<template lang="pug">
  #sign.mdc-typography
    form(@submit.prevent="sign")
      img.logo(src="https://images.weserv.nl/?w=300&url=dl.dropboxusercontent.com/s/a6mfkumhayvm7o4/Character_CodingCamp.png")
      .mdc-form-field
        .mdc-text-field(data-mdc-auto-init="MDCTextField")
          input#id.mdc-text-field__input(v-model.trim="id" type="text" required)
          label.mdc-text-field__label(for="id") Yourname
          .mdc-line-ripple
      .mdc-form-field
        .errMsg(v-if="errMsg") {{errMsg}}
        button(type="submit" data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised Login
</template>

<script>
  import axios from 'axios';
  //  import {MDCRipple} from '@material/ripple';

  export default {
    name: 'sign',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
//        id: 'user3@ex.co',
        id: '',
        errMsg: ''
      }
    },
    mounted() {
//      Array.from(document.querySelectorAll('.mdc-text-field')).forEach(v => MDCRipple.attachTo(v));
      window.mdc.autoInit();
    },
    methods: {
      sign() {
        const url = `${process.env.API}/sign`;
        const data = {
          username: this.id,
        };

        axios.post(url, data)
          .then(response => {
            localStorage.setItem('token', response.data.token);
            // console.log(this.$router);
            this.$router.push(this.$route.query.redirect);
          })
          .catch(error => {
            console.log(error);
            const {status, data} = error.response;
            if (status !== 200) {
              this.errMsg = `${data.message}. Re-check your authentication.`;
            } else {
              console.log(error);
            }
          })
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  /*@import '~@material/textfield/mdc-text-field';*/
  /*@import '~@material/button/mdc-button';*/
  #sign {
  }

  .logo {
    /*width: 300px;*/

    margin: 0 auto 4rem auto;
    display: block;
  }

  form {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -55%);
    max-width: 30rem;
    width: 100%;
  }

  button {
    margin-top: 2rem;
  }

  .mdc-form-field {
    display: block;
  }

  .logo, .mdc-text-field, .mdc-button {
    width: 100%;
  }

</style>
