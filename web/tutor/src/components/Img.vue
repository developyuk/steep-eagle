<template lang="pug">
  //- transition(appear-active-class="animated fadeIn" leave-active-class="animated fadeOut")
  img(:src="src" :key="src")
</template>

<script>
export default {
  props: ["src"],
  computed: {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
