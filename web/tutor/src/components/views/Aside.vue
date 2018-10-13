<template lang="pug">
  #TemplateMain
    my-header
    #main-content(:key="$route.fullPath")
      slot
</template>

<script>
import TabBottom from "@/components/TabBottom";
import MyHeader from "@/components/headers/Aside";
export default {
  components: { TabBottom, MyHeader },
  data() {
    return {};
  },
  mounted() {
    // console.log(this.$route.fullPath);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.cell.auto {
  overflow: auto;
}
</style>
