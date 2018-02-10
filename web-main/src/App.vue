<template lang="pug">
  #app
    .mdc-layout-grid
      router-view
</template>

<script>
  export default {
    name: 'app',
    mounted() {
      window.mdc.autoInit();
    }
  }
</script>

<style lang="scss">
  :root {
    --mdc-theme-primary: #ED235C;
  }
  #app {
    background-color: #f9f9f9;
  }

  .mdc-layout-grid {
    background-color: #fff;
    max-width: 30rem;
    padding: 0;
  }
  h1.title{
    background-color: var(--mdc-theme-primary,#6200ee);
    color:#fff;
    text-align: center;
    padding:1rem 0;
  }
</style>
