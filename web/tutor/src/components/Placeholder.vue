<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-active-class="")
    span.placeholder(:class="classList" :key="value") {{newValue}}
</template>

<script>
  export default {
    props: ['value', 'length', 'valEmpty'],
    computed: {
      newValue() {
        if (!this.value) {
          return !!this.valEmpty ? this.valEmpty : 'lorem ipsum';
        }
        return this.value;
      },
      classList() {
        return [!!this.value ? '' : 'is-wait']
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "~sass-bem";

  /*$ph-bg: #fff !default;
  $ph-color: #ced4da !default;
  $ph-border-radius: 2px !default;

  $ph-gutter: 30px !default;
  $ph-spacer: 15px !default;

  $ph-avatar-border-radius: 50% !default;*/

  .placeholder {
    position: relative;
    @include is('wait') {
      background-color: #C3C3C3;
      color: #C3C3C3;
      &::selection {
        background: #C3C3C3;
        color: #C3C3C3;
      }

      /*&::before {
        content: " ";
        position: absolute;
        top: 0;
        right: 0;
        width: 100%;
        height: 100%;
        background: linear-gradient(to right, rgba($ph-bg, 0) 46%, rgba($ph-bg, .35) 50%, rgba($ph-bg, 0) 54%) 50% 50%;

        animation: phAnimation .8s linear infinite;
      }*/
    }
  }

  /*@keyframes phAnimation {
    0% {
      transform: translate3d(-30%, 0, 0);
    }
    100% {
      transform: translate3d(30%, 0, 0);
    }
  }*/
</style>
