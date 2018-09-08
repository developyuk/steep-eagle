<template lang="pug">
  tab(:items="tabItems" elevation="z10").tab-bottom
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
          to: "/",
          icon: "home",
          text: "Home"
        },
        {
          to: "/students",
          icon: "face",
          text: "Student"
        },
        {
          to: "/progress/classes",
          icon: "assessment",
          text: "Progress"
        }
      ]
    };
  },
  methods: {},
  mounted() {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss">
@import "src/assets/shared";
.tab-bottom {
  position: fixed;
  bottom: 0;
  background-color: #fff;
  max-width: 30rem;
  .mdc-tab__icon {
    padding: 0.125rem 0;
    $size: 2rem;
    font-size: $size;
    width: $size;
  }
  .mdc-tab__text-label {
    font-size: 0.5rem;
    text-align: center;
    padding-bottom: 0.25rem;
  }
  .mdc-tab--stacked {
    height: 3rem;
  }
}
</style>
