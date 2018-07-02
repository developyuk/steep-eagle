<template lang="pug">
  div(:class='sidebarClasses' :data-background-color='backgroundColor' :data-active-color='activeColor')
    h1.logo M
    .sidebar-wrapper(ref='sidebarScrollArea')
      slot
      ul(:class='navClasses')
        slot(name='links')
          sidebar-item(v-for='(link, index) in sidebarLinks' :key='link.name + index' :link='link')
            sidebar-item(v-for='(subLink, index) in link.children' :key='subLink.name + index' :link='subLink')
</template>
<script>
  export default {
    props: {
      title: {
        type: String,
        default: 'Paper Dashboard Pro'
      },
      type: {
        type: String,
        default: 'sidebar',
        validator: (value) => {
          let acceptedValues = ['sidebar', 'navbar']
          return acceptedValues.indexOf(value) !== -1
        }
      },
      backgroundColor: {
        type: String,
        default: 'white',
        validator: (value) => {
          let acceptedValues = ['white', 'brown', 'black']
          return acceptedValues.indexOf(value) !== -1
        }
      },
      activeColor: {
        type: String,
        default: 'success',
        validator: (value) => {
          let acceptedValues = ['primary', 'info', 'success', 'warning', 'danger']
          return acceptedValues.indexOf(value) !== -1
        }
      },
      logo: {
        type: String,
        default: 'static/img/vue-logo.png'
      },
      sidebarLinks: {
        type: Array,
        default: () => []
      }
    },
    computed: {
      sidebarClasses () {
        if (this.type === 'sidebar') {
          return 'sidebar'
        } else {
          return 'collapse navbar-collapse off-canvas-sidebar'
        }
      },
      navClasses () {
        if (this.type === 'sidebar') {
          return 'nav'
        } else {
          return 'nav navbar-nav'
        }
      }
    },
    methods: {
      async initScrollBarAsync () {
        await import('perfect-scrollbar/dist/css/perfect-scrollbar.css')
        const PerfectScroll = await import('perfect-scrollbar')
        PerfectScroll.initialize(this.$refs.sidebarScrollArea)
      }
    },
    mounted () {
      this.initScrollBarAsync()
    },
    beforeDestroy () {
      if (this.$sidebar.showSidebar) {
        this.$sidebar.showSidebar = false
      }
    }
  }

</script>
<style lang="scss">
  @import "../../assets/scss/_shared.scss";

  @media (min-width: 992px) {
    .navbar-search-form-mobile,
    .nav-mobile-menu{
      display: none;
    }
  }
</style>
