<template lang="pug">
  header.mdc-top-app-bar.mdc-top-app-bar--fixed
    .mdc-top-app-bar__row
      section.mdc-top-app-bar__section.mdc-top-app-bar__section--align-start
        a.material-icons.mdc-top-app-bar__navigation-icon(href="#" @click.prevent="$router.push('/')") arrow_back
        span.mdc-top-app-bar__title 
          .input
            input(type="text" v-model='q' placeholder='Search')
      section.mdc-top-app-bar__section.mdc-top-app-bar__section--align-end(role="toolbar")
        a.search.material-icons.mdc-top-app-bar__action-item(href="#" @click.prevent="$router.push('/')" aria-label="Close" alt="Close") close

</template>

<script>
import { mapState } from "vuex";
import { MDCTopAppBar } from "@material/top-app-bar/index";

export default {
  components: {},
  computed: {
    ...mapState(["currentAuth", "currentStats"])
  },
  data() {
    return {
      q: ""
    };
  },
  methods: {},
  destroyed() {},
  mounted() {
    const topAppBarElement = this.$el;
    const topAppBar = new MDCTopAppBar(topAppBarElement);
    topAppBar.setScrollTarget(document.getElementById("main-content"));

    this.$el.querySelector("input").focus();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";
@import "~sass-bem";
@import "@material/top-app-bar/mdc-top-app-bar";

.input {
  input {
    color: #fff;
    background-color: transparent;
    border: none;
    width: 100%;
    height: 100%;
    padding: 0.5rem;
  }
}
</style>
