<template>
<section>
  <header class="card-panel">

    <router-link :to="'/profile'"><i class="material-icons">arrow_back</i> </router-link> Settings

  </header>
  <main>
    <div class="container">
      <ul class="collection">
        <li class="collection-item">
          <router-link :to="'/notifications'">Notifications</router-link>
        </li>
        <li class="collection-item">
          <router-link :to="'/how-to-use'">How to use</router-link>
        </li>
        <li class="collection-item"><a :href="'#'" @click.prevent="signout">Sign out</a></li>
      </ul>
    </div>
  </main>
</section>
</template>

<script>
export default {
  name: 'profile-settings',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  },
  methods: {
    signout() {
      localStorage.removeItem('token');
      window.location.reload();

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
header {
    padding: 1rem;
    .material-icons {
        vertical-align: middle;
    }
}
.container {
    margin-top: 1rem;
    font-size: 1rem;
    .collection {
        border: none;
        .collection-item:last-child {
            border-bottom: 1px solid #e0e0e0;
        }
    }
    a {
        color: col;
    }
}
</style>
