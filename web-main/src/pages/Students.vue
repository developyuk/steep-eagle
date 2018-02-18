<template lang="pug">
  #students.mdc-typography
    header1
    ul.mdc-list
      li.mdc-list-item(v-for="(v,i) in students")
        .mdc-list-item__graphic
          img(':src'="v.photo")
        span.mdc-list-item__text(v-if="v.name") {{v.name}}
          span.mdc-list-item__secondary-text(v-if="v.class") {{v.class.branch.name}} {{v.class.day}} {{v.class.time}}
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
        students: []
      }
    },
    methods: {
      getStudentsSessions() {

        const url = `${process.env.API}/tutors/3/sessions?students=null`;

        axios.get(url)
          .then(response => {
            response.data._embedded.forEach(v => {
              v._links.students.forEach(v2 => {
                axios.get(`${process.env.API}${v2.href}`)
                  .then(response2 => {
                    const i = this.students.length;
                    this.$set(this.students, i, response2.data);

                    axios.get(`${process.env.API}${v._links.class.href}`)
                      .then(response3 => {
                        this.$set(this.students[i], 'class', response3.data);

                        axios.get(`${process.env.API}${response3.data._links.branch.href}`)
                          .then(response4 => {
                            this.$set(this.students[i]['class'], 'branch', response4.data);
                          })
                          .catch(error => console.log(error));
                      })
                      .catch(error => console.log(error));

                    console.log(this.students);
                  })
                  .catch(error => console.log(error))
              })
            });

          })
          .catch(error => console.log(error))
      }
    },
    mounted() {
      this.getStudentsSessions();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #students {
    position: relative;
    height: 100vh;
  }
</style>
