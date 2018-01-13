<template>
<section>
  <header>header</header>
  <main>main</main>
</section>
</template>

<script>
export default {
  name: 'classes',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
$roboto-font-path: "~materialize-css/dist/fonts/roboto/";
@import "~materialize-css/sass/components/color";
$primary-color: color("grey", "base");
$secondary-color: color("grey", "lighten-1");
@import "~materialize-css/sass/materialize.scss";
</style>
