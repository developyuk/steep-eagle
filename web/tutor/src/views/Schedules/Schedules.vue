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
          item(v-for="(v2,i2) in v._items" :item="v2" :key="`${v.date}-${v2._id}`" @click-start="onClickStart")

    form(@submit.prevent="checkPin($event)")
      my-dialog(@mounted="onMountedDialog")
        span Insert 1234 to activate&nbsp;
          placeholder(:value="currentClass.module.name.toUpperCase()" val-empty="lorem ipsum")
          | .
        p
        input(type="text" name="username" v-model.trim="pin")
        .errMsg(v-if="errMsg") {{errMsg}}
        template(slot="footer")
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--cancel(type='button') No
          button.mdc-button.mdc-dialog__footer__button(type='submit') Yes
    snackbar(@mounted="onMountedSnackbar")
</template>

<script>
import axios from "axios";
import moment from "moment";
import _findIndex from "lodash/findIndex";
import _range from "lodash/range";
import { mapState, mapMutations } from "vuex";
import mqtt from "mqtt";
import { MDCDialog } from "@material/dialog";
import { MDCRipple } from "@material/ripple";

import TemplateMain from "@/components/views/Main";
import MyDialog from "@/components/Dialog";
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
  components: { TemplateMain, MyDialog, Empty, Item, Placeholder, Snackbar },
  data() {
    return {
      pin: null,
      classes: placeholderSchedules,
      currentClass: {
        _id: 0,
        module: { name: "" }
      },
      dialog: null,
      snackbar: null,
      errMsg: null,
      mqtt: null,

      undoItem: null
    };
  },
  computed: {
    ...mapState(["currentAuth"])
  },
  methods: {
    ...mapMutations(["nextMqtt"]),
    onMountedSnackbar(e) {
      this.snackbar = e;
    },
    onMountedDialog(e) {
      this.dialog = e;
    },
    onClickStart(e) {
      this.dialog.show();
      const { i, i2 } = this.findClassById(e.id);
      this.currentClass = this.classes[i]._items[i2];
    },
    checkPin(e) {
      if (this.pin === "1234") {
        const data = {
          class: this.currentClass._id
        };
        console.log(data);

        axios
          .post(`${process.env.VUE_APP_API}/attendances_tutors`, data)
          .then(response => {
            this.pin = null;
            this.undoItem = response.data;
          });

        this.snackbar.show({
          message: `You started a class`
          // actionText: "Undo",
          // actionHandler: () => {
          //   if (this.undoItem) {
          //     axios
          //       .delete(
          //         `${process.env.VUE_APP_API}/attendances_tutors/${
          //           this.undoItem._id
          //         }`,
          //         { headers: { "If-Match": this.undoItem._etag } }
          //       )
          //       .then(response => {
          //         this.snackbar.show({
          //           message: `You undo a class`
          //           // message: `Undo ${msgClass.module.name.toUpperCase()}`
          //         });
          //       });
          //   }
          // }
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
    },
    findClassById(id) {
      let i, i2;
      i = _findIndex(this.classes, v => {
        i2 = _findIndex(v._items, { _id: id });
        return i2 > -1;
      });
      //        console.log(i, ii, this.classes);
      return { i, i2 };
    }
  },
  mounted() {
    this.dialog = new MDCDialog(this.$el.querySelector("#my-mdc-dialog"));

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
