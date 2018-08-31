<template lang="pug">
  #header
    header
      .grid-x.align-middle.align-justify
        .cell.shrink
          img.logo(src="img/logo.svg")
        //- .cell.shrink
        //-   span.search(@click="onClickSearch") search?

</template>

<script>
import _debounce from "lodash/debounce";
import _throttle from "lodash/throttle";
import { mapState, mapMutations, mapActions } from "vuex";

export default {
  props: ['drawer'],
  components: {
  },
  data() {
    return {
      q: "",
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
    this.$el
      .querySelector("img.logo")
      .addEventListener("click", () => (this.drawer.open = true));
    this.updateStats();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "../../assets/shared";
@import "~sass-bem";

header {
  position: relative;
  background-color: $mdc-theme-primary;
  overflow: hidden;
}

img.logo {
  width: 1.5rem;
  padding: 1rem;

  filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, 0.5));
}

.search {
  font-weight: 500;
  color: #fff;
  margin-right: 1rem;
}
</style>
