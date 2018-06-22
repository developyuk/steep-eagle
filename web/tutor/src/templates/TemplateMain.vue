<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
    #TemplateMain
      header1
      main: slot
      tab-bottom
</template>

<script>
  export default {
    components: {
      'tab-bottom': () => import('@/components/TabBottom'),
      'header1': () => import('@/components/Header'),
    },
    data() {
      return {}
    },
    mounted() {
      console.log(this.$route.fullPath)
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  main {
    min-height: calc(500px - 8rem);
    height: calc(100vh - 8rem);
    overflow: auto;
  }
</style>
