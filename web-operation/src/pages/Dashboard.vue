<template lang="pug">
  section.row
    section.col.s2.card-panel
      nav-side
    section.col.s10
      header-menu
      section: h1 Dashboard
</template>

<script>
  import NavSide from '@/components/NavSide'
  import Header from '@/components/Header'

  export default {
    name: 'dashboard',
    components: {
      'nav-side': NavSide,
      'header-menu': Header
    },
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA'
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  section.col {
    padding: 0 !important;
  }

  section.row {
    margin: 0;
  }

  .card-panel.col {
    margin: 0;
    height: 100vh;
  }
</style>
