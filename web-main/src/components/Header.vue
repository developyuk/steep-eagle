<template lang="pug">
  #header
    header
      img.logo(src="static/img/logo.svg")
      span.search(@click="comingSoon($event)") search?
    aside.mdc-drawer.mdc-drawer--temporary.mdc-typography
      nav.mdc-drawer__drawer
        header.mdc-drawer__header
          .mdc-drawer__header-content
            .photo: img(:src="'https://images.weserv.nl/?il&q=100&w=64&h=64&t=square&shape=circle&url='+currentAuth.photo")
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
  export default {
    name: 'header',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        currentAuth: {}
      }
    },
    methods: {
      signOut(e) {
        localStorage.removeItem('token');
        window.location.reload();
      },
      comingSoon(e) {
        e.target.classList.remove('animated', 'fadeIn');
        if (e.target.innerText.toLowerCase() !== 'coming soon') {
          e.target.innerText = 'coming soon';
        } else {
          e.target.innerText = 'search?';
        }
        e.target.classList.add('animated', 'fadeIn');
      },
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
        auth.photo = auth.photo.replace('https://', '').replace('http://', '');
        this.currentAuth = auth;
      });

    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  header {
    position: relative;
    background-color: var(--mdc-theme-primary, #6200ee);
  }

  img.logo {
    width: 1.5rem;
    padding: 1rem;

    filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, .5));
  }

  span.search {
    color: #fff;
    text-shadow: 2px 2px 4px rgba(0, 0, 0, .5);
    position: absolute;
    right: 1rem;
    top: 50%;
    transform: translateY(-50%);
    font-weight: 500;
    text-transform: capitalize;
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
      color: #9B51E0;
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
    }
    .mdc-list-item, .mdc-list-divider {
      position: absolute;
      width: 100%;
    }
    .mdc-list-item {
      &, .material-icons {
        color: var(--mdc-theme-primary);
      }
      bottom: 0;
    }
    .mdc-list-divider {
      bottom: 3rem;
    }
  }
</style>
