<template lang="pug">
  template-main#progress
    .content
      img(src="img/uc.gif")
      br
      br
      | we are still under construction ~~
</template>

<script>
  import TemplateMain from '@/templates/TemplateMain';

  export default {
    components: {
      'template-main': TemplateMain,
    },
    data() {
      return {
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
    img {
      width: 100%;
    }
  }
</style>
