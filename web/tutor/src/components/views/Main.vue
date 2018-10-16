<template lang="pug">
  #TemplateMain
    my-drawer(v-model="drawer")
      template(slot="header")
        h3.mdc-drawer__title
          my-img(:src="currentAuth.photo")
          | {{currentAuth.name || currentAuth.username}}
        h6.mdc-drawer__subtitle {{currentAuth.email}}

      ul.mdc-list.mdc-list--non-interactive
        li.mdc-list-item(style="height:auto")
          //- span.mdc-list-item__text Inbox
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
        li.mdc-list-item(style="height:auto")
          //- span.mdc-list-item__text Outgoing
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

        li.mdc-list-divider(role="separator")
        a.mdc-list-item.mdc-list-item--activated(href="#" aria-selected="true" @click.prevent="signOut")
          i.material-icons.mdc-list-item__graphic(aria-hidden="true") power_settings_new
          span.mdc-list-item__text Just sign me out !

    .mdc-drawer-scrim
    
    my-header(:drawer="drawer")
    #mainContent(:key="$route.fullPath")
      .mdc-top-app-bar--fixed-adjust
        slot
    tab-bottom
</template>

<script>
import MyDrawer from "@/components/Drawer";
import { mapState, mapMutations, mapActions } from "vuex";
import { MDCList } from "@material/list";
import MyImg from "@/components/Img";
import TabBottom from "@/components/TabBottom";
import MyHeader from "@/components/headers/Header";

export default {
  components: { MyDrawer, MyImg, TabBottom, MyHeader },
  computed: {
    ...mapState(["currentAuth", "currentStats"])
  },
  data() {
    return {
      drawer: null
    };
  },
  methods: {
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
@import "src/assets/shared";
@import "@material/top-app-bar/mdc-top-app-bar";

.mdc-drawer {
  .mdc-layout-grid__inner {
    grid-template-columns: repeat(3, minmax(0, 1fr));
  }
  .mdc-layout-grid__cell {
    text-align: center;
  }

  img {
    display: block;
    width: 4rem;
    border-radius: 50%;
  }
  //   .mdc-list-item,
  //   .mdc-list-divider {
  //     // position: absolute;
  //     // width: 100%;
  //   }
  //   .mdc-list-item {
  //     &,
  //     .material-icons {
  //       color: #fff;
  //       height: auto;
  //     }
  //     // bottom: 0;
  //   }
  //   .mdc-list-divider {
  //     border-bottom-color: rgba(0, 0, 0, 0.12);
  //   }
}
.mdc-list {
  // position: absolute;
  // bottom: 0;
  // width: 100%;
}
.mdc-list-item {
  &,
  a,
  .material-icons {
    color: #fff;
  }
  .mdc-tab__icon {
    opacity: 1;
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
      $size: 2rem;
      font-size: $size;
      width: $size;
    }
  }
  @include m(texts) {
    height: calc(100vh - #{151px+62px+40px+16px+32px});
    font-size: 0.75rem;
    color: map_get($palettes, red);
    width: 100%;
    .mdc-card {
      background-color: #fff;
      margin: 0.75rem 0;
      padding: 0.5rem 1rem;
      border-radius: 0.325rem;
    }
  }
}
// .sign-out {
//   *:focus {
//     outline: 0;
//   }
//   height: auto;
//   // height: 3rem;
//   width: 100%;
//   .mdc-list-item {
//     height: 100%;
//     padding: 0.25rem 0.75rem;
//     margin: 0;
//   }
//   a {
//     color: #fff;
//     text-decoration: none;
//     display: block;
//     width: 100%;
//     height: 2rem;
//     // padding-top: 1.5rem;
//   }
//   .mdc-list-item__graphic {
//     vertical-align: middle;
//   }
// }
</style>
