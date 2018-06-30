<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-active-class="animated fadeOut")
    #sign.grid-y.grid-frame.align-center
      .cell.auto
        form(@submit.prevent="sign").login
          h1.logo M
          <!--img.logo(src="https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg")-->
          .mdc-text-field
            input#my-text-field.mdc-text-field__input(v-model.trim="username" name="username" type="text" @focus="errMsg=null")
            label.mdc-floating-label(for="my-text-field") Enter your name
            .mdc-line-ripple

          transition(enter-active-class="animated bounceInUp" leave-active-class="animated bounceOutDown")
            .errMsg(v-if="!!errMsg") {{errMsg}}
          button.mdc-fab(aria-label="Login")
            span.mdc-fab__icon.material-icons arrow_forward
          .faq
            .faq__item
              a(@click="onClickFaq($event,0)") Why no password?
            .faq__item
              a(@click="onClickFaq($event,1)") I'm not sure if i registered?
      .cell.shrink
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
              this.errMsg = `Check your username.`;
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
  @import "~foundation-sites/scss/util/util";

  #sign {
    min-height: 21rem;
  }

  form.login {
    @include absolute-center;
    max-width: 30rem;
    width: calc(100% - 5rem);

    .mdc-text-field{
      width: 100%;
    }
    .logo {

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
        font-size: .875rem;
      }
    }
  }

  .powered {
    height: 1rem;
    margin: .5rem 0;
    text-align: center;
    font-size: .75rem;
    img {
      margin: 0 .5rem;
    }
  }
</style>
