<template lang="pug">
  #TemplateMain
    my-drawer(@mounted="onMountedDrawer")
      template(slot="header")
        .photo
          my-img(:src="currentAuth.photo")
        .name.mdc-drawer__title {{currentAuth.name || currentAuth.username}}
        .email.mdc-drawer__subtitle {{currentAuth.email}}

      .mdc-card.mdc-elevation--z4
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
          li.mdc-list-item(tabindex="-1")
            a(href='#' @click.prevent="signOut($event)")
                i.material-icons.mdc-list-item__graphic(aria-hidden='true') power_settings_new
                | Just sign me out !
    .mdc-drawer-scrim

    .grid-y.grid-frame
      //- .cell.shrink
      my-header(:drawer="drawer")
      //- transition(enter-active-class="animated fadeIn" leave-active-class="animated fadeOut")
      .cell.auto(:key="$route.fullPath")
        slot
      //- .cell.shrink
      tab-bottom
</template>

<script>
import MyDrawer from "@/components/Drawer";
import { mapState, mapMutations, mapActions } from "vuex";
import { MDCList } from "@material/list";

export default {
  components: {
    MyDrawer,
    "my-img": () => import("@/components/Img"),
    "tab-bottom": () => import("@/components/TabBottom"),
    "my-header": () => import("@/components/headers/Header")
  },
  computed: {
    ...mapState(["currentAuth", "currentStats"])
  },
  data() {
    return {
      drawer: null
    };
  },
  methods: {
    onMountedDrawer(e) {
      this.drawer = e;
    },
    signOut(e) {
      localStorage.removeItem("token");
      window.location.reload();
    }
  },
  mounted() {
    // console.log(this.$route.fullPath);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "~sass-bem";
@import "../../assets/shared";
.grid-frame {
  padding-top: 56px;
  padding-bottom: 3rem;
}
.cell.auto {
  overflow: auto;
}

.mdc-drawer {
  .mdc-layout-grid__inner {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
  .mdc-layout-grid__cell {
    text-align: center;
  }
  .photo {
  }
  img {
    width: 4rem;
    height: 4rem;
    border-radius: 50%;
    margin-top: 1rem;
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
      opacity: 1;
      $size: 2rem;
      font-size: $size;
      width: $size;
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
.stats-list {
  height: calc(100% - 3rem);
  // display: block
  min-height: 20rem;
  // height:100%;
}
.sign-out {
  *:focus {
    outline: 0;
  }
  height: auto;
  // height: 3rem;
  width: 100%;
  .mdc-list-item {
    height: 100%;
    padding: 0.25rem 0.75rem;
    margin: 0;
  }
  a {
    color: #fff;
    text-decoration: none;
    display: block;
    width: 100%;
    height: 2rem;
    // padding-top: 1.5rem;
  }
  .mdc-list-item__graphic {
    vertical-align: middle;
  }
}
</style>
