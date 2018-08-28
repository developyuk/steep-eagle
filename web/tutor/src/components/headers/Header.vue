<template lang="pug">
  #header
    header
      .grid-x.align-middle.align-justify
        .cell.shrink
          img.logo(src="img/logo.svg")
        //- .cell.shrink
        //-   span.search(@click="onClickSearch") search?
    my-drawer(@mounted="onMountedDrawer")
      template(slot="header")
        .photo
          my-img(:src="currentAuth.photo" myIs="tutor")
        .name.mdc-typography--headline6 {{currentAuth.name || currentAuth.username}}
        .email {{currentAuth.email}}

      ul.mdc-list.stats-list
        li.mdc-list-item
          .stats.stats--icons.mdc-layout-grid
            .mdc-layout-grid__inner
              .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
                i.material-icons.mdc-tab__icon(aria-hidden="true") class
                .text {{currentStats.classes}} classes
              .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
                i.material-icons.mdc-tab__icon(aria-hidden="true") watch_later
                .text {{currentStats.hours}} hours
              .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
                i.material-icons.mdc-tab__icon(aria-hidden="true") stars
                .text {{currentStats.feedbacks}} feedbacks
        li.mdc-list-item
          .stats.stats--texts
            .mdc-card.ratings
              .title ratings given
              .text {{currentStats.ratings}}%
            .mdc-card.reviews
              .title reviews given
              .text {{currentStats.reviews}}%
            .mdc-card.attendances
              .title teaching attendances
              .text {{currentStats.attendances}}%

      ul.mdc-list.sign-out
        li.mdc-list-divider(role="separator")
        li.mdc-list-item
          a(href='#' @click.prevent="signOut($event)")
              i.material-icons.mdc-list-item__graphic(aria-hidden='true') power_settings_new
              | Just sign me out !

</template>

<script>
import _debounce from "lodash/debounce";
import _throttle from "lodash/throttle";
import { mapState, mapMutations, mapActions } from "vuex";
import MyDrawer from "@/components/Drawer";

export default {
  components: {
    MyDrawer,
    "my-img": () => import("@/components/Img")
  },
  computed: {
    ...mapState(["currentAuth", "currentStats"])
  },
  data() {
    return {
      q: "",
      $drawer: null
    };
  },
  methods: {
    ...mapMutations(["nextSearch"]),
    ...mapActions(["updateStats"]),
    onMountedDrawer(e) {
      this.$drawer = e;
    },
    signOut(e) {
      localStorage.removeItem("token");
      window.location.reload();
    },
    onClickSearch(e) {
      this.$router.push("/search");
    }
  },
  destroyed() {},
  mounted() {
    this.$el
      .querySelector("img.logo")
      .addEventListener("click", () => (this.$drawer.open = true));
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

.mdc-drawer {
  .mdc-layout-grid__inner {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
  .mdc-layout-grid__cell {
    text-align: center;
  }
}

.mdc-drawer {
  img {
    width: 8rem;
    border-radius: 50%;
  }
  .mdc-list-item,
  .mdc-list-divider {
    // position: absolute;
    // width: 100%;
  }
  .mdc-list-item {
    &,
    .material-icons {
      color: #fff;
      height: auto;
    }
    // bottom: 0;
  }
  .mdc-list-divider {
    // bottom: 3rem;

    border-bottom-color: rgba(0, 0, 0, 0.12);
  }
}

.stats {
  color: #fff;
  .title {
    text-transform: capitalize;
  }
  @include m(icons) {
    font-size: 0.675rem;
    padding: 0;
    width: 100%;
    .material-icons {
      font-size: 2.5rem;
      width: 2.5rem;
      height: 2.5rem;
    }
  }
  @include m(texts) {
    font-size: 0.75rem;
    color: map_get($palettes, red);
    width: 100%;
    .mdc-card {
      background-color: #fff;
      margin: 0.75rem 0;
      padding: 0.5rem 1rem;
      border-radius: 0.325rem;
      // box-shadow: 0 0.25rem rgba(0, 0, 0, 0.25);
    }
  }
}
.search {
  font-weight: 500;
  color: #fff;
  margin-right: 1rem;
}
.stats-list{
  height: calc(100% - 3rem);
  // display: block
  min-height:20rem;
  // height:100%;
}
.sign-out {
  height: 3rem;
  width: 100%;
  .mdc-list-item {
    height: 100%;
    padding: 0 1rem;
  }
  a {
    color: #fff;
    text-decoration: none;
    display: block;
    width: 100%;
    height: 100%;
    padding-top: 1.5rem;
  }
  .mdc-list-item__graphic {
    vertical-align: middle;
  }
}
</style>
