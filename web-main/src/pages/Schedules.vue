<template>
<section>
  <header>
    <div class="input-field container">
      <i class="material-icons prefix">search</i>
      <input id="search" type="text" class="validate" placeholder='Try "Pondok Indah"'>
      <!-- <label for="search">Search</label> -->
    </div>
  </header>
  <main class="grey lighten-3">
    <div class="container" v-for="v in classes">
      <div class="card small z-depth-1 class">
        <div class="card-image">
          <router-link :to="`/class/${v.id}`"><img :src="v.image"> </router-link>
          <span class="card-title">{{v.name}}</span>
          <button @click="activate(v.id,$event)" class="btn-floating halfway-fab waves-effect waves-light red"><i class="material-icons">check</i></button>
        </div>
        <div class="card-content">
          <div class="row">
            <div class="col s6 class__day"><i class="material-icons">date_range</i> {{v.day}}</div>
            <div class="col s6 class__time"><i class="material-icons">schedule</i>{{v.time}}</div>
          </div>
        </div>
      </div>
    </div>
  </main>
  <footer>
    <nav-bottom></nav-bottom>
  </footer>
</section>
</template>

<script>
import NavBottom from '@/components/NavBottom'
import axios from 'axios'

export default {
  name: 'schedules',
  components: {
    'nav-bottom': NavBottom
  },
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA',
      classes: []
    }
  },
  methods: {
    getSchedules(page = 1) {
      const url = `${process.env.API}/classes?day=today`;

      axios.get(url)
        .then(response => this.classes = response.data._embedded)
        .catch(error => console.log(error))

    },
    activate(id, e) {
      const url = `${process.env.API}/classes/${id}/sessions`;
      axios.post(url)
        .then(response => console.log(response))
        .catch(error => console.log(error))
      console.log(id, e);
    }
  },
  mounted() {
    this.getSchedules()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "~sass-bem";
main,
section {
    min-height: 100%;
    height: 100%;
}
header {
    input[type=text]:not(.browser-default) {
        border-bottom: none;
        margin-bottom: 0;
    }
}
main {
    padding: 0.75rem 0;
    .card {
        height: auto;
        .card-image {
            max-height: 10rem;
            position: static;

        }
        .btn-floating.halfway-fab {
            bottom: 2.5rem;
        }
        .card-content {
            padding: 1rem;
            font-size: 1rem;
            line-height: 1.25rem;

            .material-icons {
                font-size: 1.25rem;
                margin-right: 0.25rem;
            }
            .row {
                margin-bottom: 0;
                padding: 0 1rem;
            }
            .col {
                padding: 0;
                &:last-child {
                    text-align: right;
                }
            }
        }
    }

}
.class {
    @include e(day) {
        text-transform: capitalize;
    }
}
</style>
