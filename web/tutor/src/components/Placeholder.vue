<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-active-class="")
    span.placeholder(:class="classList" :key="value") {{newValue}}
</template>

<script>
export default {
  props: ["value", "valEmpty"],
  computed: {
    newValue() {
      if (!this.value) {
        return !!this.valEmpty ? this.valEmpty : "lorem ipsum";
      }
      return this.value;
    },
    classList() {
      return [!!this.value ? "" : "is-wait"];
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "~sass-bem";
@import "src/assets/shared";

$ph-color: map-get($palettes, gainsboro) !default;

.placeholder {
  position: relative;
  @include is("wait") {
    background-color: $ph-color;
    color: $ph-color;
    &::selection {
      background: $ph-color;
      color: $ph-color;
    }
  }
}
</style>
