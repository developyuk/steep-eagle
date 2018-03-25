<template lang="pug">
  li#card
    .mdc-list-item
      .mdc-list-item__graphic
        img(':src'="student.photo")
      span.mdc-list-item__text {{student.name}}
    hr.mdc-list-divider(v-if="isActive")

    component(:is="currentComponent" :sid="sid" :uid="student.id" :name="student.name" :index="index" class="")
</template>

<script>
  import axios from 'axios';
  import _debounce from 'lodash/debounce';
  import vue from '../../assets/vue';
  import {getCorrectEventName} from '@material/animation';

  export default {
    name: 'card',
    props: ['index', 'sid', 'student', 'isActive'],
    watch: {
      isActive(v) {
        this.currentComponent = v ? 'form-rate-review' : 'empty'
      }
    },
    components: {
      'form-rate-review': () => import('./FormRateReview'),
      'empty': () => import('./Empty'),
    },
    data() {
      return {
        currentComponent: 'empty',
        msg: 'Welcome to Your Vue.js PWA',
        direction: null
      }
    },
    methods: {
      setPosition(v = 0) {
        this.$el.style.marginLeft = `${v}px`;
      }
    },
    mounted() {
      console.log();
      const $el = this.$el.querySelector('.mdc-list-item');
//      const $form = this.$el.querySelector('.mdc-list-item').nextSibling.nextSibling;
      const hammertime = new Hammer($el, {});
      hammertime
        .on('panend', e => {
          if (Math.abs(e.deltaX) > this.$el.closest('.mdc-list').offsetWidth * (1 / 3)) {
            this.$el.classList.add('animated', `slideOut${this.direction}Height`);
            this.$el.addEventListener(getCorrectEventName(window, 'animationend'), e => {
              if (['slideOutRightHeight', 'slideOutLeftHeight'].indexOf(e.animationName) >= 0) {
                this.$bus.$emit('onSuccessRateReview', {
                  $el: this.$el,
                  sid: this.sid,
                  uid: this.student.id,
                  name: this.student.name
                });
              }
            });
            const url = `${process.env.API}/sessions/${this.sid}/students/${this.student.id}`;

            axios.post(url, {
              interaction: 0,
              creativity: 0,
              cognition: 0,
              review: "",
              status: 0,
            })
              .then(response => {
              })
              .catch(error => {
                console.log(error);

                this.$bus.$emit('onUndoRateReview', {
                  $el: this.$el,
                  sid: this.sid,
                  uid: this.student.id,
                  name: this.student.name
                });
              });
          } else {
            this.setPosition();
          }

        })
        .on('panleft panright', e => {
          this.setPosition(e.deltaX);
        })
        .on('panleft', e => this.direction = 'Left')
        .on('panright', e => this.direction = 'Right')
        .on('tap', e => {
          this.$bus.$emit('onTapStudent', this.index)
        });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">
  @import "../../assets/shared";
  @import "../../assets/animate";
  /*@import "@material/animation/functions";*/

  #card {
    max-width: 100%;
    min-width: 100%;
  }

  .mdc-list-item {
    background-color: #fff;
    min-width: 100%;

    box-sizing: border-box;
    height: 4rem;
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
</style>
