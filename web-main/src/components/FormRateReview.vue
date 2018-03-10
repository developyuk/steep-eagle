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

  export default {
    name: 'form-rate-review',
    props: ['sid', 'uid', 'name'],
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        ratingInteraction: 0,
        ratingCreativity: 0,
        ratingCognition: 0,
        review: '',
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
        $rating.querySelectorAll(`.material-icons`).forEach(v => v.classList.remove('is-active'));
        [...Array(parseInt(value)).keys()].forEach(v => {
          $rating.querySelector(`.material-icons[data-value='${v + 1}']`).classList.add('is-active')
        });
      },
      submit() {
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;
        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: true,
        })
          .then(response => {
            this.$bus.$emit('onAfterSubmitRateReview', response.data);
          })
          .catch(error => console.log(error))
      },
      absence() {
        const url = `${process.env.API}/sessions/${this.sid}/students/${this.uid}`;

        axios.post(url, {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition),
          review: this.review,
          status: false,
        })
          .then(response => {
            this.$bus.$emit('onAfterSubmitRateReview', response.data);
          })
          .catch(error => console.log(error));
        console.log(this.sid, this.uid);
      }
    },
    mounted() {
      mdc.ripple.MDCRipple.attachTo(document.querySelector('button'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  $mdc-theme-primary: #ED235C;
  #form-rate-review {
    padding: 0rem 2rem;
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
      .name {
        text-transform: capitalize;
      }
    }
  }

  .submit, .absence {
    width: 100%;
    text-align: center;
  }

  .submit button {
    width: 100%;
    background-color: #1FEEB2;
    margin: 1rem 0;
  }

  .absence a {
    color: #EB5757;
    margin-bottom: 1rem;
  }

  .mdc-text-field--textarea .mdc-text-field__input {
    padding-top: 16px;
  }
</style>
