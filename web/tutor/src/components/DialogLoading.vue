<template lang="pug">
    my-dialog.mdc-dialog--loading(v-model="dialog" escapeKeyAction="" scrimClickAction="")
      //- span Loading..
      center
        img(src="https://loading.io/spinners/ring/index.rotate-ring-spinner.svg")
</template>

<script>
import MyDialog from "@/components/Dialog";
export default {
  components: { MyDialog },

  data() {
    return {
      dialog: null
    };
  },
  mounted() {
    this.$emit("input", this.dialog);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">
.mdc-dialog--loading {
  .mdc-dialog__content {
    padding: 0;
  }
}
</style>
