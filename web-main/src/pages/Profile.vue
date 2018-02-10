<template lang="pug">
    #profile.mdc-typography
      h1.title Profile
      tab-bottom
</template>

<script>
  import TabBottom from '@/components/TabBottom';
  export default {
    name: 'profile',
    components: {
      'tab-bottom': TabBottom
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA'
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  #profile{
    position: relative;
    height: calc(100vh - 2rem);
  }
</style>
