<template lang="pug">
  template-main#students
    empty(v-if="!!sessions && !sessions.length")
    .mdc-list-group(v-else)
      template(v-for="(v, i) in sessions")
        h3.mdc-list-group__subheader
          placeholder(:value="v.session.class_.module.name").module
          br
          placeholder(:value="v.session.class_.branch.name").branch
          | &nbsp;&nbsp;
          .day-time
            placeholder(:value="v.session.class_.start_at" val-empty="00:00")
            | &nbsp;-&nbsp;
            placeholder(:value="v.session.class_.finish_at" val-empty="00:00")
        ul.mdc-list
          template(v-for="(vv,ii) in v.session.class_.students")
            card(:key="`${i}.${ii}`" :index="`${i}.${ii}`" :stid="v.id" :sid="v.session.id" :tid="v.tutor" :student="vv.student" :isActive="vv.isActive" @tap-student="onTapStudent")
            hr.mdc-list-divider

    snackbar(@mounted="onMountedSnackbar")
</template>

<script>
  import axios from 'axios';
  import moment from 'moment';
  import _isEqual from 'lodash/isEqual';
  import _cloneDeep from 'lodash/cloneDeep';
  import _findIndex from 'lodash/findIndex';
  import {mapState, mapMutations} from 'vuex';
  import mqtt from "mqtt";
  import _range from 'lodash/range';
  import TemplateMain from '@/components/views/Main';

  const placeholderStudents =
    _range(2).map(v => {
      const students = _range(3).map(vv => {
        return {
          "class_id": 12,
          "student_id": 84,
          "id": 71,
          "student": {
            "username": "bruce banner",
            "_updated": "Fri, 06 Jul 2018 21:46:21 GMT",
            "pass_": null,
            "email": null,
            "role": "student",
            "profile": {
              "dob": null,
              "photo": null,
              "user_id": 84,
              "name": ""
            },
            "id": 84
          }
        }
      });
      return {

        "session": {

          "class_id": 12,
          "class_": {

            "branch_id": 1,
            "start_at": "",
            "module_id": 10,
            "module": {
              "name": "",
              "image": "https://images.weserv.nl/?il&bg=2084C6&w=96&h=96&t=square&url=dl.dropboxusercontent.com/s/kt4ww4h0s4ptldr/html.png",
              "id": 10
            },
            students,
            "tutor_id": 83,
            "finish_at": "",
            "branch": {
              "name": "",
              "address": null,
              "id": 1
            },
            "id": 12,
            "day": "saturday"
          },
          "id": 29,

        },
        "tutor_id": 83,

        "session_id": 29,
        "id": 29
      }
    });

  export default {
    components: {
      TemplateMain,
      'placeholder': () => import('@/components/Placeholder'),
      'card': () => import('./Card'),
      'empty': () => import('./Empty'),
      'snackbar': () => import('@/components/Snackbar'),
    },
    computed: {
      ...mapState(['currentAuth']),
    },
    data() {
      return {
        sessions: placeholderStudents,

        snackbar: null,
        mqtt: null,
      }
    },
    methods: {
      ...mapMutations(['nextMqtt']),
      onMountedSnackbar(e) {
        this.snackbar = e
      },
      onTapStudent(e) {
        const {sid, uid} = e;
        const [i, ii] = this.getSessionIndex(sid, uid);

        const item = _cloneDeep(this.sessions[i].session.class_.students[ii]);
        this.sessions.forEach((v, i, a) => {
          v.session.class_.students.forEach((v2, i2, a2) => {
            this.$set(a2[i2], 'isActive', false)
          });
        });

        this.$set(this.sessions[i].session.class_.students[ii], 'isActive', !item['isActive']);
      },
      getSessionIndex(sid, uid) {
        let i, ii;
        sid = parseInt(sid);
        uid = parseInt(uid);

        i = _findIndex(this.sessions, v => {
          ii = _findIndex(v.session.class_.students, vv => {
            return vv.student.id === uid && v.id === sid
          });
          return ii > -1;
        });
        return [i, ii]
      },
      getStudentsSessions(params = {forceRefresh: false}) {
        const url = `${process.env.VUE_APP_DBAPI}/students`;
        const headers = {};
        if (params.forceRefresh) {
          headers['Cache-Control'] = 'no-cache'
        }
        axios.get(url, {headers})
          .then(response => {
            this.sessions = !!response.data._items.length ? response.data._items : [];
          })
          .catch(error => console.log(error));
      }
    },
    mounted() {
      this.getStudentsSessions();

      this.mqtt = mqtt.connect(process.env.VUE_APP_WS);

      this.mqtt
        .on('connect', () => {
          const topic = 'students';
          this.nextMqtt({topic, mqtt: this.mqtt});
          this.mqtt.subscribe(topic);
        })
        .on('message', (topic, message) => {
          console.log(topic, message.toString());
          const parsedMessage = JSON.parse(message.toString());

          const {sid, uid, name, sts} = parsedMessage;
//          const [i, ii] = this.getSessionIndex(sid, uid);

          switch (parsedMessage.on) {
            case 'undoRateReview': {
              this.getStudentsSessions({forceRefresh: true});
              let snackbarOpts = {
                message: `Undo ${name.split(" ")[0].toUpperCase()}`
              };
              this.snackbar.show(snackbarOpts);
              break;
            }
            case 'successRateReview': {
              const {by: msgBy} = parsedMessage;
              this.getStudentsSessions({forceRefresh: true});

              let snackbarOpts = {
                message: `You submitted a progress`,
                // message: `Submit ${name.split(" ")[0].toUpperCase()}`,
              };
              if (msgBy.id === this.currentAuth.id) {
                snackbarOpts = Object.assign(snackbarOpts, {
                  actionText: 'Undo',
                  actionHandler: () => {
                    const url = `${process.env.VUE_APP_DBAPI}/sessions_students/${sts.id}`;

                    axios.delete(url, {headers: {'If-Match': sts.et}})
                      .then(response => {
                        this.mqtt
                          .publish('students', JSON.stringify(Object.assign(parsedMessage, {on: 'undoRateReview'})));
                      })
                      .catch(error => {
                        console.log(error);
                      });
                  }
                })
              }
              this.snackbar.show(snackbarOpts);
              break;
            }
          }
        });

    },
    beforeDestroy() {
      console.log('beforeDestroy');

      this.mqtt.end();
      this.nextMqtt(null);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";
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

  span.branch, span.day-time {
    text-transform: capitalize;
  }

  .day-time {
    float: right;
  }

</style>
