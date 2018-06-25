<template lang="pug">
  #header
    header
      img.logo(src="img/logo.svg")
      span.search(@click="onClickSearch($event)").animated
        i.material-icons search
        input(type="text" name="q" v-model.trim="q" )
        .remove(@click="onClickClearSearch($event)") x
    my-drawer(@mounted="onMountedDrawer")
      template(slot="header")
        .photo
          my-img(:src="currentAuth.photo" myIs="profile")
        .name {{currentAuth.name}}
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
    watch: {
      q: _debounce(function (val) {
        if (!val) {
          this.nextSearch(null);
          const $cont = this.$el.querySelector('.search');
          $cont.classList.toggle('is-opened');
          $cont.classList.remove('fadeInRight');
          $cont.classList.add('fadeOutRight');
          setTimeout(() => $cont.classList.remove('fadeOutRight'), 100);
        }
        this.nextSearch(val);
      }, 500)
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
      onClickClearSearch(e) {
        this.q = '';
      },
      onClickSearch(e) {
        const $cont = e.target.closest('.search');
        const $input = $cont.querySelector('input');
        const $icon = $cont.querySelector('.material-icons');
        if (!$cont.classList.contains('is-opened')) {
          $cont.classList.toggle('is-opened');
          $cont.classList.remove('fadeOutRight');
          $cont.classList.add('fadeInRight');
          setTimeout(() => $input.focus(), 500);
        } else {
          $cont.classList.toggle('is-opened');
          $cont.classList.remove('fadeInRight');
          $cont.classList.add('fadeOutRight');
          setTimeout(() => $cont.classList.remove('fadeOutRight'), 100);
        }
      }
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

  span.search {
    color: #fff;
    /*text-shadow: 2px 2px 4px rgba(0, 0, 0, .5);*/
    position: absolute;
    top: 50%;
    width: 12.5rem;
    left: calc(100% - 2rem);
    transform: translateY(-50%);
    &.is-opened {
      /*color: var(--mdc-theme-primary, #6200ee);*/
      color: black;
      right: 1rem;
      left: auto;
      top: 20%;
      width: 12rem;

      input {
        visibility: visible;
        width: 9rem;
      }
      .material-icons {
        position: absolute;
        top: 50%;
        transform: translateY(-50%);
        left: .5rem;
        z-index: 1;
      }
    }
    .remove {
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
      right: 1rem;
      font-size: 1.25rem;

    }
    .material-icons {
      vertical-align: middle;
    }
    input {
      visibility: hidden;
      padding: {
        top: .5rem;
        right: .5rem;
        bottom: .5rem;
        left: 2.5rem;
      }
      width: 8rem;
      border: none;
      border-radius: 1rem;
    }
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
    /*.content {
      position: absolute;
      top: 60%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 80%;
      text-align: center;
      font-size: .675rem;
    }*/
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
      .material-icons {
        font-size: 3rem;
        width: 3rem;
        height: 3rem;
      }
    }
    @include m(texts) {
      color: map_get($palettes, red);
      > div {
        background-color: #fff;
        margin: .5rem 1rem;
        padding: .5rem 1rem;
        border-radius: .5rem;
      }
    }
  }
</style>
