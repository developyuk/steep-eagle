<template lang="pug">
  template-aside#search
    tab-top
</template>

<script>
import TemplateAside from "@/components/views/Aside";
export default {
  components: {
    TemplateAside,
    "tab-top": () => import("./TabTop")
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
