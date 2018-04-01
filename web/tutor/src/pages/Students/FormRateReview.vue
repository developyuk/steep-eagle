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

  export default {
    name: 'form-rate-review',
    props: ['sid', 'uid', 'name', 'index'],
    data() {
      return {
        ratingInteraction: 0,
        ratingCreativity: 0,
        ratingCognition: 0,
        review: '',
      }
    },
    watch: {
      review(val){
        this.$bus.$emit('onClickRating', {
          sid: this.sid,
          uid: this.uid,
          name: this.name,
          form: {
            interaction: parseInt(this.ratingInteraction),
            creativity: parseInt(this.ratingCreativity),
            cognition: parseInt(this.ratingCognition),
            review: this.review,
          },
        });
      }
    },
    methods: {
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
        this.$bus.$emit('onClickRating', {
          sid: this.sid,
          uid: this.uid,
          name: this.name,
          form: {
            interaction: parseInt(this.ratingInteraction),
            creativity: parseInt(this.ratingCreativity),
            cognition: parseInt(this.ratingCognition),
            review: this.review,
          },
        });
      },
      submit() {
        const $el = this.$el.closest('li');
        $el.classList.add('animated', `slideOutUpHeight`);
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;
        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: true,
        })
          .catch(error => {
            this.$bus.$emit('onUndoRateReview', {
              sid: this.sid,
              uid: this.uid,
              name: this.name,
            });
            console.log(error)
          })
      },
      absence() {
        const $el = this.$el.closest('li');
        $el.classList.add('animated', `slideOutUpHeight`);
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;

        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: false,
        })
          .catch(error => {
            console.log(error);
            this.$bus.$emit('onUndoRateReview', {
              sid: this.sid,
              uid: this.uid,
              name: this.name,
            });
          });
      }
    },
    mounted() {
      mdc.ripple.MDCRipple.attachTo(document.querySelector('button'));
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
