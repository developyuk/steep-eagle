<template lang="pug">
  li#card
    .mdc-list-item
      .mdc-list-item__graphic
        my-img(:src="student.photo")
      span.mdc-list-item__text
        placeholder(:value="student.name")
    hr.mdc-list-divider(v-if="isActive")

    //transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
    component(:is="currentComponent" :sid="sid" :tid="tid" :uid="student._id" :name="student.name" v-model="currentAttendance")
</template>

<script>
import axios from "axios";
import _debounce from "lodash/debounce";
import { getCorrectEventName } from "@material/animation";
import { mapState } from "vuex";
import Hammer from "hammerjs";
import MyImg from "@/components/Img";
import FormRateReview from "./FormRateReview";
import Placeholder from "@/components/Placeholder";

export default {
  components: { FormRateReview, MyImg, Placeholder },
  props: ["student", "isActive", "index", "sid", "tid"],
  computed: {
    ...mapState(["currentAuth", "currentMqtt"])
  },
  data() {
    return {
      currentComponent: "",
      currentAttendance: {},
      direction: null,
      hammertime: null
    };
  },
  watch: {
    isActive(v) {
      this.currentComponent = v ? "form-rate-review" : "";
    },
    currentAttendance(v, ov) {
      this.$emit("input", v);
    }
  },
  methods: {
  },
  mounted() {
    const $el = this.$el.querySelector(".mdc-list-item");

    this.hammertime = new Hammer($el, {});
    this.hammertime.get("pan").set({ direction: Hammer.DIRECTION_HORIZONTAL });
    this.hammertime.on("tap", e => {
      this.$emit("tap-student", this.index);
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";
/*@import "@material/animation/functions";*/

#card {
  // max-width: 100%;
  // min-width: 100%;
}

.mdc-list-item {
  background-color: #fff;
  // min-width: 100%;
  // box-sizing: border-box;
  // height: 4rem;
}
$size: 40px;
.mdc-list-item__text {
  text-transform: capitalize;
  line-height: $size;
}

.mdc-list-item__graphic {
  &,
  img {
    width: $size;
    height: $size;
    border-radius: 50%;
  }
}
</style>
