<template lang="pug">
  #header
    header
      img.logo(src="img/logo.svg")
      //span.search search?
    my-drawer(@mounted="onMountedDrawer")
      template(slot="header")
        .photo
          my-img(:src="currentAuth.photo" myIs="profile")
        .name.mdc-typography--headline6 {{currentAuth.name}}
        .email {{currentAuth.email}}

      .stats.stats--icons.mdc-layout-grid
        .mdc-layout-grid__inner
          .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
            i.material-icons.mdc-tab__icon(aria-hidden="true") class
            .text {{currentStats.classes}} classes
          .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
            i.material-icons.mdc-tab__icon(aria-hidden="true") access_time
            .text {{currentStats.hours}} hours
          .mdc-layout-grid__cell.mdc-layout-grid__cell--span-1
            i.material-icons.mdc-tab__icon(aria-hidden="true") stars
            .text {{currentStats.feedbacks}} feedbacks
      .stats.stats--texts
        .ratings
          .title ratings given
          .text {{currentStats.ratings}}%
        .reviews
          .title reviews given
          .text {{currentStats.reviews}}%
        .attendances
          .title teaching attendances
          .text {{currentStats.attendances}}%
      .mdc-list-divider(role="separator")
      a.mdc-list-item(href='#' @click.prevent="signOut($event)")
        i.material-icons.mdc-list-item__graphic(aria-hidden='true') power_settings_new
        | Just sign me out !

</template>

<script>
  import _debounce from "lodash/debounce";
  import _throttle from "lodash/throttle";
  import {mapState, mapMutations, mapActions} from 'vuex'
  import MyDrawer from '@/components/Drawer';

  export default {
    components: {
      MyDrawer,
      'my-img': () => import('@/components/Img'),
    },
    computed: {
      ...mapState(['currentAuth', 'currentStats']),
    },
    data() {
      return {
        q: '',
        $drawer: null,
      }
    },
    methods: {
      ...mapMutations(['nextSearch']),
      ...mapActions(['updateStats']),
      onMountedDrawer(e) {
        this.$drawer = e;
      },
      signOut(e) {
        localStorage.removeItem('token');
        window.location.reload();
      },
    },
    destroyed() {
    },
    mounted() {
      this.$el.querySelector('img.logo').addEventListener('click', () => this.$drawer.open = true);
      this.updateStats();
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "../assets/shared";
  @import "~sass-bem";

  header {
    position: relative;
    background-color: $mdc-theme-primary;
    overflow: hidden;
  }

  img.logo {
    width: 1.5rem;
    padding: 1rem;

    filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, .5));
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
    .mdc-list-item, .mdc-list-divider {
      position: absolute;
      width: 100%;
    }
    .mdc-list-item {
      &, .material-icons {
        color: #fff;
      }
      bottom: 0;
    }
    .mdc-list-divider {
      bottom: 3rem;
    }
  }

  .stats {
    color: #fff;
    .title {
      text-transform: capitalize;
    }
    @include m(icons) {
      font-size: .675rem;
      .material-icons {
        font-size: 2.5rem;
        width: 2.5rem;
        height: 2.5rem;
      }
    }
    @include m(texts) {
      font-size: .75rem;
      color: map_get($palettes, red);
      > div {
        background-color: #fff;
        margin: .325rem 1rem;
        padding: .325rem 1rem;
        border-radius: .5rem;
      }
    }
  }
</style>
