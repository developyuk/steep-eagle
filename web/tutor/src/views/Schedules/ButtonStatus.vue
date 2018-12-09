<template lang="pug">
  #buttonStatus
    button(v-if="status === 'start'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    button.start-other(v-if="status === 'start-ongoing'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    //- button(v-if="status === 'start-late-ongoing'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Activated
    button(v-if="status === 'disabled'" disabled @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'late'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Activate
    //- span.ongoing(v-if="status ==='ongoing'") Ongoing
    //- span.late-ongoing(v-if="status ==='late-ongoing'") Activated
</template>

<script>
import moment from "moment";
import { mapState } from "vuex";
import { MDCRipple } from "@material/ripple";

export default {
  props: ["item", "index"],
  data() {
    return {};
  },
  computed: {
    ...mapState(["currentAuth"]),
    status() {
      const item = this.item;
      const msts = moment(item.start);
      const mfts = moment(item.finish);
      const mnow = moment();

      let status = "disabled";
      const ls = item.last_attendances._items;
      // console.log(item,item.last_attendances._items);

      if (!!ls.length && msts.diff(mnow, "days") < 1) {
        const lsItems = ls;
        const mls = moment(lsItems[0]._created);

        if (
          !!this.currentAuth &&
          !!lsItems.length &&
          !!lsItems[0].tutor &&
          !!lsItems[0].tutor._id &&
          !lsItems.filter(v => {
            return v.tutor._id === this.currentAuth._id;
          }).length >= 1
        ) {
          if (mls.isBefore(mfts)) {
            status = "start-ongoing";
          }
          // if (mls.isAfter(mfts)) {
          //   status = "start-late-ongoing";
          // }
          return status;
        }
        // else {
        //   if (mls.isBefore(mfts)) {
        //     status = "ongoing";
        //   }
        //   if (mls.isAfter(mfts)) {
        //     status = "late-ongoing";
        //   }
        // }
      } else {
        // console.log(item,item.start,item.finish,msts, mfts, mnow);
        if (msts.diff(mnow, "minutes") < 5 && mfts.diff(mnow, "minutes") > 0) {
          status = "start";
        }
        if (mnow.isAfter(mfts)) {
          status = "late";
        }
      }
      // console.log(status)

      return status;
    }
  },
  methods: {
    start(e) {
      this.$emit("input", this.item);
    }
  },
  mounted() {
    const $el = this.$el.querySelector(".mdc-button");
    if (!!$el) {
      new MDCRipple(this.$el.querySelector(".mdc-button"));
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";

.mdc-button,
// span.ongoing,
// span.late-ongoing
 {
  position: absolute;
  right: 1rem;
  font-size: 1rem * (1/1.75);
  background-color: map-get($palettes, green);
  width: 5rem;
  top: 1.5rem;
}

.mdc-button {
  &:disabled {
    background-color: #999999;
    color: #fff;
  }
}

// span.ongoing,
// span.late-ongoing,
.start-other {
  text-transform: uppercase;
  background-color: map-get($palettes, blue);
  color: #fff;
  line-height: 2.25rem;
  text-align: center;
}
</style>
