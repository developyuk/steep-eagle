<template lang="pug">
  template-main#schedules
    empty(v-if="!!classes && !classes.length")
    .mdc-list-group
      template(v-for="(v,i) in classes")
        h3.mdc-list-group__subheader
          span.date-day
            .date
              placeholder(:value="v.dateDay" val-empty="00")
            .day
              placeholder(:value="v.day" val-empty="lorem ipsum")
          span.text
            placeholder(:value="v.text")
        .mdc-list
          item(v-for="(vv,ii) in v._items" :item="vv" :key="`${v.date}-${vv._id}`" @click-start="onClickStart")

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
      startAt: "",
      finishAt: "",
      finishAtTs: "",
      startAtTs: "",
      tutor: {
        name: "",
        username: "",
        email: ""
      },
      branch: {
        name: ""
      },
      last_attendances:{_items:[]}
    };
  })
  };
});

export default {
  components: {
    TemplateMain,
    MyDialog,
    empty: () => import("./Empty"),
    item: () => import("./Item"),
    placeholder: () => import("@/components/Placeholder"),
    snackbar: () => import("@/components/Snackbar")
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
      snackbar: null,
      errMsg: null,
      mqtt: null
    };
  },
  computed: {
    ...mapState(["currentAuth"])
  },
  methods: {
    onMountedSnackbar(e) {
      this.snackbar = e;
    },
    onMountedDialog(e) {
      this.dialog = e;
    },
    onClickStart(e) {
      this.dialog.show();
      const { i, ii } = this.findClassById(e.id);
      this.currentClass = this.classes[i]._items[ii];
    },
    checkPin(e) {
      if (this.pin === "1234") {
        const data = {
          'class': this.currentClass._id
        };

        axios
          .post(`${process.env.VUE_APP_API}/attendances_tutors`, data)
          .then(response => {
            this.pin = null;

            axios
              .put(`${process.env.VUE_APP_API}/schedules`, data)
              .then(response => console.log(response))
              .catch(error => console.log(error));

            axios
              .put(`${process.env.VUE_APP_API}/students/attendances`)
              .then(response => console.log(response))
              .catch(error => console.log(error));
          })
          .catch(error => console.log(error));

        this.dialog.close();
      } else {
        this.errMsg = "invalid. Check pin again!";
      }
    },
    getSchedules(params = { forceRefresh: false }) {
      const config = {};
      config["headers"] = {};
      if (params.forceRefresh) {
        headers["Cache-Control"] = "no-cache";
      }
      axios
        .get(`${process.env.VUE_APP_API}/schedules`, config)
        .then(response => {
          this.classes = response.data._items;
        })
        .catch(error => console.log(error));
    },
    findClassById(id) {
      let i, ii;
      i = _findIndex(this.classes, v => {
        ii = _findIndex(v._items, { _id:id });
        return ii > -1;
      });
      //        console.log(i, ii, this.classes);
      return { i, ii };
    }
  },
  mounted() {
    //      new MDCRipple(this.$el.querySelector('.mdc-button'));
    this.dialog = new MDCDialog(this.$el.querySelector("#my-mdc-dialog"));

    this.getSchedules();

    this.mqtt = {
      topic: "schedules",
      m: mqtt.connect(process.env.VUE_APP_WS)
    };

    this.mqtt.m
      .on("connect", () => {
        this.mqtt.m.subscribe(this.mqtt.topic);
      })
      .on("message", (topic, message) => {
        console.log(topic, message.toString());
        const parsedMessage = JSON.parse(message.toString());
        const { on: msgOn, by: msgBy } = parsedMessage;
        const isCurrentUser = msgBy.id === this.currentAuth._id;
        const by = isCurrentUser ? "You" : msgBy.username;

        if (msgOn === "startYes") {
          const { class: msgClass, item: MsgItem } = parsedMessage;

          setTimeout(_ => this.getSchedules({ forceRefresh: true }), 100);
          let snackbarOpts = {
            message: `${by} started a class`
            // message: `Start ${msgClass.module.name.toUpperCase()}`
          };

          if (isCurrentUser) {
            snackbarOpts = Object.assign(snackbarOpts, {
              actionText: "Undo",
              actionHandler: () => {
                let url = `${process.env.VUE_APP_API}/attendances_tutors/${
                  MsgItem.id
                }`;

                axios
                  .delete(url, { headers: { "If-Match": MsgItem._etag } })
                  // .then(response => { })
                  .catch(error => console.log(error));
              }
            });
          }
          this.snackbar.show(snackbarOpts);
        }
        if (msgOn === "undo") {
          const { class: msgClass } = parsedMessage;

          let snackbarOpts = {
            message: `${by} undo a class`
            // message: `Undo ${msgClass.module.name.toUpperCase()}`
          };
          this.snackbar.show(snackbarOpts);
          setTimeout(_ => this.getSchedules({ forceRefresh: true }), 100);
        }
      });
  },
  beforeDestroy() {
    //      console.log('beforeDestroy');

    this.mqtt.m.end();
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
