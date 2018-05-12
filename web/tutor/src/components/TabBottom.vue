<template lang="pug">
  nav#icon-text-tab-bar.mdc-tab-bar.mdc-tab-bar--icons-with-text.mdc-elevation--z8(data-mdc-auto-init="MDCTabBar")
    router-link(to="/" :class="{'mdc-tab--active': isActivePath('/')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") home
      span.mdc-tab__icon-text Home
    router-link(to="/students" :class="{'mdc-tab--active': isActivePath('/students')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") face
      span.mdc-tab__icon-text Student
    router-link(to="/progress" :class="{'mdc-tab--active': isActivePath('/progress')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") assessment
      span.mdc-tab__icon-text Progress

    //span.mdc-tab-bar__indicator
</template>

<script>
  //  import {MDCRipple} from '@material/ripple';

  export default {
    name: 'tabBottom',
    data() {
      return {
      }
    },
    methods: {
      isActivePath(path) {
        return this.$route.path === path
      }
    },
    mounted() {
      window.mdc.autoInit();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../assets/shared";

  .mdc-tab-bar {
    position: absolute;
    bottom: 0;
    max-width: 30rem;
    width: 100%;
    background-color: #fff;
    z-index: 4;
  }

  .mdc-tab--active, .mdc-tab--active .mdc-tab__icon {
    color: $mdc-theme-primary;
  }

  .mdc-tab {
    padding: 0 2rem;
  }

  .mdc-tab__icon {
    font-size: 2.5rem;
    width: auto;
    height: auto;
    line-height: 1.25rem;
    vertical-align: middle;
    margin-top: 1rem;
  }

  .mdc-tab__icon-text {
    font-size: .625rem;
  }
</style>
