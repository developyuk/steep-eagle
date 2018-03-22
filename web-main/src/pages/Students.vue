<template lang="pug">
  #students.mdc-typography
    header1
    spinner(v-if="!sessions")
    .empty(v-if="!!sessions && !sessions.length")
      img(src="https://images.weserv.nl/?w=300&url=dl.dropboxusercontent.com/s/6wrsi1smelwdhj7/face-2.png")
      p It’s holiday time, man! #[br]Pls enjoy yout time by being away from your phone.
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
            li.mdc-list-item(:data-sid="vv.id" :data-uid="vv._embedded.student.id")
              .mdc-list-item__graphic
                img(':src'="vv._embedded.student.photo")
              span.mdc-list-item__text {{vv._embedded.student.name}}
            hr.mdc-list-divider
            .wrapper(v-if="!!vv._embedded.student && !!currentView && !!currentView[vv.id] && !!currentView[vv.id][vv._embedded.student.id]")
              component(:is="currentView[vv.id][vv._embedded.student.id]" :sid="vv.id" :uid="vv._embedded.student.id" :name="vv._embedded.student.name")
    tab-bottom
</template>

<script>
  import axios from 'axios';
  import moment from 'moment';
  import Hammer from 'hammerjs';
  import sharedHal from '../mixins/hal';

  export default {
    name: 'students',
    mixins: [sharedHal],
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
        currentView: {},

        lastId: ''
      }
    },
    methods: {
      onClickList(e) {
        const $el = e.target.closest('.mdc-list-item').nextSibling.nextSibling.childNodes[0];
        let sid;
        let uid;
        const is = $el.getAttribute('id');
//        console.log(this.lastId);
        if (!!this.lastId) {
          const lastId = this.lastId.split('-');
          sid = lastId[0];
          uid = lastId[1];
          this.$set(this.currentView[sid], uid, 'empty');
        }
        sid = $el.getAttribute('sid');
        uid = $el.getAttribute('uid');
        console.log(sid, uid);
        if (is === 'empty') {
          this.$set(this.currentView[sid], uid, 'form-rate-review');
          this.lastId = `${sid}-${uid}`;
        }
        console.log(!!this.currentView, !!this.currentView[sid], !!this.currentView[sid][uid], JSON.stringify(this.currentView));
//        console.log('clicked', sid, uid, this.lastId);
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
//              console.log(sessions);
            }
            else {
              sessions = [];
            }
            this.sessions = sessions;

            let j = 0;
            const d = 200;
            const currentView = [];
            if (this.sessions) {
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
                        if (!currentView[v2.id]) {
                          currentView[v2.id] = [];
                        }
                        currentView[v2.id][item.id] = 'empty';
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
            }
            this.currentView = currentView;
          })
          .catch(error => console.log(error));


      }
    },
    mounted() {
      this.$bus.$on('currentAuth', (auth) => {
        this.currentAuth = auth;
        this.getStudentsSessions();
        setTimeout(() => {
          Array.from(document.querySelectorAll(".mdc-list-item")).forEach(v => {
            var hammertime = new Hammer(v, {});
            hammertime
              .on('panend', e => {
                const $el = e.target.closest(".mdc-list-item");
                if (Math.abs(e.deltaX) > e.target.closest('.mdc-list').offsetWidth * (1 / 3)) {
                  const url = `${process.env.API}/sessions/${$el.dataset.sid}/students/${$el.dataset.uid}`;

                  axios.post(url, {
                    interaction: 0,
                    creativity: 0,
                    cognition: 0,
                    review: "",
                    status: false,
                  })
                    .then(response => {
                      this.$bus.$emit('onAfterSubmitRateReview', response.data);
                    })
                    .catch(error => console.log(error));
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
      this.$bus.$on('onAfterSubmitRateReview', (resp) => {
        Array.from(document.querySelectorAll(".mdc-list-item")).forEach(v => v.style.marginLeft = 0);
        this.getStudentsSessions();
      });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
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
  }

  .mdc-list {
    background-color: #F29E4C;
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
