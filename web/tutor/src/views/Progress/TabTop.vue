<template lang="pug">
  tab(:items="tabItems" :indicator="true" elevation="z2").tab-progress
</template>

<script>
export default {
  components: {
    tab: () => import("@/components/Tab")
  },
  data() {
    return {
      tabItems: [
        {
          to: "/progress/classes",
          text: "Classes"
        },
        {
          to: "/progress/students",
          text: "Students"
        }
      ]
    };
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">
.tab-progress {
  .mdc-tab{
    height: 2.5rem;
  }
  .mdc-tab__text-label {
    font-size: 0.75rem;
  }
}
</style>
