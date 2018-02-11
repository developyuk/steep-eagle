<template lang="pug">
  nav#icon-text-tab-bar.mdc-tab-bar.mdc-tab-bar--icons-with-text.mdc-elevation--z8(data-mdc-auto-init="MDCTabBar")
    router-link(to="/").mdc-tab.mdc-tab--with-icon-and-text(:class="{'mdc-tab--active': isActivePath('/')}")
      i.material-icons.mdc-tab__icon(aria-hidden="true") home
      span.mdc-tab__icon-text Home
    router-link(to="/students").mdc-tab.mdc-tab--with-icon-and-text(:class="{'mdc-tab--active': isActivePath('/students')}")
      i.material-icons.mdc-tab__icon(aria-hidden="true") face
      span.mdc-tab__icon-text Student

    //span.mdc-tab-bar__indicator
</template>

<script>
  import {MDCRipple} from '@material/ripple';

  export default {
    name: 'tabBottom',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA'
      }
    },
    methods: {
      isActivePath(path) {
        return this.$route.path === path
      }
    },
    mounted() {
//      mdc.tabs.MDCTabBar.attachTo(document.querySelector('#icon-text-tab-bar'));
      MDCRipple.attachTo(document.querySelector('.mdc-tab-bar'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .mdc-tab-bar {
    position: absolute;
    bottom: 0;
    max-width: 30rem;
    width: 100%;
    background-color: #fff;
  }
  .mdc-tab--active, .mdc-tab--active .mdc-tab__icon{
    color:#ED235C;
  }
  .mdc-tab__icon{
    font-size: 2rem;
    width: auto;
    height: auto;
    line-height: 1.25rem;
  }
  .mdc-tab__icon-text{
    font-size: .75rem;
  }
</style>
