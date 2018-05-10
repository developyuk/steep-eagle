<template lang="pug">
  li#card(:data-index="index" :data-sid="sid" :data-student="student" :data-isActive="isActive")
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
  import {getCorrectEventName} from '@material/animation';
  import {mapState} from 'vuex';

  export default {
    name: 'card',
    components: {
      'form-rate-review': () => import('./FormRateReview'),
      'empty': () => import('./Empty'),
    },
    computed: {
      ...mapState(['currentMqtt']),
    },
    props: ['index', 'sid', 'student', 'isActive'],
    watch: {
      isActive(v) {
        this.currentComponent = v ? 'form-rate-review' : 'empty'
      }
    },
    data() {
      return {
        currentComponent: 'empty',
        direction: null,
        hammertime: null
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

      this.$el.addEventListener(getCorrectEventName(window, 'animationend'), e => {
        console.log(e.animationName);
        if (['slideOutRightHeight', 'slideOutLeftHeight', 'slideOutUpHeight'].indexOf(e.animationName.split('-')[0]) >= 0) {
          this.currentMqtt.mqtt
            .publish(this.currentMqtt.topic, JSON.stringify({
              sid: this.sid,
              uid: this.student.id,
              name: this.student.name,
              on: "successRateReview",
            }));
        }
      });

      this.hammertime = new Hammer($el, {});
      this.hammertime
        .on('panend', e => {
          if (Math.abs(e.deltaX) > this.$el.closest('.mdc-list').offsetWidth * (1 / 3)) {
            this.$el.classList.add('animated', `slideOut${this.direction}Height`);
            const path = `/sessions/${this.sid}/students/${this.student.id}`;

            axios.post(`${process.env.API}${path}`, {
              interaction: 0,
              creativity: 0,
              cognition: 0,
              review: "",
              status: 0,
            })
              .catch(error => {
                console.log(error);

                this.currentMqtt.mqtt
                  .publish(this.currentMqtt.topic, JSON.stringify({
                    sid: this.sid,
                    uid: this.student.id,
                    name: this.student.name,
                    on: "undoRateReview",
                  }));
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
          this.currentMqtt.mqtt
            .publish(this.currentMqtt.topic, JSON.stringify({
              sid: this.sid,
              uid: this.student.id,
              name: this.student.name,
              on: "tapStudent",
            }));
        });
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
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
