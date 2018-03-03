<template lang="pug">
  #students.mdc-typography
    header1
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
            component(:is="currentView[`${v.id}`][`${vv.users.id}`]" :sid="v.id" :uid="vv.users.id"  :name="vv.name")
    tab-bottom
</template>

<script>
    import TabBottom from '@/components/TabBottom';
    import Header from '@/components/Header';
    import FormRateReview from '@/components/FormRateReview';
    import Empty from '@/components/Empty';
  import axios from 'axios';

  export default {
    name: 'students',
    components: {
      TabBottom,
//      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': Header,
//      'header1': () => import('@/components/Header'),
      'form-rate-review': FormRateReview,
      'empty': Empty
//      'form-rate-review': () => import('@/components/FormRateReview'),
//      'empty': () => import('@/components/Empty')
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        students: [],
        sessions: [],
        currentAuth: null,
        currentView: {},

        lastId: ''
      }
    },
    methods: {
      onClickList(e) {
        const $el = e.target.nextSibling.nextSibling;
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
        if(is === 'empty') {
          this.$set(this.currentView[sid], uid, 'form-rate-review');
          this.lastId = `${sid}-${uid}`
        }
        console.log('clicked', sid, uid,this.lastId);
//        else {
//        }
//        this.currentView[sid][uid] = 'form-rate-review';
      },
      getStudentsSessions() {

        this.$bus.$on('currentAuth', (auth) => {
          const url = `${process.env.API}/tutors/${auth.id}/sessions`;

          axios.get(url)
            .then(response => {
              this.sessions = response.data._embedded;
              const currentView = [];
              this.sessions.forEach(v => {
                v._embedded.class._embedded.students.forEach(v2 => {
//                  this.$set(this.currentView, `x${v.id}`, {});
                  if (!currentView[v.id]) {
                    currentView[v.id] = [];
                  }
                  currentView[v.id][v2.users.id] = 'empty';
//                  this.$set(this.currentView[`x${v.id}`], `x${v2.users.id}`, 'empty');
//                  this.currentView[v.id] = [];
//                  this.currentView[v.id][v2.users.id] = 'empty'
                });

              });
              console.log(currentView);
              this.currentView = currentView;
            })
            .catch(error => console.log(error))

        });
      }
    },
    mounted() {
      this.getStudentsSessions();
      this.$bus.$on('currentAuth', (auth) => {
//        console.log(auth);
        this.currentAuth = auth;
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
