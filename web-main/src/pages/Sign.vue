<template lang="pug">
  #sign.mdc-typography
    form(@submit.prevent="sign")
      .mdc-form-field
        .mdc-text-field.mdc-text-field--box.mdc-text-field--with-leading-icon(data-mdc-auto-init="MDCTextField")
          i.material-icons.mdc-text-field__icon(tabindex="0") account_circle
          input#id.mdc-text-field__input(v-model.trim="id" type="text" required)
          label.mdc-text-field__label.mdc-text-field__label--float-above(for="id") Enter your email
          .mdc-line-ripple
      .mdc-form-field
        .mdc-text-field.mdc-text-field--box.mdc-text-field--with-leading-icon(data-mdc-auto-init="MDCTextField")
          i.material-icons.mdc-text-field__icon(tabindex="1") lock_outline
          input#pwd.mdc-text-field__input(v-model.trim="pwd" type="password" required)
          label.mdc-text-field__label.mdc-text-field__label--float-above(for="pwd") Enter your password
          .mdc-line-ripple
      button(type="submit" data-mdc-auto-init="MDCRipple").mdc-button.mdc-button--raised Sign in
</template>

<script>
  import axios from 'axios';
  import {MDCRipple} from '@material/ripple';

  export default {
    name: 'sign',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA',
        id: 'user3@ex.com',
        pwd: 'asdqwes'
      }
    },
    mounted() {
      Array.from(document.querySelectorAll('.mdc-text-field')).forEach(v => MDCRipple.attachTo(v));
      MDCRipple.attachTo(document.querySelector('.mdc-button'));
    },
    methods: {
      sign() {
        const url = `${process.env.API}/sign`;
        const data = {
          email: this.id,
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

//              $(this.$el).find('.modal').modal()
//                .find('.modal-content').html(message).end()
//                .modal('open');
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

  form {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    width: 30rem;
  }

  .mdc-form-field {
    display: block;
  }

  .mdc-text-field, .mdc-button {
    width: 100%;
  }

</style>
