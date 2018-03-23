<template>
<main>
  <form class="container" @submit.prevent="sign">
    <div class="input-field">
      <i class="material-icons prefix">account_circle</i>
      <input id="email" type="email" v-model.trim="email" class="validate" placeholder="Enter your email">
    </div>
    <div class="input-field">
      <i class="material-icons prefix">lock_outline</i>
      <input id="password" type="password" v-model.trim="pwd" class="validate" placeholder="Enter your password">
    </div>
    <button class="btn waves-effect waves-light" type="submit">Sign in </button>
  </form>

  <div id="modal1" class="modal bottom-sheet">
    <div class="modal-content"> </div>
    <div class="modal-footer">
      <button class="modal-action modal-close waves-effect waves-green btn-flat">Close</button>
    </div>
  </div>
</main>
</template>

<script>
import axios from 'axios'

export default {
  name: 'sign',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA',
      email: 'user7@ex.com',
      pwd: 'asdqwes'
    }
  },
  methods: {
    sign() {
      const url = `${process.env.API}/sign`;
      const data = {
        email: this.email,
        pwd: this.pwd,
      };

      axios.post(url, data)
        .then(response => {
          localStorage.setItem('token', response.data.token);
          // console.log(this.$router);
          this.$router.push(this.$route.query.redirect);
        })
        .catch(error => {
          console.log(error);
          const {
            status,
            data
          } = error.response;
          if (status === 401) {
            const message = `${data.message}<br/> Check your authentication.`;

            $(this.$el).find('.modal').modal()
              .find('.modal-content').html(message).end()
              .modal('open');
          } else {
            console.log(error);
          }
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.container {

    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%,-50%);
}
button {
    float: right;
}
</style>
