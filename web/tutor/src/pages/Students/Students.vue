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
  import _isEqual from 'lodash/isEqual';
  import _cloneDeep from 'lodash/cloneDeep';

  export default {
    name: 'students',
    mixins: [mixinHal, mixinDom],
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
            let j = 0;
            const d = 200;

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
                setTimeout(() => this.parseEmbedded('branch', item._links.branch.href, item['_embedded']), j * d);
                setTimeout(() => this.parseEmbedded('module', item._links.module.href, item['_embedded']), j * d);
                return item
              }), j * d);

              setTimeout(() => {
                if (!!v._embedded.items) {
                  v._embedded.items.forEach((v2, i2, a2) => {
                    this.$set(a2[i2], 'isActive', false);
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

      this.$bus.$on('onTapStudent', (key) => {
        const [i, ii] = key.split('.');
        const item = _cloneDeep(this.sessions[i]._embedded.items[ii]);

        this.sessions.forEach((v, i, a) => {
          v._embedded.items.forEach((v2, i2, a2) => {
            this.$set(a2[i2], 'isActive', false)
          });
        });
        console.log(key, i, ii, item['isActive']);
        this.$set(this.sessions[i]._embedded.items[ii], 'isActive', !item['isActive']);
      });

      this.$bus.$on('onUndoRateReview', (v) => {
        let {sid, uid, name, $el} = v;
        $el.className = 'animated slideInDownHeight';
        $el.style.marginLeft = 0;
      });
      this.$bus.$on('onSuccessRateReview', (v) => {
//        console.log(resp, this.currentClickedStudent);
        let {sid, uid, name, $el, index: key} = v;
//        this.getStudentsSessions();
        if (!!key) {
          const [i, ii] = key.split('.');
          this.$set(this.sessions[i]._embedded.items[ii], 'isActive', false);
        }
        this.snackbar.show({
          message: `${name.split(" ")[0].toUpperCase()} has been saved`,
          actionText: 'Undo',
          actionHandler: function () {
            _this.$bus.$emit('onUndoRateReview', v);
            const url = `${process.env.API}/sessions/${sid}/students/${uid}`;

            axios.delete(url)
              .then(response => {
                _this.$bus.$emit('onUndoRateReview', v);
              })
              .catch(error => {
                console.log(error);
                _this.$bus.$emit('onSuccessRateReview', v);
              });
          }
        });
      });

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
  }

  span.branch, span.day-time {
    text-transform: capitalize;
  }

  span.day-time {
    float: right;
  }

</style>
