<template lang="pug">
  aside.mdc-drawer.mdc-drawer--modal
    .mdc-drawer__header
      slot(name="header")
    .mdc-drawer__content
      slot
</template>

<script>
import { MDCDrawer } from "@material/drawer";

export default {
  mounted() {
    const $drawer = new MDCDrawer(this.$el);
    this.$emit("mounted", $drawer);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";

.mdc-drawer {
  width: 90%;
  max-width: 24rem;
  background-color: map-get($palettes, red);
  color: #fff;

  .mdc-drawer__header {
    // display: block;
    /*padding-top: 3rem;*/
    > * {
    //   width: 320px;
    //   display: block;
    //   line-height: 1.25rem;
      color: #fff;
    }
    img {
      // width: 4rem;
      // height: 4rem;
      // margin-bottom: 1rem;
    }
  }

  .mdc-drawer__content {
    // height: 100%;
    > .mdc-card {
      height: calc(100% - 9rem);
      background-color: transparent;
      margin-top: 5px;
      box-shadow: 0px -2px 4px -1px rgba(0, 0, 0, 0), 0px -4px 5px 0px rgba(0, 0, 0, 0.1), 0px -1px 0 0px rgba(0, 0, 0, 0.1);
    }
  }

}
</style>
