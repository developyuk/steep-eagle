<template lang="pug">
  .wrapper
    side-bar(type='sidebar' :sidebar-links='$sidebar.sidebarLinks')
      //user-menu
      form.navbar-form.navbar-left.navbar-search-form.navbar-search-form-mobile(role='search')
        .input-group
          span.input-group-addon
            i.fa.fa-search
          input.form-control(type='text' value='' placeholder='Search...')
      mobile-menu
    .main-panel
      top-navbar
      dashboard-content(@click.native='toggleSidebar')
      //content-footer

</template>
<script>
  import TopNavbar from './TopNavbar.vue'
  //  import ContentFooter from '../components/Dashboard/Layout/ContentFooter.vue'
  import DashboardContent from '../components/Dashboard/Layout/Content.vue'
  import UserMenu from 'src/components/UIComponents/SidebarPlugin/UserMenu.vue'
  import MobileMenu from 'src/components/UIComponents/SidebarPlugin/MobileMenu.vue'

  export default {
    components: {
      TopNavbar,
//      ContentFooter,
      DashboardContent,
      UserMenu,
      MobileMenu
    },
    methods: {
      toggleSidebar() {
        if (this.$sidebar.showSidebar) {
          this.$sidebar.displaySidebar(false)
        }
      }
    }
  }

</script>

<style lang="scss" scoped>


</style>
