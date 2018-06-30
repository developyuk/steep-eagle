<template lang="pug">
  nav#icon-text-tab-bar.mdc-tab-bar.mdc-tab-bar--icons-with-text.mdc-elevation--z8()
    router-link(to="/" :class="{'mdc-tab--active': isActivePath('/')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") home
      span.mdc-tab__icon-text Home
    router-link(to="/students" :class="{'mdc-tab--active': isActivePath('/students')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") face
      span.mdc-tab__icon-text Student
    router-link(to="/progress" :class="{'mdc-tab--active': isActivePath('/progress')}").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") assessment
      span.mdc-tab__icon-text Progress
    span.mdc-tab-bar__indicator.hide
</template>

<script>
//  import {MDCTab, MDCTabFoundation} from '@material/tabs';
  import {MDCTabBar, MDCTabBarFoundation} from '@material/tabs';

  export default {
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
      new MDCTabBar(document.querySelector('#icon-text-tab-bar'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../assets/shared";

  .mdc-tab-bar {
    width: 100%;
    background-color: #fff;
  }

  .mdc-tab--active, .mdc-tab--active .mdc-tab__icon {
    color: $mdc-theme-primary;
  }

  .mdc-tab {
    line-height: 1rem;
    padding: .25rem;
  }

  .mdc-tab__icon {
    font-size: 2rem;
    width: auto;
    height: auto;
    margin: 0;
  }

  .mdc-tab__icon-text {
    font-size: .675rem;
    margin: 0;
    padding: 0;
  }
  .mdc-tab-bar--icons-with-text{
    height: auto;
  }
</style>
