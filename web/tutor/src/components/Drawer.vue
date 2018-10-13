<template lang="pug">
  .my-drawer
    aside.mdc-drawer.mdc-drawer--modal
      .mdc-drawer__header
        slot(name="header")
      .mdc-drawer__content
        slot
    .mdc-drawer-scrim
</template>

<script>
import { MDCDrawer } from "@material/drawer";
import { MDCList } from "@material/list";

export default {
  mounted() {
    const $drawer = new MDCDrawer(this.$el.querySelector(".mdc-drawer"));
    new MDCList(this.$el.querySelector(".mdc-list"));
    this.$emit("mounted", $drawer);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";

.mdc-drawer {
  width: 85%;
  max-width: 24rem;
  background-color: map-get($palettes, red);
  color: #fff;

  .mdc-drawer__header {
    > * {
      color: #fff;
    }
  }
}
</style>
