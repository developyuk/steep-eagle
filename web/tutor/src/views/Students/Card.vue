<template lang="pug">
  li#card(:data-index="index" :data-sid="stid" :data-student="student" :data-isActive="isActive")
    .mdc-list-item
      .mdc-list-item__graphic
        my-img(:src="student.photo")
      span.mdc-list-item__text
        placeholder(:value="student.name")
    hr.mdc-list-divider(v-if="isActive")

    //transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
    component(:is="currentComponent" :stid="stid" :sid="sid" :tid="tid" :uid="student.id" :name="student.name")
</template>

<script>
import axios from "axios";
import _debounce from "lodash/debounce";
import { getCorrectEventName } from "@material/animation";
import { mapState } from "vuex";
import Hammer from "hammerjs";

export default {
  components: {
    "form-rate-review": () => import("./FormRateReview"),
    "my-img": () => import("@/components/Img"),
    placeholder: () => import("@/components/Placeholder")
  },
  props: ["stid", "student", "isActive", "index", "sid", "tid"],
  computed: {
    ...mapState(["currentAuth", "currentMqtt"])
  },
  watch: {
    isActive(v) {
      this.currentComponent = v ? "form-rate-review" : "";
    }
  },
  data() {
    return {
      currentComponent: "",
      direction: null,
      hammertime: null
    };
  },
  methods: {
    // setPosition(v = 0) {
    //   this.$el.style.marginLeft = `${v}px`;
    // }
  },
  mounted() {
    const $el = this.$el.querySelector(".mdc-list-item");

    this.hammertime = new Hammer($el, {});
    this.hammertime.get("pan").set({ direction: Hammer.DIRECTION_HORIZONTAL });

    this.hammertime
      //        .on('panend', e => {
      //          if (Math.abs(e.deltaX) > this.$el.closest('.mdc-list').offsetWidth * (1 / 3)) {
      //
      //
      //            let url = `${process.env.VUE_APP_DBAPI}/attendances_students`;
      //            let params = {
      //              attendance_tutor: this.stid,
      //              student: this.student.id,
      //
      //              rating_interaction: 0,
      //              rating_creativity: 0,
      //              rating_cognition: 0,
      //              feedback: "",
      //              is_presence: false,
      //            };
      //            axios.post(url, params)
      //              .then(response => {
      //                console.log(response.data);
      //                this.currentMqtt.mqtt
      //                  .publish(this.currentMqtt.topic, JSON.stringify({
      //                    sid: this.stid,
      //                    sts: {
      //                      id: response.data.id,
      //                      et: response.data._etag,
      //                    },
      //                    uid: this.student.id,
      //                    name: this.student.name,
      //                    by: this.currentAuth,
      //                    on: "successRateReview",
      //                  }));
      //              })
      //              .catch(error => {
      //                console.log(error);
      //              });
      //          } else {
      //            this.setPosition();
      //          }
      //
      //        })
      //        .on('panleft panright', e => {
      //          console.log(e);
      //          this.setPosition(e.deltaX);
      //        })
      .on("tap", e => {
        this.$emit("tap-student", {
          sid: this.stid,
          uid: this.student.id,
          name: this.student.name
        });
      });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "../../assets/shared";
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
