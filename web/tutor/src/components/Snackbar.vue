<template lang="pug">
  .mdc-snackbar(aria-live='assertive' aria-atomic='true' aria-hidden='true')
    .mdc-snackbar__text
    .mdc-snackbar__action-wrapper
      button.mdc-snackbar__action-button(type='button')
</template>

<script>

  import {MDCSnackbar} from '@material/snackbar';

  export default {
    mounted() {
      const snackbar = new MDCSnackbar(this.$el);
      this.$emit('mounted', snackbar);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">

</style>
