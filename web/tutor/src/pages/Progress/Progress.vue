<template lang="pug">
  #Progress
    header1
    section.main
      .content
        img(src="static/img/uc.gif")
        br
        br
        | we are still under construction ~~
    tab-bottom
</template>

<script>
export default {
  name: 'progress',
  components: {
    'tab-bottom': () => import('@/components/TabBottom'),
    'header1': () => import('@/components/Header'),
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";

  .content {
    color: map-get($palettes, purple-darken1);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 50%;
    text-align: center;
    font-size: .675rem;
    img{
      width: 100%;
    }
  }
</style>
