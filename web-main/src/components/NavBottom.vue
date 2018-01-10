<template>
<nav>
  <ul class="row">
    <li class="col s3">
      <router-link to="/"><i class="material-icons">calendar_today</i></router-link>
    </li>
    <li class="col s3">
      <router-link to="/feedbacks"><i class="material-icons">assignment</i></router-link>
    </li>
    <li class="col s3">
      <router-link to="/progress"><i class="material-icons">assessment</i></router-link>
    </li>
    <li class="col s3">
      <router-link to="/profile"><i class="material-icons">person</i></router-link>
    </li>
  </ul>
</nav>
</template>

<script>
export default {
  name: 'nav-bottom',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
nav {
    position: fixed;
    bottom: 0;
    li {
        text-align: center;
    }
    a {
        padding: 0;
    }
    .material-icons {
        font-size: 2.5rem;
    }
}
</style>
