<template lang="pug">
  #students.mdc-typography
    header1
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
    tab-bottom
</template>

<script>
  import axios from 'axios';
  import moment from 'moment';
  import Hammer from 'hammerjs';
  import mixinHal from '../../mixins/hal';
  import mixinDom from '../../mixins/dom';
  import mixinImage from '../../mixins/image';
  import _isEqual from 'lodash/isEqual';
  import _cloneDeep from 'lodash/cloneDeep';
  import _findIndex from 'lodash/findIndex';

  export default {
    name: 'students',
    mixins: [mixinHal, mixinDom, mixinImage],
    components: {
      'spinner': () => import('@/components/Spinner'),
      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': () => import('@/components/Header'),
      'card': () => import('./Card')
    },
    data() {
      return {
        sessions: null,
        currentAuth: null,
        currentStudent: {
          sid: "0",
          uid: "0",
          name: "",
          is: "empty",
        },

        snackbar: null,
        ws: null,
      }
    },
    methods: {
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

        return Array.from(document.querySelectorAll('ul.mdc-list > li')).filter(v => {
          return v.dataset.index === `${i}.${ii}`
        })[0];
      },
      getStudentsSessions() {
        const url = `${process.env.API}/tutors/${this.currentAuth.id}/sessions`;

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
                      photo: "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=",
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
                    this.parseEmbedded('student', v2._links.student.href, a2[i2]['_embedded'], item => {
                      item.photo = this.normalizeImage(item.photo);
                      return item;
                    });
                  })
                }
              }, timeout);

            });
          })
          .catch(error => console.log(error));
      }
    },
    mounted() {
      const _this = this;
      this.snackbar = mdc.snackbar.MDCSnackbar.attachTo(document.querySelector('.mdc-snackbar'));
      this.$bus.$on('currentAuth', (auth) => {
        if (!!this.currentAuth) {
          return;
        }
        this.currentAuth = auth;
        this.getStudentsSessions();

        window.mdc.autoInit();
      });
      this.$bus.$on('onCreatedWs', (data) => {
        if (!!this.ws) {
          return;
        }
        const {wsStudents} = data;
        this.ws = wsStudents;

        this.ws.onmessage = (e) => {
          console.log(e.data);
          const data = JSON.parse(e.data);
          console.log(data);
          const {sid, uid, name} = data.v;
          const $el = _this.getSessionElement(sid, uid);
          const [i, ii] = _this.getSessionIndex(sid, uid);

          switch (data.on) {
            case 'tapStudent': {
              const item = _cloneDeep(_this.sessions[i]._embedded.items[ii]);

              _this.sessions.forEach((v, i, a) => {
                v._embedded.items.forEach((v2, i2, a2) => {
                  _this.$set(a2[i2], 'isActive', false)
                });
              });
              console.log(i, ii, item['isActive']);
              _this.$set(_this.sessions[i]._embedded.items[ii], 'isActive', !item['isActive']);
              break;
            }
            case 'clickRating': {
              const {form} = data.v;
              for (const k in form) {
                const value = form[k];
                if (value) {
                  if (k !== 'review') {
                    const $elIcon = $el.querySelector(`.${k} .material-icons[data-value="${value}"]`);
                    console.log(`.${k} .material-icons[data-value="${value}"]`, $elIcon);
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
              console.log(form);
              break;
            }
            case 'undoRateReview': {
              $el.className = "animated slideInLeftHeight";
              $el.style.marginLeft = 0;
              break;
            }
            case 'successRateReview': {
              $el.className = "hide";

              _this.sessions.forEach((v, i, a) => {
                v._embedded.items.forEach((v2, i2, a2) => {
                  _this.$set(a2[i2], 'isActive', false)
                });
              });
//            this.$set(this.sessions[i]._embedded.items[ii], 'isActive', false);

              _this.snackbar.show({
                message: `${name.split(" ")[0].toUpperCase()} has been saved`,
                actionText: 'Undo',
                actionHandler: function () {
                  _this.$bus.$emit('onUndoRateReview', data.v);
                  const url = `${process.env.API}/sessions/${sid}/students/${uid}`;

                  axios.delete(url)
                    .then(response => {
                      _this.$bus.$emit('onUndoRateReview', data.v);
                    })
                    .catch(error => {
                      console.log(error);
                      _this.$bus.$emit('onSuccessRateReview', data.v);
                    });
                }
              });
              break;
            }
          }

        };
      }).$emit('reqCreatedWs')
        .$on('onTapStudent', (v) => {
          console.log(this.ws);
          this.ws.send(JSON.stringify({on: 'tapStudent', v}))
        })
        .$on('onClickRating', (v) => {
          this.ws.send(JSON.stringify({on: 'clickRating', v}))
        })
        .$on('onUndoRateReview', (v) => {
          this.ws.send(JSON.stringify({on: 'undoRateReview', v}));
        })
        .$on('onSuccessRateReview', (v) => {
          this.ws.send(JSON.stringify({on: 'successRateReview', v}));
        });

    },
    beforeDestroy() {
      this.$bus
        .$off('onCreatedWs')
        .$off('onTapStudent')
        .$off('onClickRating')
        .$off('onUndoRateReview')
        .$off('onSuccessRateReview');
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";
  @import "@material/animation/functions";

  #students {
    position: relative;
    height: 100vh;
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
