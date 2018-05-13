<template lang="pug">
  #form-rate-review
    form(@submit.prevent="submit")
      .rate
        h4.title How was
          span.name  {{name}}
          |  ?
        .cognition.clearfix
          span.title cognition
          span.stars
            .rating
              i.material-icons(v-for="v in [5,4,3,2,1]" :data-value="v" @click="onClickRating($event)")
        .creativity.clearfix
          span.title creativity
          span.stars
            .rating
              i.material-icons(v-for="v in [5,4,3,2,1]" :data-value="v" @click="onClickRating($event)")
        .interaction.clearfix
          span.title interaction
          span.stars
            .rating
              i.material-icons(v-for="v in [5,4,3,2,1]" :data-value="v" @click="onClickRating($event)")
      .review
        h4.title Comment for
          span.name  {{name}}
        .mdc-text-field.mdc-text-field--textarea.mdc-text-field--fullwidth
          textarea(v-model="review").mdc-text-field__input
      .submit: button(type='submit').mdc-button.mdc-button--raised submit
      .absence: a(href="#" @click.prevent="absence") or absence ?
</template>

<script>
  import axios from 'axios';
  import {getCorrectEventName} from '@material/animation';
  import {mapState, mapMutations} from 'vuex';
  import {MDCRipple} from '@material/ripple';
  import {MDCTextField} from '@material/textfield';

  export default {
    props: ['sid', 'uid', 'name', 'index'],
    computed: {
      ...mapState(['currentAuth', 'currentMqtt']),
    },
    data() {
      return {
        ratingInteraction: 0,
        ratingCreativity: 0,
        ratingCognition: 0,
        review: '',
      }
    },
    watch: {
      review(val) {
        this.nextStudentSession({
          sid: this.sid,
          uid: this.uid,
          name: this.name,
          form: {
            interaction: parseInt(this.ratingInteraction),
            creativity: parseInt(this.ratingCreativity),
            cognition: parseInt(this.ratingCognition),
            review: this.review,
          },
          on: "clickRating",
        })
      }
    },
    methods: {
      ...mapMutations(['nextStudentSession']),
      onClickRating(e) {
        const value = e.target.dataset.value;
        const ParentClassList = e.target.closest('.clearfix').classList;
        const $rating = e.target.closest('.rating');
        const isInteraction = ParentClassList.contains('interaction');
        const isCognition = ParentClassList.contains('cognition');
        const isCreativity = ParentClassList.contains('creativity');
        if (isInteraction) {
          this.ratingInteraction = value;
        }
        if (isCognition) {
          this.ratingCognition = value;
        }
        if (isCreativity) {
          this.ratingCreativity = value;
        }

        this.nextStudentSession({
          sid: this.sid,
          uid: this.uid,
          name: this.name,
          form: {
            interaction: parseInt(this.ratingInteraction),
            creativity: parseInt(this.ratingCreativity),
            cognition: parseInt(this.ratingCognition),
            review: this.review,
          },
          on: "clickRating",
        })
      },
      submit() {
        const $el = this.$el.closest('li');
        $el.classList.add('animated', `slideOutUpHeight`);
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;

        this.currentMqtt.mqtt
          .publish(this.currentMqtt.topic, JSON.stringify({
            sid: this.sid,
            uid: this.uid,
            name: this.name,
            by: this.currentAuth,
            on: "successRateReview",
          }));
        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: true,
        })
          .then(response => {
            console.log(response.data);
          })
          .catch(error => {
            this.currentMqtt.mqtt
              .publish(this.currentMqtt.topic, JSON.stringify({
                sid: this.sid,
                uid: this.uid,
                name: this.name,
                by: this.currentAuth,
                on: "undoRateReview",
              }));
            console.log(error)
          })
      },
      absence() {
        const $el = this.$el.closest('li');
        $el.classList.add('animated', `slideOutUpHeight`);
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;

        this.currentMqtt.mqtt
          .publish(this.currentMqtt.topic, JSON.stringify({
            sid: this.sid,
            uid: this.uid,
            name: this.name,
            by: this.currentAuth,
            on: "successRateReview",
          }));
        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: false,
        })
          .then(response => {
            console.log(response.data);
          })
          .catch(error => {
            console.log(error);
            this.currentMqtt.mqtt
              .publish(this.currentMqtt.topic, JSON.stringify({
                sid: this.sid,
                uid: this.uid,
                name: this.name,
                by: this.currentAuth,
                on: "undoRateReview",
              }));
          });
      }
    },
    mounted() {
      new MDCTextField(this.$el.querySelector('.mdc-text-field'));
      new MDCRipple(this.$el.querySelector('.mdc-button'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "../../assets/shared";
  @import "../../assets/animate";

  #form-rate-review {
    padding: 0 2rem;
    background-color: #fff;
    /*min-width: 90%;*/
    max-width: 100%;
  }

  .rate {
    .cognition, .creativity, .interaction {
      text-transform: capitalize;
      .title {
        font-weight: 500;
      }
    }
    .title .title {
      float: left;
    }
    .stars {
      float: right;
    }
    .rating {
      unicode-bidi: bidi-override;
      direction: rtl;
    }
    i.material-icons {
      display: inline-block;
      position: relative;
      width: 1.1em;
      color: $mdc-theme-primary;
      &.is-active {

        &:before {
          content: 'star';
        }
      }
      &:before {
        content: 'star_outline';
      }
      &:hover:before,
      &:hover ~ i:before {
        content: 'star'
      }
    }

  }

  .rate, .review {
    > .title {
      color: #BDBDBD;
      font-size: .75rem;
      margin: 0;
      padding: 1rem 0;
      .name {
        text-transform: capitalize;
      }
    }
  }

  .submit, .absence {
    width: 100%;
    text-align: center;
    margin-bottom: 1rem;
    display: inline-block;
  }

  .submit button {
    width: 100%;
    background-color: map-get($palettes, green);
    margin: 1rem 0;
  }

  .absence a {
    color: map-get($palettes, orange);
  }

  .mdc-text-field--textarea .mdc-text-field__input {
    padding-top: 16px;
  }
</style>
