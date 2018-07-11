<template lang="pug">
  #TemplateMain
    .grid-y.grid-frame
      .cell.shrink
        my-header
      transition(enter-active-class="animated fadeIn" leave-active-class="animated fadeOut")
        .cell.auto(:key="$route.fullPath")
          slot
</template>

<script>
  export default {
    components: {
      'tab-bottom': () => import('@/components/TabBottom'),
      'my-header': () => import('@/components/headers/Aside'),
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
