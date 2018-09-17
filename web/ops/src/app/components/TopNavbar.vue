<template lang="pug">
  nav.navbar.navbar-default
    .container-fluid
      .navbar-minimize
        button.btn.btn-fill.btn-icon(@click='minimizeSidebar')
          i(:class="$sidebar.isMinimized ? 'ti-menu-alt' : 'ti-more-alt'")
      .navbar-header
        button.navbar-toggle(type='button' :class='{toggled: $sidebar.showSidebar}' @click='toggleSidebar')
          span.sr-only Toggle navigation
          span.icon-bar.bar1
          span.icon-bar.bar2
          span.icon-bar.bar3
        a.navbar-brand {{this.$route.name}}
      .collapse.navbar-collapse
        form.navbar-form.navbar-left.navbar-search-form(role='search')
          .input-group
            input.form-control(type='text' value='' placeholder='Search...')
            span.input-group-addon
              i.fa.fa-search
        ul.nav.navbar-nav.navbar-right
          //li.open
          //  router-link.dropdown-toggle.btn-magnify(to='/admin/stats' data-toggle='dropdown')
          //    i.ti-panel
          //    p Stats
          drop-down(tag='li' title='5' icon='ti-bell')
            template(slot="title")
              a.dropdown-toggle.btn-rotate(data-toggle='dropdown' href='javascript:void(0)')
                i(class='icon')
                p.notification
                  img(:src="currentAuth.photo")
                  b.caret
            li
              router-link(:to="`/admin/${currentAuth.role}s/${currentAuth.id}`") Profile
            li
              a(href='#') Settings
            li
              a(@click.prevent="onClickSignOut") Sign Out
          //li
          //  router-link.btn-rotate(to='/admin/overview')
          //    i.ti-settings
</template>
<script>
  import {mapState} from 'vuex'
  import MyImg from './Img.vue'

  export default {
    computed: mapState(['currentAuth']),
    components: {
      MyImg
    },
    data() {
      return {
        activeNotifications: false
      }
    },
    methods: {
      onClickSignOut(e) {
        localStorage.removeItem('token');
        window.location.reload();
      },
      capitalizeFirstLetter(string) {
        return string.charAt(0).toUpperCase() + string.slice(1)
      },
      toggleNotificationDropDown() {
        this.activeNotifications = !this.activeNotifications
      },
      closeDropDown() {
        this.activeNotifications = false
      },
      toggleSidebar() {
        this.$sidebar.displaySidebar(!this.$sidebar.showSidebar)
      },
      hideSidebar() {
        this.$sidebar.displaySidebar(false)
      },
      minimizeSidebar() {
        this.$sidebar.toggleMinimize()
      }
    }
  }

</script>
<style scoped lang="scss">
  .notification {
    img {
      $size: 4rem;
      height: $size;
      width: $size;
      border-radius: 50%;
    }
  }
</style>
