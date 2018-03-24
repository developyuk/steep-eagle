<template lang="pug">
  #header
    header
      img.logo(src="static/img/logo.svg")
      span.search(@click="onClickSearch($event)").animated
        i.material-icons search
        input(type="text" name="q" v-model.trim="q" )
        .remove(@click="onClickClearSearch($event)") x
    aside.mdc-drawer.mdc-drawer--temporary.mdc-typography
      nav.mdc-drawer__drawer
        header.mdc-drawer__header
          .mdc-drawer__header-content
            .photo: img(:src="currentAuth.photo")
            .name {{currentAuth.name}}
            .email {{currentAuth.email}}
        nav#icon-with-text-demo.mdc-drawer__content.mdc-list
          .content
            img(src="static/img/uc.gif")
            br
            br
            | we are still under construction ~~
            br
            | or perhaps you just want to sign out?
          .mdc-list-divider(role="separator")
          a.mdc-list-item(href='#' @click.prevent="signOut($event)")
            i.material-icons.mdc-list-item__graphic(aria-hidden='true') power_settings_new
            | Just sign me out !
</template>

<script>
  import _debounce from "lodash/debounce";
  import _throttle from "lodash/throttle";

  export default {
    name: 'header',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        currentAuth: {},
        q: ''
      }
    }, watch: {
      q: _debounce(function (val) {
        if (!val) {
          const $cont = document.querySelector('.search');
          $cont.classList.toggle('is-opened');
          $cont.classList.remove('fadeInRight');
          $cont.classList.add('fadeOutRight');
          setTimeout(() => $cont.classList.remove('fadeOutRight'), 100);
        }
        this.$bus.$emit('onKeyupSearch', val);
      }, 500)
    },
    methods: {
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
        }
      }
    },
    destroyed() {
      this.$bus.$off('currentAuth');
    },
    mounted() {
      let drawer = new mdc.drawer.MDCTemporaryDrawer(document.querySelector('.mdc-drawer--temporary'));
      document.querySelector('img.logo').addEventListener('click', () => {
        drawer.open = true;
      });
      this.$bus.$on('currentAuth', (auth) => {
        auth.photo = auth.photo ? auth.photo : "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=";
        if (auth.photo.indexOf('data:image/gif') < 0) {
          auth.photo = auth.photo.replace('https://', '').replace('http://', '');
          auth.photo = `//images.weserv.nl/?il&q=100&w=64&h=64&t=square&shape=circle&url=${auth.photo}`;
        }
        this.currentAuth = auth;
      });

    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "../assets/shared";

  header {
    position: relative;
    background-color: $mdc-theme-primary;
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
    .mdc-drawer__header-content {
      display: block;
      /*padding-top: 3rem;*/
      > * {
        width: 320px;
        display: block;
        line-height: 1.25rem;
        color: #fff;
      }
      img {
        width: 4rem;
        height: 4rem;
        margin-bottom: 1rem;
      }
    }
    .mdc-drawer__content {
      color: map-get($palettes, purple-darken1);
    }
    .content {
      position: absolute;
      top: 60%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 80%;
      text-align: center;
      font-size: .675rem;
    }
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
        color: $mdc-theme-primary;
      }
      bottom: 0;
    }
    .mdc-list-divider {
      bottom: 3rem;
    }
  }
</style>
