<template lang="pug">
  aside#my-mdc-dialog.mdc-dialog(role="alertdialog" aria-labelledby="my-mdc-dialog-label" aria-describedby="my-mdc-dialog-description")
    .mdc-dialog__surface
      header.mdc-dialog__header
        h2#my-mdc-dialog-label.mdc-dialog__header__title
          slot(name="header")
      section#my-mdc-dialog-description.mdc-dialog__body
        slot
      footer.mdc-dialog__footer
        slot(name="footer")
    .mdc-dialog__backdrop
</template>

<script>
  import {MDCDialog} from '@material/dialog';
  import {MDCRipple} from '@material/ripple';

  export default {
    mounted() {
      const $dialog = new MDCDialog(this.$el);
//      new MDCRipple(this.$el.querySelector('.mdc-button'));

      this.$emit('mounted', $dialog);
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
