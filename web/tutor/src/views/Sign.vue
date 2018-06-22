<template lang="pug">
  template-empty#sign
    form(@submit.prevent="sign").login
      img.logo(src="https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg")
      .mdc-text-field
        input#my-text-field.mdc-text-field__input(v-model.trim="username" name="username" type="text" required)
        label.mdc-text-field__label(for="my-text-field") Yourname
        .mdc-line-ripple
      .errMsg(v-if="errMsg" v-html="errMsg")
      button(type="submit").mdc-button.mdc-button--raised Login
    .powered powered by
      img(src="https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png")
</template>

<script>
  import axios from 'axios';
  import {MDCRipple} from '@material/ripple';
  import {MDCTextField} from '@material/textfield';
  import TemplateEmpty from '@/templates/TemplateEmpty';

  export default {
    components: {
      'template-empty': TemplateEmpty,
    },
    data() {
      return {
        username: '',
        errMsg: ''
      }
    },
    mounted() {
      const $textfield =new MDCTextField(this.$el.querySelector('.mdc-text-field'));
      const $button = new MDCRipple(this.$el.querySelector('.mdc-button'));
    },
    methods: {
      sign() {
        const url = `${process.env.VUE_APP_API}/sign`;
        const data = {
          username: this.username.toLowerCase(),
        };

        axios.post(url, data)
          .then(response => {
            localStorage.setItem('token', response.data.token);
            // console.log(this.$router);
            this.$router.push('/');
          })
          .catch(error => {
            console.log(error);
            const {status, data} = error.response;
            if (status !== 200) {
              this.errMsg = `${data.message}.<br/> Re-check your authentication.`;
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
  #sign {
  }

  form.login {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -55%);
    max-width: 30rem;
    width: calc(100% - 5rem);
  }

  button {
    margin-top: 2rem;
  }
</style>
