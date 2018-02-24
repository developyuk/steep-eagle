<template lang="pug">
  #students.mdc-typography
    header1
    .mdc-list-group
      template(v-for="(v,i) in sessions")
        h3.mdc-list-group__subheader
          span.module {{v.module.name}}
          | &nbsp;&nbsp;
          span.branch {{v.branch.name}}
        ul.mdc-list
          template(v-for="(vv,ii) in v.students")
            li.mdc-list-item
              .mdc-list-item__graphic
                img(':src'="vv.photo")
              span.mdc-list-item__text {{vv.name}}
                //span.mdc-list-item__secondary-text {{v._embedded.class.day}} {{v._embedded.class.time}}
            hr.mdc-list-divider
    .action
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
          const url = `${process.env.API}/tutors/${auth.id}/sessions?students=null`;

          axios.get(url)
            .then(response => {
              this.sessions = response.data._embedded;
              this.sessions.forEach((v, i) => {
                this.$set(this.sessions[i], 'branch', {'name': ''});
                this.$set(this.sessions[i], 'module', {'name': ''});
                this.$set(this.sessions[i], 'students', []);

                axios.get(`${process.env.API}${v._embedded.class._links.branch.href}`)
                  .then(response2 => {
                    this.$set(this.sessions[i], 'branch', response2.data);
                  })
                  .catch(error => console.log(error));

                axios.get(`${process.env.API}${v._embedded.class._links.module.href}`)
                  .then(response2 => {
                    this.$set(this.sessions[i], 'module', response2.data);
                  })
                  .catch(error => console.log(error));

                if (v._links.students) {
                  v._links.students.forEach((v2, i2) => {
                    axios.get(`${process.env.API}${v2.href}`)
                      .then(response2 => {
                        const ii = this.sessions[i].students.length;
                        this.$set(this.sessions[i]['students'], ii, response2.data);

//                      console.log(this.sessions);
                      })
                      .catch(error => console.log(error))
                  })
                }
              });

            })
            .catch(error => console.log(error))

        });
      },
      getPosition(el) {
        let xPos = 0;
        let yPos = 0;

        while (el) {
          if (el.tagName == "BODY") {
            // deal with browser quirks with body/window/document and page scroll
            const xScroll = el.scrollLeft || document.documentElement.scrollLeft;
            const yScroll = el.scrollTop || document.documentElement.scrollTop;

            xPos += (el.offsetLeft - xScroll + el.clientLeft);
            yPos += (el.offsetTop - yScroll + el.clientTop);
          } else {
            // for all other non-BODY elements
            xPos += (el.offsetLeft - el.scrollLeft + el.clientLeft);
            yPos += (el.offsetTop - el.scrollTop + el.clientTop);
          }

          el = el.offsetParent;
        }
        return {
          x: xPos,
          y: yPos
        };
      }

    },
    mounted() {
      this.getStudentsSessions();
      this.$bus.$on('currentAuth', (auth) => {
        console.log(auth);
        this.currentAuth = auth;
      });

      const cls = this;
      setTimeout(() => {
        const $action = document.querySelector('.action');
        Array.from(document.querySelectorAll('#students .mdc-list')).forEach((v, i) => {
          Array.from(v.querySelectorAll('.mdc-list-item')).forEach((v2, i2) => {

            const mc = new Hammer(v2);
            let actionStatus = false;

            mc.on("panleft panright panend", function (ev) {
              const panLimit = ev.target.offsetWidth / 3;
              if (ev.type === 'panright' || ev.type === 'panleft') {

                if (ev.deltaX > 0) {
                  $action.className = 'action action--right';
                  actionStatus = false;
                  if (ev.deltaX > panLimit) {
                    actionStatus = true;
                  }
                }
                if (ev.deltaX < 0) {
                  $action.className = 'action action--left';
                  actionStatus = false;
                  if (ev.deltaX < panLimit * -1) {
                    actionStatus = true;
                  }
                }
                v2.style.transform = `translateX(${ev.deltaX}px)`;
                $action.style.top = `${cls.getPosition(v2).y}px`;
                $action.style.height = "4rem";
              }
              if (ev.type === 'panend') {
                v2.style.transform = `translateX(0px)`;

                if (Math.abs(ev.deltaX) > panLimit) {
                  const session = cls.sessions[i - 1];
                  cls.$delete(session.students, i2);
                  if (!session.students.length) {
                    cls.$delete(cls.sessions, i - 1);
                  }
                  $action.style.height = "0";
                  console.log('lewat center', i - 1, i2);
                  actionStatus = false;
                }
              }
            });
          });
        });
        console.log('add hammer event');
      }, 999);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #students {
    position: relative;
    height: 100vh;
  }

  .mdc-list-group {
    height: calc(100vh - 8rem);
    overflow: scroll;
  }

  .mdc-list-item {
    padding: .5rem 1rem;
    background-color: #fff;
    z-index: 2;
  }

  .mdc-list-item__text {
    text-transform: capitalize;
  }

  span.module {
    text-transform: uppercase;
  }

  span.branch {
    text-transform: capitalize;
  }

  .action {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 4rem;
    z-index: 1;
  }

  .action--right {;
    background-color: #1FEEB2;
  }

  .action--left {
    background-color: #F29E4C
  }

</style>
