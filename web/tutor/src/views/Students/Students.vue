<template lang="pug">
  template-main#students
    empty(v-if="!!attendances && !attendances.length")
    .mdc-list-group(v-else)
      template(v-for="(v, i) in attendances")
        h3.mdc-list-group__subheader
          placeholder(:value="v.attendance.module.name").module
          br
          placeholder(:value="v.attendance.class.branch.name").branch
          | &nbsp;&nbsp;
          .day-time
            placeholder(:value="v.attendance.class._start" val-empty="00:00")
            | &nbsp;-&nbsp;
            placeholder(:value="v.attendance.class._finish" val-empty="00:00")
        ul.mdc-list
          template(v-for="(vv,ii) in v.students")
            card(:key="`${i}.${ii}`" :index="`${i}.${ii}`" :stid="v._id" :sid="v.attendance._id" :tid="v.tutor" :student="vv.student" :isActive="vv.isActive" @tap-student="onTapStudent" @absenced="onAbsenced" @presenced="onPresenced")
            hr.mdc-list-divider

    my-dialog(@mounted="onMountedDialog")
      span Loading..
    snackbar(@mounted="onMountedSnackbar")
</template>

<script>
import axios from "axios";
import moment from "moment";
import _isEqual from "lodash/isEqual";
import _cloneDeep from "lodash/cloneDeep";
import _findIndex from "lodash/findIndex";
import { mapState, mapMutations } from "vuex";
import mqtt from "mqtt";
import _range from "lodash/range";

import MyDialog from "@/components/Dialog";
import TemplateMain from "@/components/views/Main";
import Placeholder from "@/components/Placeholder";
import Card from "./Card";
import Empty from "./Empty";
import Snackbar from "@/components/Snackbar";

const placeholderStudents = _range(3).map(v => {
  return {
    attendance: {
      class: {
        _start: "",
        _finish: "",
        branch: {
          name: "",
          id: v
        },
        _id: v,
        day: ""
      },
      _id: v,
      module: {
        _id: v,
        name: ""
      }
    },
    tutor: v,
    _id: v,
    students: _range(3).map(vv => {
      return {
        class: vv,
        _id: vv,
        student: {
          username: "",
          name: "",
          photo:
            "https://images.weserv.nl/?il&w=64&url=www.shareicon.net/download/2016/01/16/249030_nobita_256x256.png",
          _id: vv
        }
      };
    })
  };
});

