<template lang="pug">
  header.mdc-top-app-bar.mdc-top-app-bar--fixed
    .mdc-top-app-bar__row
      section.mdc-top-app-bar__section.mdc-top-app-bar__section--align-start
        a.mdc-top-app-bar__navigation-icon 
          img.logo(src="img/logo.svg")
      //-   span.mdc-top-app-bar__title Title
      //- section.mdc-top-app-bar__section.mdc-top-app-bar__section--align-end(role="toolbar")
      //-   a.search.material-icons.mdc-top-app-bar__action-item(href="#" @click.prevent="onClickSearch" aria-label="Search" alt="Search") search

</template>

<script>
import _debounce from "lodash/debounce";
import _throttle from "lodash/throttle";
import { mapState, mapMutations, mapActions } from "vuex";
import { MDCTopAppBar } from "@material/top-app-bar/index";

export default {
  props: ["drawer"],
  components: {},
  data() {
    return {
      q: ""
    };
  },
  methods: {
    ...mapActions(["updateStats"]),
    onClickSearch(e) {
      this.$router.push("/search");
    }
  },
  destroyed() {},
  mounted() {
    const topAppBarElement = this.$el;
    const topAppBar = new MDCTopAppBar(topAppBarElement);
    topAppBar.setScrollTarget(document.querySelector("#mainContent"));
    topAppBar.listen("MDCTopAppBar:nav", () => {
      this.drawer.open = !this.drawer.open;
    });
    this.updateStats();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "src/assets/shared";
@import "~sass-bem";
@import "@material/top-app-bar/mdc-top-app-bar";

header {
  background-color: $mdc-theme-primary;
  max-width: 30rem;
}

img.logo {
  width: 1.5rem;
  padding: 1rem;

  filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, 0.5));
}

// .search {
//   font-weight: 600;
// }
</style>
