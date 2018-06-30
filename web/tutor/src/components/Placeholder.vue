<template lang="pug">
  transition(enter-active-class="animated fadeIn")
    span.placeholder()
      slot
</template>

<script>
  export default {}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "~sass-bem";

  $ph-bg: #fff !default;
  $ph-color: #ced4da !default;
  $ph-border-radius: 2px !default;

  $ph-gutter: 30px !default;
  $ph-spacer: 15px !default;

  $ph-avatar-border-radius: 50% !default;

  .placeholder {
    position: relative;
    @include is('inline') {
      margin-left: .25rem;
    }
    @include is('wait') {
      display: flex;
      min-height: 1rem;

      margin:{
        top:.125rem;
        bottom:.125rem;
      };
      min-width: 5rem;
      background-color: $ph-color;

      &.is-inline {
        display: inline-flex;
        margin-left: .25rem;
      }
      &::before {
        content: " ";
        position: absolute;
        top: 0;
        right: 0;
        width: 100%;
        height: 100%;
        background: linear-gradient(to right, rgba($ph-bg, 0) 46%, rgba($ph-bg, .35) 50%, rgba($ph-bg, 0) 54%) 50% 50%;

        animation: phAnimation .8s linear infinite;
      }
    }
  }

  @keyframes phAnimation {
    0% {
      transform: translate3d(-30%, 0, 0);
    }
    100% {
      transform: translate3d(30%, 0, 0);
    }
  }
</style>
