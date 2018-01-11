<template>
<section>
  <header class="card-panel">
    <div class="row">
      <div class="col s2 offset-s10">
        <router-link :to="'/profile/settings'"><i class="material-icons">settings</i> </router-link>
      </div>
    </div>
  </header>
  <main></main>
  <footer>
    <nav-bottom></nav-bottom>
  </footer>
</section>
</template>

<script>
import NavBottom from '@/components/NavBottom'
export default {
  components: {
    'nav-bottom': NavBottom
  },
  name: 'profile',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
header {
    padding: 1rem;
    // clear: both;
    a {
        float: right;
    }
    .row {
        margin-bottom: 0;
        .col {
            padding: 0;
        }
    }
}
</style>
