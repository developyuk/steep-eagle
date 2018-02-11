<template lang="pug">
    #students.mdc-typography
      header
        img.logo(src="static/img/logo.svg")
      tab-bottom
</template>

<script>
  import TabBottom from '@/components/TabBottom';
  export default {
    name: 'students',
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
  #students{
    position: relative;
    height: 100vh;
  }
</style>
