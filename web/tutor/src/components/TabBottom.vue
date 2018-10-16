<template lang="pug">
  tab.tab-bottom(:items="tabItems" elevation="z10")
</template>

<script>
import Tab from "@/components/Tab";
export default {
  components: { Tab },
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

  .mdc-tab__text-label {
    font-size: 0.675rem;
  }
  .mdc-tab--stacked {
    height: 3rem;
    .mdc-tab__icon {
      padding-top: 0.325rem;
    }
    .mdc-tab__text-label {
      padding-bottom: 0.325rem;
    }
  }
}
</style>
