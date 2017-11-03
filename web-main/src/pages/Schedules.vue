<template>
<section>
  <header></header>
  <main>
    <div class="card small" v-for="class in classes">
    </div>
  </main>
  <footer>
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
  </footer>
</section>
</template>

<script>
import axios from 'axios'
export default {
  name: 'schedules',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA',
      classes: []
    }
  },
  methods: {
    getSchedules() {
      const url = `${process.env.API}/classes?day=today`;

      axios.get(url).then(
        response => {
          this.classes = response.data;
        },
        error => console.log(error))
    }
  },
  mounted() {
    this.getSchedules()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
footer {
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
}
</style>
