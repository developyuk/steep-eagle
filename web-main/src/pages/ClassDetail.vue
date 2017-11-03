<template>
<section>
  <header></header>
  <main>{{$route.params.id}}</main>
  <footer></footer>
</section>
</template>

<script>
export default {
  name: 'classDetail',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
</style>
