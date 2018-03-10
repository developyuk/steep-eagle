<template lang="pug">
  #students.mdc-typography
    header1
    spinner(v-if="!sessions")
    .empty(v-if="!!sessions && !sessions.length")
      img(src="https://images.weserv.nl/?w=300&url=dl.dropboxusercontent.com/s/6wrsi1smelwdhj7/face-2.png")
      p It’s holiday time, man! #[br]Pls enjoy yout time by being away from your phone.
    .mdc-list-group
      template(v-for="(v,i) in sessions")
        h3.mdc-list-group__subheader
          span.module {{v._embedded.class._embedded.module.name}}
          | &nbsp;-&nbsp;
          span.branch {{v._embedded.class._embedded.branch.name}}
          | &nbsp;&nbsp;
          span.day-time {{v._embedded.class.start_at}} - {{v._embedded.class.finish_at}}
        ul.mdc-list
          template(v-for="(vv,ii) in v._embedded.class._embedded.students")
            li.mdc-list-item(@click="onClickList($event)")
              .mdc-list-item__graphic
                img(':src'="vv.photo")
              span.mdc-list-item__text {{vv.name}}
                //span.mdc-list-item__secondary-text {{v._embedded.class.day}} {{v._embedded.class.time}}
            hr.mdc-list-divider
            component(:is="currentView[`${v.id}`][`${vv.user_id}`]" :sid="v.id" :uid="vv.user_id"  :name="vv.name")
    tab-bottom
</template>

<script>
  import axios from 'axios';

  export default {
    name: 'students',
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
        students: [],
        sessions: '',
        currentAuth: null,
        currentView: {},

        lastId: ''
      }
    },
    methods: {
      onClickList(e) {
        const $el = e.target.closest('.mdc-list-item').nextSibling.nextSibling;
        console.log('clicked', $el);
        let sid;
        let uid;
        const is = $el.getAttribute('id');
        if (this.lastId) {
          const lastId = this.lastId.split('-');
          sid = lastId[0];
          uid = lastId[1];
          this.$set(this.currentView[sid], uid, 'empty');
        }
        sid = $el.getAttribute('sid');
        uid = $el.getAttribute('uid');
        if (is === 'empty') {
          this.$set(this.currentView[sid], uid, 'form-rate-review');
          this.lastId = `${sid}-${uid}`
        }
        console.log('clicked', sid, uid, this.lastId);
//        else {
//        }
//        this.currentView[sid][uid] = 'form-rate-review';
      },
      getStudentsSessions() {
        const url = `${process.env.API}/tutors/${this.currentAuth.id}/sessions`;
        axios.get(url)
          .then(response => {
            this.sessions = response.data._embedded;
            const currentView = [];
            this.sessions.forEach((v, i, a) => {
              v['_embedded']['class']['_embedded']['students'].forEach((v2, i2, a2) => {
                console.log(v2);
                if (!currentView[v.id]) {
                  currentView[v.id] = [];
                }
                currentView[v.id][v2.user_id] = 'empty';

                let image = !!v2['photo'] ? v2['photo'] : 'https://image.flaticon.com/icons/png/128/201/201818.png';
                image = image.replace('https://', '').replace('http://', '');
                image = `//images.weserv.nl/?output=png&il&q=100&w=96&h=96&t=square&url=${image}`;
                this.$set(a[i]['_embedded']['class']['_embedded']['students'][i2], 'photo', image);
              });

            });
            this.currentView = currentView;
          })
          .catch(error => console.log(error));


      }
    },
    mounted() {
      this.$bus.$on('currentAuth', (auth) => {
//        console.log(auth);
        this.currentAuth = auth;
        this.getStudentsSessions();
      });
      this.$bus.$on('onAfterSubmitRateReview', (resp) => {
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

  .mdc-list-item {
    padding: .5rem 1rem;
    background-color: #fff;
    z-index: 2;
  }

  .mdc-list-item__text {
    text-transform: capitalize;
  }

  .mdc-list-item__graphic {
    &, img {
      width: 40px;
      height: 40px;
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