export default {
  components: { TemplateMain, Placeholder, Card, Empty, Snackbar, MyDialog },
  computed: {
    ...mapState(["currentAuth"])
  },
  data() {
    return {
      attendances: placeholderStudents,

      snackbar: null,
      dialog: null,
      mqtt: null
    };
  },
  methods: {
    ...mapMutations(["nextMqtt"]),
    onMountedDialog(e) {
      this.dialog = e;
      this.dialog.escapeKeyAction = "";
      this.dialog.scrimClickAction = "";
    },
    onProgress(e) {
      const { url, data } = e;
      this.dialog.open();

      axios
        .post(url, data)
        .then(response => {
          const item = response.data;

          this.snackbar.show({
            message: `You submitted a progress`,
            actionText: "Undo",
            actionHandler: () => {
              this.snackbar.show({
                message: `You undo a progress`
              });
              axios.delete(
                `${process.env.VUE_APP_API}/attendances_students/${item._id}`,
                { headers: { "If-Match": item._etag } }
              );
              // .then(response => { })
              // .catch(error => console.log(error));
            }
          });

          this.dialog.close();
        })
        .catch(error => {
          this.dialog.close();
          this.snackbar.show({
            message: `You failed submit a progress`
          });
        });
    },
    onAbsenced(e) {
      this.onProgress(e);
    },
    onPresenced(e) {
      this.onProgress(e);
    },
    onMountedSnackbar(e) {
      this.snackbar = e;
    },
    onTapStudent(e) {
      const { sid, uid } = e;
      const [i, ii] = this.getAttendanceIndex(sid, uid);

      const item = _cloneDeep(this.attendances[i].students[ii]);
      this.attendances.forEach((v, i, a) => {
        v.students.forEach((v2, i2, a2) => {
          this.$set(a2[i2], "isActive", false);
        });
      });
      setTimeout(_ => {
        this.$set(
          this.attendances[i].students[ii],
          "isActive",
          !item["isActive"]
        );
      }, 50);
    },
    getAttendanceIndex(sid, uid) {
      let i, ii;
      sid = parseInt(sid);
      uid = parseInt(uid);

      i = _findIndex(this.attendances, v => {
        ii = _findIndex(v.students, vv => {
          return vv.student._id === uid && v._id === sid;
        });
        return ii > -1;
      });
      return [i, ii];
    },
    getStudentsAttendances(params = { forceRefresh: false }) {
      const headers = {};
      if (params.forceRefresh) {
        headers["Cache-Control"] = "no-cache";
      }
      axios
        .get(`${process.env.VUE_APP_API}/students/attendances`, { headers })
        .then(response => {
          this.attendances = response.data._items;
          // console.log(this.attendances);
          this.attendances.forEach((v, i) => {
            this.$set(v.attendance, "module", {
              _id: v.attendance.module,
              name: ""
            });
            this.$set(v.attendance, "class", {
              _id: v.attendance.class,
              _start: "",
              _finish: "",
              day: "",
              branch: {
                _id: 0,
                name: ""
              }
            });
            v.students.forEach(v2 => {
              this.$set(v2, "student", {
                username: "",
                name: "",
                photo:
                  "https://images.weserv.nl/?il&w=64&url=www.shareicon.net/download/2016/01/16/249030_nobita_256x256.png",
                _id: v2.student
              });
            });
            setTimeout(_ => {
              axios
                .get(
                  `${process.env.VUE_APP_API}/modules/${
                    v.attendance.module._id
                  }`
                )
                .then(response =>
                  this.$set(v.attendance, "module", response.data)
                );
              axios
                .get(
                  `${process.env.VUE_APP_API}/classes/${v.attendance.class._id}`
                )
                .then(response => {
                  this.$set(v.attendance, "class", response.data);
                  this.$set(v.attendance.class, "branch", {
                    _id: v.attendance.class.branch,
                    name: ""
                  });
                  axios
                    .get(
                      `${process.env.VUE_APP_API}/branches/${
                        v.attendance.class.branch._id
                      }`
                    )
                    .then(response =>
                      this.$set(v.attendance.class, "branch", response.data)
                    );
                });
              v.students.forEach(v2 => {
                const config = {
                  // params: { where: { role: "student" } }
                };
                axios
                  .get(
                    `${process.env.VUE_APP_API}/users/${v2.student._id}`,
                    config
                  )
                  .then(response => this.$set(v2, "student", response.data));
              });
            }, i * 500);

            // this.$set(v.attendance, "class", { branch: "" });
          });
        });
    }
  },
  mounted() {
    this.getStudentsAttendances();

    this.mqtt = mqtt
      .connect(process.env.VUE_APP_WS)
      .on("connect", () => {
        const topic = `${process.env.VUE_APP_PROJECT_NAME}.students`;
        this.nextMqtt({ topic, mqtt: this.mqtt });
        this.mqtt.subscribe(topic);
      })
      .on("message", (topic, message) => {
        console.log(topic, message.toString());
        const parsedMessage = JSON.parse(message.toString());
        const { update, by } = parsedMessage;

        if (update) {
          this.getStudentsAttendances({ forceRefresh: true });
        }
      });
  },
  beforeDestroy() {
    // console.log("beforeDestroy");

    this.mqtt.end();
    this.nextMqtt(null);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";
@import "~@material/animation/functions";

#students {
  position: relative;
}

#form-rate-review {
  transition: mdc-animation-enter(height, 300ms);
}

.mdc-list {
  background-color: map-get($palettes, orange-lighten1);
  overflow: hidden;
}

span.module {
  text-transform: uppercase;

  /*max-width: 8rem;*/
  /*text-overflow: ellipsis;*/
  /*overflow: hidden;*/
  /*white-space: nowrap;*/
  /*display: inline-block;*/
  /*vertical-align: middle;*/
}

span.branch,
span.day-time {
  text-transform: capitalize;
}

.day-time {
  float: right;
}
</style>
