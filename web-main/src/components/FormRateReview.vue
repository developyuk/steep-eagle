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
              select(v-model="ratingCognition")
                option(v-for="v in [0,1,2,3,4,5]" :value=v) {{v}}
              i.material-icons(v-for="v in [1,2,3,4,5]")
        .creativity.clearfix
          span.title creativity
          span.stars
            .rating
              select(v-model="ratingCreativity").hide
                option(v-for="v in [0,1,2,3,4,5]" :value=v) {{v}}
              i.material-icons(v-for="v in [1,2,3,4,5]")
        .imagination.clearfix
          span.title imagination
          span.stars
            .rating
              select(v-model="ratingImagination").hide
                option(v-for="v in [0,1,2,3,4,5]" :value=v) {{v}}
              i.material-icons(v-for="v in [1,2,3,4,5]")
      .review
        h4.title Comment for
          span.name  {{name}}
        .mdc-text-field.mdc-text-field--textarea.mdc-text-field--fullwidth
          textarea(v-model="review").mdc-text-field__input
      .submit: button(type='submit').mdc-button.mdc-button--raised submit
      .absence: a(href='#' @click.prevent="absence") or absence ?
</template>

<script>
  export default {
    name: 'form-rate-review',
    props: ['sid', 'uid', 'name'],
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        ratingImagination: 0,
        ratingCreativity: 0,
        ratingCognition: 0,
        review: '',
      }
    },
    methods: {
      submit() {
        console.log(this.sid, this.uid,
          parseInt(this.ratingImagination), parseInt(this.ratingCreativity), parseInt(this.ratingCognition),
          this.review);
      },
      absence() {
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
  #form-rate-review {
    padding: 0rem 2rem;
  }

  .rate {
    .cognition, .creativity, .imagination {
      text-transform: capitalize;
    }
    .title .title {
      float: left;
    }
    .stars {
      float: right;
    }
    .rating{
      unicode-bidi: bidi-override;
      direction: rtl;
    }
    i.material-icons {
      display: inline-block;
      position: relative;
      width: 1.1em;
      &:before {
        content: 'star_outline';
      }
      &:hover:before,
      &:hover ~ i:before{
        content: 'star'
      }
    }

  }

  .rate, .review {
    > .title {
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
</style>
