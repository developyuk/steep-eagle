<template lang="pug">
  template-main#students
    spinner(v-if="!sessions")
    .empty(v-if="!!sessions && !sessions.length")
      router-link(to="/" data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Start Class
      p or
      router-link(to="/" data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised.mdc-button--compact Activate
    .mdc-list-group(v-else)
      template(v-for="(v,i) in sessions")
        h3.mdc-list-group__subheader
          span.module(v-if="!!v._embedded.class._embedded.module") {{v._embedded.class._embedded.module.name}}
          | &nbsp;-&nbsp;
          span.branch(v-if="!!v._embedded.class._embedded.branch") {{v._embedded.class._embedded.branch.name}}
          | &nbsp;&nbsp;
          span.day-time {{v._embedded.class.start_at}} - {{v._embedded.class.finish_at}}
        ul.mdc-list
          template(v-for="(vv,ii) in v._embedded.items")
            card(:key="`${i}.${ii}`" :index="`${i}.${ii}`" :sid="vv.id" :student="vv._embedded.student" :isActive="vv.isActive")
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
  import {MDCRipple} from '@material/ripple';
  import TemplateMain from '@/templates/TemplateMain';

  export default {
    name: 'students',
    mixins: [mixinHal, mixinDom],
    components: {
      'template-main': TemplateMain,
      'spinner': () => import('@/components/Spinner'),
      'card': () => import('./Card')
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
            const item = _cloneDeep(this.sessions[i]._embedded.items[ii]);
            this.sessions.forEach((v, i, a) => {
              v._embedded.items.forEach((v2, i2, a2) => {
                this.$set(a2[i2], 'isActive', false)
              });
            });
//            console.log(i, ii, item['isActive']);
            this.$set(this.sessions[i]._embedded.items[ii], 'isActive', !item['isActive']);
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
        sessions: null,

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
          ii = _findIndex(v._embedded.items, vv => {
            return vv._embedded.student.id === studentId && vv.id === sessionId
          });
          return ii > -1;
        });
        return [i, ii]
      },
      getSessionElement(sessionId, studentId) {
        const [i, ii] = this.getSessionIndex(sessionId, studentId);

        return Array.from(this.$el.querySelectorAll('ul.mdc-list > li')).filter(v => {
          return v.dataset.index === `${i}.${ii}`
        })[0];
      },
      getStudentsSessions() {
        const url = `${process.env.VUE_APP_API}/tutors/${this.currentAuth.id}/sessions`;

        axios.get(url)
          .then(response => {
            let sessions = response.data._embedded.items;
            if (!!sessions) {
              sessions = sessions.map(v => {
                v._embedded['class'] = {
                  _embedded: {
                    module: {name: "..."},
                    branch: {name: "..."}
                  },
                  start_at: "...",
                  finish_at: "...",
                };
                v._embedded.items = v._embedded.items.map(v2 => {
                  v2._embedded = {
                    student: {
                      name: "...",
                    }
                  };
                  return v2
                });
                return v;
              });
            }
            else {
              sessions = [];
            }
            this.sessions = sessions;
            const timeout = 1;

            if (!this.sessions.length) {
              return;
            }
            this.sessions.forEach((v, i, a) => {
              setTimeout(() => this.parseEmbedded('class', v._links.class.href, a[i]['_embedded'], (item) => {
                if (!item['_embedded']) {
                  this.$set(item, '_embedded', {
                    module: {name: "..."},
                    branch: {name: "..."},
                  });
                }
                setTimeout(() => this.parseEmbedded('branch', item._links.branch.href, item['_embedded']), timeout);
                setTimeout(() => this.parseEmbedded('module', item._links.module.href, item['_embedded']), timeout);
                return item
              }), timeout);

              setTimeout(() => {
                if (!!v._embedded.items) {
                  v._embedded.items.forEach((v2, i2, a2) => {
                    this.$set(a2[i2], 'isActive', false);
                    if (!a2[i2]['_embedded']) {
                      this.$set(a2[i2], '_embedded', {});
                    }
                    this.parseEmbedded('student', v2._links.student.href, a2[i2]['_embedded']);
                  })
                }
              }, timeout);

            });
          })
          .catch(error => console.log(error));
      }
    },
    mounted() {
      if (!!this.$el.querySelector('.mdc-button')) {
        new MDCRipple(this.$el.querySelector('.mdc-button'));
      }
      this.snackbar = new MDCSnackbar(this.$el.querySelector('.mdc-snackbar'));
      setTimeout(() => this.getStudentsSessions(), 1);

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

          const {sid, uid, name} = parsedMessage;
          const $el = this.getSessionElement(sid, uid);
          const [i, ii] = this.getSessionIndex(sid, uid);

          switch (parsedMessage.on) {
            case 'undoRateReview': {
//              console.log($el);
              $el.className = "animated slideInLeftHeight";
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
                v._embedded.items.forEach((v2, i2, a2) => {
                  this.$set(a2[i2], 'isActive', false)
                });
              });
//            this.$set(this.sessions[i]._embedded.items[ii], 'isActive', false);
              let snackbarOpts = {
                message: `Submit ${name.split(" ")[0].toUpperCase()}`,
              };
              if (msgBy.id === this.currentAuth.id) {
                snackbarOpts = Object.assign(snackbarOpts, {
                  actionText: 'Undo',
                  actionHandler: () => {
                    this.mqtt
                      .publish('students', JSON.stringify(Object.assign(parsedMessage, {on: 'undoRateReview'})));
                    const url = `${process.env.VUE_APP_API}/sessions/${sid}/students/${uid}`;

                    axios.delete(url)
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

  .empty {
    text-align: center;
    position: absolute;
    top: 50%;
    transform: translateY(-70%);
    width: 100%;
    p {
      text-align: center;
      margin: .75rem 0;
    }
    .mdc-button {
      font-size: .675rem;
      background-color: map-get($palettes, green);
      width: 5rem;
    }
  }

  .mdc-list {
    background-color: map-get($palettes, orange-lighten1);
    overflow: hidden;
  }

  span.module {
    text-transform: uppercase;

    max-width: 8rem;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    display: inline-block;
    vertical-align: middle;
  }

  span.branch, span.day-time {
    text-transform: capitalize;
  }

  span.day-time {
    float: right;
  }

</style>
