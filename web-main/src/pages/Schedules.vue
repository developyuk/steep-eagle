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
            <router-link :to="`/class/${v.id}`"><img :src="v.image"></router-link>
            <span class="card-title">{{v.day}} {{v.time}} @ {{v.branch.name}}</span>
            <button @click="confirmActivate(v.id)" class="btn-floating halfway-fab waves-effect waves-light red"><i
              class="material-icons">check</i></button>
          </div>
          <div class="card-content">
            <div class="class__daytime">
              <i class="material-icons">date_range</i> {{v.day}}
              <i class="material-icons">schedule</i>{{v.time}}
            </div>
            <div class="class__location"><i class="material-icons">place</i>{{v.branch.name}}</div>
          </div>
        </div>
      </div>


      <div id="modal1" class="modal bottom-sheet">
        <div class="modal-content"> Are you sure ?</div>
        <div class="modal-footer">
          <button class="modal-action modal-close waves-effect waves-green btn-flat">No</button>
          <button class="waves-effect waves-green btn-flat" @click="activate">Yes</button>
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
        id: 0,
        classes: [],

        $modalElement: {},
      }
    },
    methods: {
      getSchedules(page = 1) {
        const url = `${process.env.API}/classes?day=today`;

        axios.get(url)
          .then(response => {
            this.classes = response.data._embedded;

            this.classes.forEach((v, i, a) => {
              this.$set(a[i], 'branch', {'name': ''});
              axios.get(`${process.env.API}${v._links.branch.href}`)
                .then(response => {
                  this.$set(a[i], 'branch', response.data);
                })
                .catch(error => console.log(error))
            });
            // console.log(this.classes[0].module.name);
          })
          .catch(error => console.log(error))

      },
      confirmActivate(id, e) {
        this.id = id;
        this.$modalElement = $(this.$el).find('.modal').modal();
        this.$modalElement.modal('open');
      },
      activate() {
        const url = `${process.env.API}/classes/${this.id}/sessions`;
        axios.post(url, {
          id: 3
        })
          .then(response => {
            console.log(response);
            this.$modalElement.modal('close');
            Materialize.toast('Class activated', 4000);
          })
          .catch(error => console.log(error))
        console.log(this.id);
      }
    },
    mounted() {
      this.getSchedules()
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "~materialize-css/sass/components/color";
  @import "~sass-bem";

  main,
  section {
    min-height: 100%;
    height: 100%;
  }

  header {
  }

  main {
    padding: 0.75rem 0;
  }

  .class {
    @include e(daytime) {
      text-transform: capitalize;
    }
  }

  .btn-floating i {
    color: #fff;
  }

  .card {
    height: auto;

    .card-image {
      max-height: 12rem;
      position: relative;
      .card-title {
        text-shadow: 0 0 1.5rem #000;
        text-transform: capitalize;
        background-color: transparentize(color('grey', 'darken-1'), .5);
        padding: 0 1rem;
      }
    }
    .btn-floating.halfway-fab {
      bottom: 0.5rem;
      right: 0.5rem;
    }
    .card-content {
      padding: 1rem;
      font-size: 1rem;
      line-height: 1.25rem;
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
    .material-icons {
      font-size: 1.25rem;
      margin-right: 0.25rem;
      vertical-align: middle;
    }
  }

  .input-field {
    margin: {
      top: 1rem;
      bottom: 1rem;
    }
  }

  input[type=text]:not(.browser-default) {
    border-bottom: none;
    margin-bottom: 0;
    height: 2rem;
  }
</style>
