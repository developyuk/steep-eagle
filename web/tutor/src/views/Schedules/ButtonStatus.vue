<template lang="pug">
  #buttonStatus
    button(v-if="status === 'start'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    button.start-other(v-if="status === 'start-ongoing'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'disabled'" disabled @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Start
    button(v-if="status === 'late'" @click='start($event)').mdc-button.mdc-button--raised.mdc-button--compact Activate
</template>

<script>
import moment from "moment";
import { mapState } from "vuex";
import { MDCRipple } from "@material/ripple";

export default {
  props: ["item", "lastAttendances"],
  data() {
    return {};
  },
  computed: {
    ...mapState(["currentAuth"]),
    status() {
      let status = "disabled";
      if (!this.item.class._id) {
        return status;
      }

      const msts = moment(this.item.nextStart);
      const mfts = moment(this.item.nextFinish);
      const mnow = moment();

      const lsItems = this.lastAttendances._items;
      // console.log(msts,mfts,mnow);

      if (!!lsItems.length && msts.diff(mnow, "days") < 1) {
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
          // console.log(mls, mfts, mls.isBefore(mfts));
          if (mls.isBefore(mfts)) {
            status = "start-ongoing";
          }

          return status;
        }
      } else {
        // console.log(msts.toISOString(), mfts.toISOString(), mnow.toISOString());
        // console.log(msts.diff(mnow, "minutes"),mfts.diff(mnow, "minutes"));
        if (msts.diff(mnow, "minutes") < 5 && mfts.diff(mnow, "minutes") > 0) {
          status = "start";
        }
        if (mnow.isAfter(mfts)) {
          status = "late";
        }
      }

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

.mdc-button {
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

.start-other {
  text-transform: uppercase;
  background-color: map-get($palettes, blue);
  color: #fff;
  line-height: 2.25rem;
  text-align: center;
}
</style>
