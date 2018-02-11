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
  html,body{
    font-family: 'Roboto', sans-serif;
  }
  :root {
    --mdc-theme-primary: #ED235C;
    /*--mdc-theme-text-primary-on-background:#ED235C;*/
    /*color: var(,rgba(0,0,0,.87));*/
  }

  #app {
    background-color: #f9f9f9;
  }

  .mdc-layout-grid {
    background-color: #fff;
    max-width: 30rem;
    padding: 0;
    /*height: 100vh;*/
  }

  header {
    background-color: var(--mdc-theme-primary, #6200ee);
    img.logo {
      width: 1.5rem;
      padding: 1rem;

      filter: drop-shadow(1px 3px 1px rgba(0, 0, 0, .5));
    }
  }

  h1.title {
    background-color: var(--mdc-theme-primary, #6200ee);
    color: #fff;
    text-align: center;
    padding: 1rem 0;
    margin: 0;
  }
</style>
