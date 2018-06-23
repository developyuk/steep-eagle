<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-class="animated fadeOut")
    #sign
      form(@submit.prevent="sign").login
        img.logo(src="https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg")
        .mdc-text-field
          input#my-text-field.mdc-text-field__input(v-model.trim="username" name="username" type="text")
          label.mdc-floating-label(for="my-text-field") Enter your name
          .mdc-line-ripple

        transition(enter-active-class="animated fadeInUp" leave-class="animated fadeOutDown")
          .errMsg(v-if="errMsg" v-html="errMsg")
        button.mdc-fab(aria-label="Login")
          span.mdc-fab__icon.material-icons arrow_forward
        .faq
          .faq__item
            a(@click="onClickFaq($event,0)") Why no password?
          .faq__item
            a(@click="onClickFaq($event,1)") I'm not sure if i registered?
      .powered powered by
        img(src="https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png")
      my-dialog(@mounted="onDialogMounted")
        span(v-html="dialogText.opts[dialogText.idx]")
        template(slot="footer")
          button.mdc-button.mdc-dialog__footer__button.mdc-dialog__footer__button--accept(type="button") Got It!
</template>

<script>
  import axios from 'axios';
  import {MDCRipple} from '@material/ripple';
  import {MDCTextField} from '@material/textfield';
  import {MDCFloatingLabel} from '@material/floating-label';
  import MyDialog from '@/components/Dialog';

  export default {
    components: {
      MyDialog
    },
    data() {
      return {
        username: '',
        errMsg: '',
        $dialog: null,
        dialogText: {
          idx: null,
          opts: [
            'We understand that you value security the most and at some point, it might be ackward to use an app without knowing security is protecting us or not, we actually tried our very best to simplify your tasks and remembering password is one of em.',
            'Please contact administrator for this matters. Don’t forget to add @cc after your name on the login form. Example: kurie@cc']
        }
      }
    },
    methods: {
      sign() {
        const url = `${process.env.VUE_APP_API}/sign`;
        const data = {
          username: this.username.toLowerCase().slice(0, -3),
        };
        this.errMsg = '';

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
      },
      onClickFaq(e, i) {
        this.dialogText.idx = i;
        this.$dialog.show()
      },
      onDialogMounted(e) {
        this.$dialog = e;
      }
    },
    mounted() {
      new MDCTextField(this.$el.querySelector('.mdc-text-field'));
      new MDCFloatingLabel(this.$el.querySelector('.mdc-floating-label'));
      new MDCRipple(this.$el.querySelector('.mdc-fab'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../assets/shared";
  @import "~sass-bem";
  @import "~@material/fab/mixins";

  #sign {
  }

  form.login {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -55%);
    max-width: 30rem;
    width: calc(100% - 5rem);

    .mdc-text-field, .mdc-button {
      width: 100%;
    }
    .logo {
      width: 5rem;

      margin: 0 auto 4rem auto;
      display: block;
    }
  }

  button {
    margin-top: 2rem;
    float: right;
    @include mdc-fab-accessible(map-get($palettes, red));
  }

  .faq {
    margin-top: 2rem;
    @include e(item) {
      margin: .75rem 0;
      &, a {
        color: #9E9E9E;
        font-size: .75rem;
      }
    }
  }

  .powered {
    position: absolute;
    bottom: 1rem;
    left: 50%;
    transform: translateX(-50%);
    width: 12.025rem;
    font-size: .75rem;
    img {
      margin: 0 .5rem;
    }
  }
</style>
