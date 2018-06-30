<template lang="pug">
  #TemplateMain
    .grid-y.grid-frame
      .cell.shrink
        header1
      transition(enter-active-class="animated fadeIn" leave-active-class="animated fadeOut")
        .cell.auto(:key="$route.fullPath")
          slot
      .cell.shrink
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
  .cell.auto {
    overflow: auto;
  }
</style>
