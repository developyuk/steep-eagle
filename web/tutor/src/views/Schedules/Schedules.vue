<template lang="pug">
  template-main#schedules
    empty(v-if="!!schedules && !schedules.length")
    .mdc-list-group
      template(v-for="(v,i) in schedules")
        h3.mdc-list-group__subheader
          span.date-day
            .date
              placeholder(:value="v.day" val-empty="00")
            .day
              placeholder(:value="v.weekday" val-empty="lorem ipsum")
          span.text
            placeholder(:value="v.delta")
        .mdc-list
          item(v-for="(v2,i2) in v._items" :item="v2" :key="`${i}-${i2}`" v-model="currentItem")

    form(@submit.prevent="checkPin($event)")
      my-dialog(v-model="dialog")
        template(slot="title") Are you sure?
        span Insert 1234 to activate&nbsp;
          placeholder(:value="currentItem.class.module.name.toUpperCase()" val-empty="lorem ipsum")
        p &nbsp;
          input(type="text" name="username" v-model.trim="pin")
        .errMsg(v-if="errMsg") {{errMsg}}
        template(slot="actions")
          button.mdc-button.mdc-dialog__button(data-mdc-dialog-action="no" type='button') No
          button.mdc-button.mdc-dialog__button(type='submit') Yes

    my-dialog-loading(v-model="dialogLoading" escapeKeyAction="" scrimClickAction="")
    snackbar(v-model="snackbar")
</template>

<script>
import axios from "axios";
import moment from "moment";
import _findIndex from "lodash/findIndex";
import _range from "lodash/range";
import { mapState, mapMutations } from "vuex";
import mqtt from "mqtt";

import TemplateMain from "@/components/views/Main";
import MyDialog from "@/components/Dialog";
import MyDialogLoading from "@/components/DialogLoading";
import Empty from "./Empty";
import Item from "./Item";
import Placeholder from "@/components/Placeholder";
import Snackbar from "@/components/Snackbar";

const phSchedule = _ => {
  return {
    date: "",
    text: "",
    dateDay: "",
    day: "",
    _items: _range(2).map(vv => {
      return {
        class: {
          _id: 0,
          module: {
            _id: 0,
            image:
              "data:image/gif;base64,R0lGODlhAQABAIAAAP///wAAACH5BAEAAAAALAAAAAABAAEAAAICRAEAOw==",
            name: ""
          },
          _start: "",
          _finish: "",
          tutor: {
            _id: 0,
            name: "",
            username: "",
            email: ""
          },
          branch: {
            _id: 0,
            name: ""
          }
        },
        nextFinish: "",
        nextStart: ""
      };
    })
  };
};

export default {
  components: {
    TemplateMain,
    MyDialog,
    MyDialogLoading,
    Empty,
    Item,
    Placeholder,
    Snackbar
  },
  data() {
    return {
      pin: null,
      // forceSchedules: false,
      currentItem: {
        class: {
          _id: 0,
          module: { name: "" }
        }
      },
      dialog: null,
      dialogLoading: null,
      snackbar: null,
      errMsg: null,
      mqtt: null
    };
  },
  asyncData: {
    schedules: {
      get() {
        const config = {
          params: {
            sort: "start",
            embedded: JSON.stringify({
              module: 1,
              branch: 1,
              tutor: 1
            })
          },
          headers: {}
        };

        // if (this.forceSchedules) {
        config["headers"]["Cache-Control"] = "no-cache";
        // }
        return axios.get(`${process.env.VUE_APP_API}/schedules`, config);
      },
      transform(r) {
        this.dialogLoading.close();
        return r.data._items;
      },
      default: _range(3).map(phSchedule)
    }
  },
  computed: {
    ...mapState(["currentAuth"])
  },
  watch: {
    currentItem(v, ov) {
      this.dialog.open();
    }
  },
  methods: {
    ...mapMutations(["nextMqtt"]),
    checkPin(e) {
      if (this.pin === "1234") {
        this.dialog.close();
        this.dialogLoading.open();
        const data = {
          class: this.currentItem.class._id
        };

        axios
          .post(`${process.env.VUE_APP_API}/attendances_tutors`, data)
          .then(response => {
            this.pin = null;
            const item = response.data;

            this.snackbar.show({
              message: `You started a class`,
              actionText: "Undo",
              actionHandler: () => {
                this.dialogLoading.open();
                axios
                  .delete(
                    `${process.env.VUE_APP_API}/attendances_tutors/${item._id}`,
                    { headers: { "If-Match": item._etag } }
                  )
                  .then(response => {
                    this.snackbar.show({
                      message: `You undo a class`
                    });
                  });
              }
            });
          })
          .catch(error => {
            this.errMsg = "failed!";
            this.snackbar.show({
              message: `You failed start a class`
            });
          });
        this.dialog.close();
      } else {
        this.errMsg = "invalid. Check pin again!";
      }
    }
  },
  mounted() {
    this.$on("schedules$reset", resettingResponse => {
      this.dialogLoading.close();
    });

    this.mqtt = mqtt
      .connect(process.env.VUE_APP_WS)
      .on("connect", () => {
        const topic = `${process.env.VUE_APP_PROJECT_NAME}.schedules`;
        this.nextMqtt({ topic, mqtt: this.mqtt });
        this.mqtt.subscribe(topic);
      })
      .on("message", (topic, message) => {
        console.log(topic, message.toString());
        const parsedMessage = JSON.parse(message.toString());
        const { update, by } = parsedMessage;

        if (update) {
          this.dialogLoading.open();
          this.schedules$refresh();
        }
      });
  },
  beforeDestroy() {
    this.mqtt.end();
    this.nextMqtt(null);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";
@import "~sass-bem";

#schedules {
  position: relative;
}

.mdc-list {
  padding: 0;
}

.empty {
  text-align: center;
  position: absolute;
  top: 50%;
  transform: translateY(-70%);
  width: 100%;
}

.errMsg {
  padding: 0.5rem 0;
  background-color: #fff;
  margin: 0.5rem 0;
}

.mdc-list-group__subheader {
  height: 1.5rem;
  .date-day {
    float: left;
    * {
      text-align: center;
      line-height: 1rem;
    }
    .date {
      font-size: 1rem;
    }
    .day {
      font-size: 0.625rem;
      text-transform: capitalize;
    }
  }
  .text {
    float: right;
    font-size: 0.7rem;
    text-transform: capitalize;
    color: map-get($palettes, green-darken1);
  }
}

// .image {
//   width: 4rem;
//   height: 4rem;
//   background-color: #808080;

//   @include m(construct2 construct2-part-ii dummies-construct) {
//     background-color: #4f4f4f;
//   }
//   @include m(unity-2d weblite project-based dummies-web) {
//     background-color: #f2f2f2;
//   }
// }
</style>
