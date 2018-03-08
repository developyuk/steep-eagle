<template lang="pug">
  #sign.mdc-typography
    form(@submit.prevent="sign")
      img.logo(src="https://images.weserv.nl/?w=300&url=dl.dropboxusercontent.com/s/a6mfkumhayvm7o4/Character_CodingCamp.png")
      .mdc-form-field
        .mdc-text-field.mdc-text-field--box.mdc-text-field--with-leading-icon(data-mdc-auto-init="MDCTextField")
          i.material-icons.mdc-text-field__icon(tabindex="0") account_circle
          input#id.mdc-text-field__input(v-model.trim="id" type="text" required)
          label.mdc-text-field__label.mdc-text-field__label--float-above(for="id") Enter your email
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
        id: 'user3@ex.co',
        errMsg: ''
      }
    },
    mounted() {
//      Array.from(document.querySelectorAll('.mdc-text-field')).forEach(v => MDCRipple.attachTo(v));
      mdc.ripple.MDCRipple.attachTo(document.querySelector('.mdc-button'));
    },
    methods: {
      sign() {
        const url = `${process.env.API}/sign`;
        const data = {
          email: this.id,
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
  .logo {
    width: 300px;
    margin: 0 auto 1rem auto;
    display: block;
  }

  form {
    position: absolute;
    top: calc(50% - 2rem) ;
    transform: translateY(-50%);
    max-width: 30rem;
  }
  button{
    margin-top: 2rem;
  }

  .mdc-form-field {
    display: block;
  }

  form, .mdc-text-field, .mdc-button {
    width: 100%;
  }

</style>
