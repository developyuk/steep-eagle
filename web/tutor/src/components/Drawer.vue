<template lang="pug">
  aside.mdc-drawer.mdc-drawer--temporary.mdc-typography
    nav.mdc-drawer__drawer
      header.mdc-drawer__header
        .mdc-drawer__header-content
          slot(name="header")
      .mdc-drawer__content
        .mdc-card.mdc-elevation--z4
          ul.mdc-list
            slot
</template>

<script>
import { MDCTemporaryDrawer } from "@material/drawer";

export default {
  mounted() {
    const $drawer = new MDCTemporaryDrawer(this.$el);
    this.$emit("mounted", $drawer);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "../assets/shared";

.mdc-drawer {
  .mdc-drawer__drawer {
    background-color: map-get($palettes, red);
    color: #fff;
  }
  .mdc-drawer__header {
    /*background-image: url(https://images.weserv.nl/?il&q=100&w=300&url=dl.dropboxusercontent.com/s/51ibl3twvtei59x/bg-drawer.jpg);*/
  }
  .mdc-drawer__header-content {
    display: block;
    /*padding-top: 3rem;*/
    > * {
      width: 320px;
      display: block;
      line-height: 1.25rem;
      color: #fff;
    }
    img {
      width: 4rem;
      height: 4rem;
      margin-bottom: 1rem;
    }
  }
  .mdc-drawer__content {
    //color: map-get($palettes, red);
    // color: #fff;
    // height: 100%;
    > .mdc-card {
      height: 100%;
      background-color: transparent;
      margin-top: 5px;
      // box-shadow: 0px -3px 1px -2px rgba(0, 0, 0, 0.2),
      //   0px -2px 2px 0px rgba(0, 0, 0, 0.14),
      //   0px -1px 0 0px rgba(0, 0, 0, 0.12);
      box-shadow: 0px -2px 4px -1px rgba(0, 0, 0, 0),
        0px -4px 5px 0px rgba(0, 0, 0, 0.1),
        0px -1px 0 0px rgba(0, 0, 0, 0.1);
    }
  }
}
</style>
