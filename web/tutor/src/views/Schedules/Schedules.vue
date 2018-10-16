<template lang="pug">
  template-main#schedules
    empty(v-if="!!classes && !classes.length")
    .mdc-list-group
      template(v-for="(v,i) in classes")
        h3.mdc-list-group__subheader
          span.date-day
            .date
              placeholder(:value="v.day" val-empty="00")
            .day
              placeholder(:value="v.weekday" val-empty="lorem ipsum")
          span.text
            placeholder(:value="v.delta")
        .mdc-list
          item(v-for="(v2,i2) in v._items" :item="v2" :key="`${v.date}-${v2._id}`" v-model="currentClass")

    form(@submit.prevent="checkPin($event)")
      my-dialog(v-model="dialog")
        template(slot="title") Are you sure?
        span Insert 1234 to activate&nbsp;
          placeholder(:value="currentClass.module.name.toUpperCase()" val-empty="lorem ipsum")
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

const placeholderSchedules = _range(3).map(v => {
  return {
    date: v,
    text: "",
    dateDay: "",
    day: "",
    _items: _range(2).map(vv => {
      return {
        _id: vv,
        module: {
          image: "https://via.placeholder.com/48?text=image",
          name: ""
        },
        _start: "",
        _finish: "",
        finish: "",
        start: "",
        tutor: {
          name: "",
          username: "",
          email: ""
        },
        branch: {
          name: ""
        },
        last_attendances: { _items: [] }
      };
    })
  };
});

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
      classes: placeholderSchedules,
      currentClass: {
        _id: 0,
        module: { name: "" }
      },
      dialog: null,
      dialogLoading: null,
      snackbar: null,
      errMsg: null,
      mqtt: null
    };
  },
  computed: {
    ...mapState(["currentAuth"])
  },
  watch: {
    currentClass(v, ov) {
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
          class: this.currentClass._id
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
            this.dialogLoading.close();
            this.errMsg = "failed!";
            this.snackbar.show({
              message: `You failed start a class`
            });
          });
        this.dialog.close();
      } else {
        this.errMsg = "invalid. Check pin again!";
      }
    },
    getSchedules(params = { forceRefresh: false }) {
      const config = {};
      config["params"] = {
        _page: 1,
        _max_results: 2
      };
      config["headers"] = {};
      if (params.forceRefresh) {
        config["headers"]["Cache-Control"] = "no-cache";
      }
      axios
        .get(`${process.env.VUE_APP_API}/schedules`, config)
        .then(response => {
          this.dialogLoading.close();
          this.classes = response.data._items;
          // console.log(this.classes);
          this.classes.forEach((v, i) => {
            v._items.forEach((v2, i2) => {
              setTimeout(_ => {
                axios
                  .get(`${process.env.VUE_APP_API}/modules/${v2.module}`)
                  .then(response => this.$set(v2, "module", response.data));
                axios
                  .get(`${process.env.VUE_APP_API}/branches/${v2.branch}`)
                  .then(response => this.$set(v2, "branch", response.data));
                axios
                  .get(`${process.env.VUE_APP_API}/users/${v2.tutor}`)
                  .then(response => this.$set(v2, "tutor", response.data));
              }, (i + i2) * 250);
            });
          });
        });
    }
  },
  mounted() {
    this.getSchedules();

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
          this.getSchedules({ forceRefresh: true });
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

.image {
  width: 4rem;
  height: 4rem;
  background-color: #c3c3c3;

  @include m(construct2 construct2-part-ii dummies-construct) {
    background-color: #4f4f4f;
  }
  @include m(unity-2d weblite project-based dummies-web) {
    background-color: #f2f2f2;
  }
}
</style>
