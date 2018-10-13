<template lang="pug">
  #my-mdc-dialog.mdc-dialog(role="alertdialog" aria-modal="true" aria-labelledby="my-dialog-title" aria-describedby="my-dialog-content")
    .mdc-dialog__container
      .mdc-dialog__surface
        h2#my-dialog-title.mdc-dialog__title
          slot(name="title")
        #my-dialog-content.mdc-dialog__content
          slot
        footer.mdc-dialog__actions
          slot(name="actions")
    .mdc-dialog__scrim

</template>

<script>
import { MDCDialog, MDCDialogFoundation } from "@material/dialog";
import { MDCRipple } from "@material/ripple";

export default {
  mounted() {
    const $dialog = new MDCDialog(this.$el);

    const buttons = [...this.$el.querySelectorAll(".mdc-button")];
    if (buttons.length) {
      buttons.forEach(v => {
        new MDCRipple(v);
      });
    }

    this.$emit("mounted", $dialog);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
