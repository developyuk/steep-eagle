<template lang="pug">
  nav#icon-text-tab-bar.mdc-tab-bar.mdc-tab-bar--icons-with-text
    router-link(to="/").mdc-tab.mdc-tab--with-icon-and-text.mdc-tab--active
      i.material-icons.mdc-tab__icon(aria-hidden="true") home
      span.mdc-tab__icon-text Home
    router-link(to="/profile").mdc-tab.mdc-tab--with-icon-and-text
      i.material-icons.mdc-tab__icon(aria-hidden="true") face
      span.mdc-tab__icon-text Student

    span.mdc-tab-bar__indicator
</template>

<script>
export default {
  name: 'tabBottom',
  data () {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  .mdc-tab-bar{
    position: absolute;
    bottom: 0;
    width: 480px;
  }
</style>
