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
            li.mdc-list-item(:data-sid="vv.id" :data-uid="vv._embedded.student.id" :data-name="vv._embedded.student.name")
              .mdc-list-item__graphic
                img(':src'="vv._embedded.student.photo")
              span.mdc-list-item__text {{vv._embedded.student.name}}
            hr.mdc-list-divider
          .wrapper
            component(:is="currentClickedStudent.is" :sid="currentClickedStudent.sid" :uid="currentClickedStudent.uid" :name="currentClickedStudent.name")

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
  import mixinHal from '../mixins/hal';
  import mixinDom from '../mixins/dom';
  import _isEqual from 'lodash/isEqual';

  export default {
    name: 'students',
    mixins: [mixinHal, mixinDom],
    components: {
      'spinner': () => import('@/components/Spinner'),
      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': () => import('@/components/Header'),
      'form-rate-review': () => import('@/components/FormRateReview'),
      'empty': () => import('@/components/Empty')
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        sessions: null,
        currentAuth: null,
        snackbar: null,

        currentClickedStudent: {
          sid: "0",
          uid: "0",
          name: "",
          is: "empty",
        }
      }
    },
    methods: {
      onClickList(e) {
        const $el = e.target.closest('.mdc-list-item');
        let {sid, uid, name} = $el.dataset;
        let {sid: currentSid, uid: currentUid, name: currentName} = this.currentClickedStudent;

        const $formWrapper = $el.closest(".mdc-list").querySelector(".wrapper");
        this.insertAfter($formWrapper, $el.nextSibling);
        if (_isEqual({sid: currentSid, uid: currentUid, name: currentName}, {sid, uid, name})) {
          this.currentClickedStudent = {sid: "0", uid: "0", name: "", is: "empty"};
        } else {
          this.currentClickedStudent = {sid, uid, name, is: 'form-rate-review'};
        }
      },
      getStudentsSessions() {
        const _this = this;
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
//              console.log(sessions);
            }
            else {
              sessions = [];
            }
            this.sessions = sessions;

            let j = 0;
            const d = 200;
            const currentView = [];
            if (!this.sessions.length) {
              return;
            }
            this.sessions.forEach((v, i, a) => {
              j++;
              setTimeout(() => this.parseEmbedded('class', v._links.class.href, a[i]['_embedded'], (item) => {
                if (!item['_embedded']) {
                  this.$set(item, '_embedded', {
                    module: {name: "..."},
                    branch: {name: "..."},
                  });
                }
                j++;
                setTimeout(() => this.parseEmbedded('branch', item._links.branch.href, item['_embedded']), j * d);
                j++;
                setTimeout(() => this.parseEmbedded('module', item._links.module.href, item['_embedded']), j * d);
                return item
              }), j * d);

              j++;
              setTimeout(() => {
                if (!!v._embedded.items) {

                  v._embedded.items.forEach((v2, i2, a2) => {
                    if (!a2[i2]['_embedded']) {
                      this.$set(a2[i2], '_embedded', {});
                    }
                    this.parseEmbedded('student', v2._links.student.href, a2[i2]['_embedded'], item => {
                      item.photo = item.photo ? item.photo : "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=";
                      if (item.photo.indexOf('data:image/gif') < 0) {
                        let image = item.photo;
                        image = image.replace('https://', '').replace('http://', '');
                        image = `//images.weserv.nl/?output=jpg&il&q=100&w=96&h=96&t=square&url=${image}`;
                        item.photo = image;
                      }
                      return item;
                    });
                  })
                }
              }, j * d);

              setTimeout(() => {
                Array.from(document.querySelectorAll(".mdc-list-item")).forEach(v => {
                  const hammertime = new Hammer(v, {});
                  hammertime
                    .on('panend', e => {
                      const $el = e.target.closest(".mdc-list-item");
                      let {sid, uid, name} = $el.dataset;
                      _this.currentClickedStudent = {sid, uid, name};

                      if (Math.abs(e.deltaX) > e.target.closest('.mdc-list').offsetWidth * (1 / 3)) {
                        const url = `${process.env.API}/sessions/${$el.dataset.sid}/students/${$el.dataset.uid}`;

                        axios.post(url, {
                          interaction: 0,
                          creativity: 0,
                          cognition: 0,
                          review: "",
                          status: 0,
                        })
                          .then(response => {
                            _this.$bus.$emit('onAfterSubmitRateReview', response.data);
                          })
                          .catch(error => {
                            console.log(error);
                            _this.$bus.$emit('onAfterSubmitRateReview', {});
                          });
                      } else {
                        $el.style.marginLeft = 0;
                      }
                    })
                    .on('panleft panright', e => {
                      e.target.closest(".mdc-list-item").style.marginLeft = `${e.deltaX}px`;
                    })
                    .on('tap', e => this.onClickList(e));
                });
              }, 1500);
            });
//            }
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
      });
      this.$bus.$on('onAfterSubmitRateReview', (resp) => {
        Array.from(document.querySelectorAll(".mdc-list-item")).forEach(v => v.style.marginLeft = 0);
//        console.log(resp, this.currentClickedStudent);
        let {sid, uid, name} = this.currentClickedStudent;
        this.currentClickedStudent = {sid: "0", uid: "0", name: "...", is: "empty"};
        this.snackbar.show({
          message: `${name.split(" ")[0].toUpperCase()} has been saved`,
          actionText: 'Undo',
          actionHandler: function () {
            const url = `${process.env.API}/sessions/${sid}/students/${uid}`;

            axios.delete(url)
              .then(response => {
                _this.getStudentsSessions();
              })
              .catch(error => console.log(error));
          }
        });
        this.getStudentsSessions();
      });

      window.mdc.autoInit();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../assets/shared";
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
  }

  .mdc-list-item {
    padding: .5rem 1rem;
    background-color: #fff;
    z-index: 2;
    max-width: 100%;
    min-width: 90%;
  }

  .mdc-list-divider {
    border-bottom-color: #dcdcdc;
  }

  .mdc-list-item__text {
    text-transform: capitalize;
  }

  .mdc-list-item__graphic {
    &, img {
      width: 40px;
      height: 40px;
      border-radius: 50%;
    }
  }

  span.module {
    text-transform: uppercase;
  }

  span.branch, span.day-time {
    text-transform: capitalize;
  }

  span.day-time {
    float: right;
  }

</style>
