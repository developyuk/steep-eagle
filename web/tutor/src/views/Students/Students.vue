<template lang="pug">
  template-main#students
    empty(v-if="!!sessions && !sessions.length")
    .mdc-list-group(v-else)
      template(v-for="(v, i) in sessions")
        h3.mdc-list-group__subheader
          placeholder(:value="v.session.class.program_module.module.name").module
          br
          placeholder(:value="v.session.class.branch.name").branch
          | &nbsp;&nbsp;
          .day-time
            placeholder(:value="v.session.class.start_at" val-empty="00:00")
            | &nbsp;-&nbsp;
            placeholder(:value="v.session.class.finish_at" val-empty="00:00")
        ul.mdc-list
          template(v-for="(vv,ii) in v.session.class.students")
            card(:key="`${i}.${ii}`" :index="`${i}.${ii}`" :sid="v.id" :student="vv.student" :isActive="vv.isActive")
            hr.mdc-list-divider
    .mdc-snackbar(aria-live='assertive' aria-atomic='true' aria-hidden='true')
      .mdc-snackbar__text
      .mdc-snackbar__action-wrapper
        button.mdc-snackbar__action-button(type='button')
</template>

<script>
  import axios from 'axios';
  import moment from 'moment';
  import mixinHal from '../../mixins/hal';
  import mixinDom from '../../mixins/dom';
  import _isEqual from 'lodash/isEqual';
  import _cloneDeep from 'lodash/cloneDeep';
  import _findIndex from 'lodash/findIndex';
  import {mapState, mapMutations} from 'vuex';
  import mqtt from "mqtt";
  import {MDCSnackbar} from '@material/snackbar';
  import TemplateMain from '@/components/views/Main';

  const placeholderStudents =
    [1,2].map(v => {
      const students = [1,2,3].map(vv => {
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
          "class": {

            "branch_id": 1,
            "start_at": "",
            "program_module": {
              "module_id": 10,
              "module": {
                "name": "",
                "image": "https://images.weserv.nl/?il&bg=2084C6&w=96&h=96&t=square&url=dl.dropboxusercontent.com/s/kt4ww4h0s4ptldr/html.png",
                "id": 10
              },
              "id": 11,
              "program_id": 1
            },
            students,
            "program_module_id": 11,
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
    name: 'students',
    mixins: [mixinHal, mixinDom],
    components: {
      TemplateMain,
      'placeholder': () => import('@/components/Placeholder'),
      'card': () => import('./Card'),
      'empty': () => import('./Empty'),
    },
    computed: {
      ...mapState(['currentAuth', 'currentStudentSession']),
    },
    watch: {
      currentStudentSession(v) {
        const {sid, uid, on: stateOn} = v;
        const [i, ii] = this.getSessionIndex(sid, uid);

        switch (stateOn) {
          case 'tapStudent': {
            const item = _cloneDeep(this.sessions[i].session.class.students[ii]);
            this.sessions.forEach((v, i, a) => {
              v.session.class.students.forEach((v2, i2, a2) => {
                this.$set(a2[i2], 'isActive', false)
              });
            });
//            console.log(i, ii, item['isActive']);
            this.$set(this.sessions[i].session.class.students[ii], 'isActive', !item['isActive']);
            break;
          }
          case 'clickRating': {
            const $el = this.getSessionElement(sid, uid);
            const {form} = v;
            for (const k in form) {
              const value = form[k];
              if (value) {
                if (k !== 'review') {
                  const $elIcon = $el.querySelector(`.${k} .material-icons[data-value="${value}"]`);
//                  console.log(`.${k} .material-icons[data-value="${value}"]`, $elIcon);
                  const $rating = $elIcon.closest('.rating');
                  $rating.querySelectorAll(`.material-icons`).forEach(v => v.classList.remove('is-active'));
                  [...Array(parseInt(value)).keys()].forEach(v => {
                    $rating.querySelector(`.material-icons[data-value='${v + 1}']`).classList.add('is-active')
                  });
                } else {
                  const $elTextbox = $el.querySelector(`.review textarea`);
                  $elTextbox.innerText = value;
                }
              }
            }
//            console.log(form);
            break;
          }
        }
      }
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
      getSessionIndex(sessionId, studentId) {
        let i, ii;
        sessionId = parseInt(sessionId);
        studentId = parseInt(studentId);

        i = _findIndex(this.sessions, v => {
          ii = _findIndex(v.session.class.students, vv => {
            return vv.student.id === studentId && vv.id === sessionId
          });
          return ii > -1;
        });
        return [i, ii]
      },
      getSessionElement(sessionTutorId, studentId) {
        const [i, ii] = this.getSessionIndex(sessionTutorId, studentId);

        return Array.from(this.$el.querySelectorAll('ul.mdc-list > li')).filter(v => {
          return v.dataset.index === `${i}.${ii}`
        })[0];
      },
      getStudentsSessions() {
        const url = `${process.env.VUE_APP_DBAPI}/students`;

        axios.get(url)
          .then(response => {
            this.sessions = [];
            this.sessions = Object.assign([], this.sessions, response.data._items);
//            this.sessions = response.data._items;
          })
          .catch(error => console.log(error));
      }
    },
    mounted() {
      this.snackbar = new MDCSnackbar(this.$el.querySelector('.mdc-snackbar'));
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

          const {sid, uid, name, stsId, stsEt} = parsedMessage;
          const $el = this.getSessionElement(sid, uid);
          const [i, ii] = this.getSessionIndex(sid, uid);

          switch (parsedMessage.on) {
            case 'undoRateReview': {
//              console.log($el);
//              $el.className = "animated slideInLeftHeight";
              $el.style.marginLeft = 0;
              let snackbarOpts = {
                message: `Undo ${name.split(" ")[0].toUpperCase()}`
              };
              this.snackbar.show(snackbarOpts);
              break;
            }
            case 'successRateReview': {
              const {by: msgBy} = parsedMessage;
              $el.className = "hide";

              this.sessions.forEach((v, i, a) => {
                v.session.class.students.forEach((v2, i2, a2) => {
                  this.$set(a2[i2], 'isActive', false)
                });
              });

              let snackbarOpts = {
                message: `Submit ${name.split(" ")[0].toUpperCase()}`,
              };
              if (msgBy.id === this.currentAuth.id) {
                snackbarOpts = Object.assign(snackbarOpts, {
                  actionText: 'Undo',
                  actionHandler: () => {
                    this.mqtt
                      .publish('students', JSON.stringify(Object.assign(parsedMessage, {on: 'undoRateReview'})));
                    const url = `${process.env.VUE_APP_DBAPI}/sessions_tutors_students/${stsId}`;

                    axios.delete(url, {headers: {'If-Match': stsEt}})
                      .then(response => {
                        console.log(response.data)
                      })
                      .catch(error => {
                        console.log(error);
                        this.mqtt
                          .publish('students', JSON.stringify(Object.assign(parsedMessage, {on: 'successRateReview'})));
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
