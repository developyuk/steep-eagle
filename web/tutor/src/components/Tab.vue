<template lang="pug">
  nav.mdc-tab-bar(role="tablist" :class="[`mdc-elevation--${elevation || 'z1'}`]")
    .mdc-tab-scroller
      .mdc-tab-scroller__scroll-area
        .mdc-tab-scroller__scroll-content
          button.mdc-tab(role="tab" aria-selected="true" v-for="v in items" :key="v.to" @click="onClick($event,v.to)" :class="{'mdc-tab--active': isActivePath(v.to),'mdc-tab--stacked':!!v.icon}")
            span.mdc-tab__content
              span.mdc-tab__icon.material-icons(v-if="v.icon") {{v.icon}}
              span.mdc-tab__text-label {{v.text}}
            span.mdc-tab-indicator(:class="{'mdc-tab-indicator--active':isActivePath(v.to),'hide':!indicator}")
              span.mdc-tab-indicator__content.mdc-tab-indicator__content--underline
            span.mdc-tab__ripple
</template>

<script>
//  import {MDCTab, MDCTabFoundation} from '@material/tabs';
import { MDCTabBar } from "@material/tab-bar";

export default {
  props: ["items", "indicator","elevation"],
  data() {
    return {};
  },
  methods: {
    isActivePath(path) {
      // if (this.$route.path === "/") {
      //   return this.$route.path === path;
      // } else {
      //   return this.$route.path.indexOf(path) === 0 && path.length > 1;
      // }
      return this.$route.path === path
    },
    onClick(e, to) {
      this.$router.push(to);
    }
  },
  mounted() {
    const tabBar = new MDCTabBar(this.$el);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "../assets/shared";
</style>
