<template lang="pug">
  #students.mdc-typography
    header1
    .mdc-list-group
      template(v-for="(v,i) in sessions")
        h3.mdc-list-group__subheader
          span.module {{v._embedded.class._embedded.module.name}}
          | &nbsp;&nbsp;
          span.branch {{v._embedded.class._embedded.branch.name}}
          | &nbsp;&nbsp;
          span.day-time {{v._embedded.class.start_at}} - {{v._embedded.class.finish_at}}
        ul.mdc-list
          template(v-for="(vv,ii) in v._embedded.class._embedded.students")
            li.mdc-list-item
              .mdc-list-item__graphic
                img(':src'="vv.photo")
              span.mdc-list-item__text {{vv.name}}
                //span.mdc-list-item__secondary-text {{v._embedded.class.day}} {{v._embedded.class.time}}
            hr.mdc-list-divider
    tab-bottom
</template>

<script>
  import TabBottom from '@/components/TabBottom';
  import Header from '@/components/Header';
  import axios from 'axios';

  export default {
    name: 'students',
    components: {
      TabBottom,
      'header1': Header
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        students: [],
        sessions: [],
        currentAuth: null,
      }
    },
    methods: {
      getStudentsSessions() {

        this.$bus.$on('currentAuth', (auth) => {
          const url = `${process.env.API}/tutors/${auth.id}/sessions`;

          axios.get(url)
            .then(response => {
              this.sessions = response.data._embedded;
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
  #students {
    position: relative;
    height: 100vh;
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
