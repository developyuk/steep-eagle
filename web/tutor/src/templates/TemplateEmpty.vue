<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
    main: slot
</template>

<script>
export default {
  data () {
    return {
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
