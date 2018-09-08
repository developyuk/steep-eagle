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
            rating(name="rating_cognition" @change="onChangeCognition")
        .creativity.clearfix
          span.title creativity
          span.stars
            rating(name="rating_creativity" @change="onChangeCreativity")
        .interaction.clearfix
          span.title interaction
          span.stars
            rating(name="rating_interaction" @change="onChangeInteraction")
      .review
        h4.title Comment for
          span.name  {{name}}
        .mdc-text-field.mdc-text-field--textarea.mdc-text-field--fullwidth
          textarea(v-model="review").mdc-text-field__input
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

export default {
  props: ["stid", "uid", "name", "sid", "tid"],
  components: {
    rating: () => import("./Rating")
  },
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
    onChangeCognition(e) {
      this.ratingCognition = e;
    },
    onChangeCreativity(e) {
      this.ratingCreativity = e;
    },
    onChangeInteraction(e) {
      this.ratingInteraction = e;
    },
    submit() {
      //        const url = `${process.env.VUE_APP_API}/attendances/${this.stid}/students/${this.uid}`;
      const url = `${process.env.VUE_APP_API}/attendances_students`;
      const data = {
        // attendance_tutor: this.stid,
        attendance: this.sid,
        tutor: this.tid,
        student: this.uid,
        rating_interaction: parseInt(this.ratingInteraction),
        rating_creativity: parseInt(this.ratingCreativity),
        rating_cognition: parseInt(this.ratingCognition),
        feedback: this.review,
        is_presence: true
      };

      axios
        .post(url, data)
        // .then(response => { })
        .catch(error => {
          console.log(error);
        });
    },
    absence() {
      const url = `${process.env.VUE_APP_API}/attendances_students`;
      const data = {
        // attendance_tutor: this.stid,
        attendance: this.sid,
        tutor: this.tid,
        student: this.uid,
        rating_interaction: parseInt(this.ratingInteraction),
        rating_creativity: parseInt(this.ratingCreativity),
        rating_cognition: parseInt(this.ratingCognition),
        feedback: this.review,
        is_presence: false
      };
      console.log(data);

      axios
        .post(url, data)
        // .then(response => { })
        .catch(error => {
          console.log(error);
        });
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
