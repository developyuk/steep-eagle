<template lang="pug">
  #TemplateMain
    header1
    transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
      main(:key="$route.fullPath")
        slot
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
<style scoped lang="scss">
  main {
    $remain: 8.5rem;
    min-height: calc(500px - #{$remain});
    height: calc(100vh - #{$remain});
    overflow: auto;
  }
</style>
