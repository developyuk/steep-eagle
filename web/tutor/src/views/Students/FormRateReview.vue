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
            rating(name="rating_cognition" v-model="ratingCognition")
        .creativity.clearfix
          span.title creativity
          span.stars
            rating(name="rating_creativity" v-model="ratingCreativity")
        .interaction.clearfix
          span.title interaction
          span.stars
            rating(name="rating_interaction" v-model="ratingInteraction")
      .review
        h4.title Comment for
          span.name  {{name}}
        .mdc-text-field.mdc-text-field--textarea.mdc-text-field--fullwidth.mdc-text-field--dense
          textarea(v-model="review" id="textarea").mdc-text-field__input

      .submit: button(type='submit').mdc-button.mdc-button--raised submit
      .absence: a(href="#" @click.prevent="absence") or absence ?
</template>

<script>
import axios from "axios";
import { getCorrectEventName } from "@material/animation";
import { mapState } from "vuex";
//  import {mapState, mapMutations} from 'vuex';
import { MDCRipple } from "@material/ripple";
import { MDCTextField } from "@material/textfield";
import Rating from "./Rating";

export default {
  props: ["uid", "name", "sid", "tid"],
  components: { Rating },
  computed: {
    ...mapState(["currentAuth", "currentMqtt"])
  },
  data() {
    return {
      ratingInteraction: 0,
      ratingCreativity: 0,
      ratingCognition: 0,
      review: ""
    };
  },
  methods: {
    submit() {
      const url = `${process.env.VUE_APP_API}/attendances/${this.sid}/students`;
      const data = {
        tutor: this.tid,
        student: this.uid,
        rating: {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition)
        },
        feedback: this.review,
        isPresence: true
      };

      this.$emit("input", { url, data });
    },
    absence() {
      const url = `${process.env.VUE_APP_API}/attendances/${this.sid}/students`;
      const data = {
        tutor: this.tid,
        student: this.uid,
        rating: {
          interaction: parseInt(this.ratingInteraction),
          creativity: parseInt(this.ratingCreativity),
          cognition: parseInt(this.ratingCognition)
        },
        feedback: this.review,
        isPresence: false
      };

      this.$emit("input", { url, data });
    }
  },
  mounted() {
    new MDCTextField(this.$el.querySelector(".mdc-text-field"));
    new MDCRipple(this.$el.querySelector(".mdc-button"));
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";

#form-rate-review {
  padding: 0 2rem;
  background-color: #fff;
  /*min-width: 90%;*/
  max-width: 100%;
}

.rate {
  .cognition,
  .creativity,
  .interaction {
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
}

.rate,
.review {
  > .title {
    color: #bdbdbd;
    font-size: 0.75rem;
    margin: 0;
    padding: 1rem 0;
    .name {
      text-transform: capitalize;
    }
  }
}

.submit,
.absence {
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
