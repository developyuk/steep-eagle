<template lang="pug">
  transition(enter-active-class="animated fadeIn" leave-active-class="animated fadeOut")
    #sign.grid-y.grid-frame.align-center
      .cell.auto
        form(@submit.prevent="sign").login
          h1.logo M

          .mdc-text-field.mdc-text-field--fullwidth.mdc-text-field--dense
            input#my-text-field.mdc-text-field__input(v-model.trim="username" name="username" type="text" @focus="errMsg=null" placeholder="Enter your name" aria-label="Enter your name")

          transition(enter-active-class="animated bounceInDown" leave-active-class="animated bounceOutUp")
            .errMsg(v-if="!!errMsg") {{errMsg}}
          button.mdc-fab(aria-label="Login")
            span.mdc-fab__icon.material-icons arrow_forward
          .faq
            .faq__item
              a(href="#" @click.prevent="onClickFaq($event,0)") Why no password?
            .faq__item
              a(href="#" @click.prevent="onClickFaq($event,1)") I'm not sure if i registered?
      .cell.shrink
        .powered powered by
          img(alt="codingcamp.id" src="https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png")

      my-dialog(v-model="dialog")
        span(v-html="dialogText.opts[dialogText.idx]")
        template(slot="actions")
          button.mdc-button.mdc-dialog__button(data-mdc-dialog-action="yes" type="button") Got It!
</template>

<script>
import axios from "axios";
import { MDCRipple } from "@material/ripple";
import { MDCTextField } from "@material/textfield";
import { MDCFloatingLabel } from "@material/floating-label";
import MyDialog from "@/components/Dialog";

export default {
  components: {
    MyDialog
  },
  data() {
    return {
      username: "",
      errMsg: "",
      dialog: null,
      dialogText: {
        idx: null,
        opts: [
          "We understand that you value security the most and at some point, it might be ackward to use an app without knowing security is protecting us or not, we actually tried our very best to simplify your tasks and remembering password is one of em.",
          "Please contact administrator for this matters. Don’t forget to add @cc after your name on the login form. Example: kurie@cc"
        ]
      }
    };
  },
  methods: {
    sign() {
      const url = `${process.env.VUE_APP_API}/sign`;
      const username = this.username.substring(0, this.username.length - 3);
      const data = {
        username: username,
        role: ["tutor"]
      };
      this.errMsg = "";

      axios
        .post(url, data)
        .then(response => {
          localStorage.setItem("token", response.data.token);
          // console.log(this.$router);
          this.$router.push("/");
        })
        .catch(error => {
          console.log(error);
          const { status, data } = error.response;
          if (status !== 200) {
            this.errMsg = `Check your username.`;
          } else {
            console.log(error);
          }
        });
    },
    onClickFaq(e, i) {
      this.dialogText.idx = i;
      this.dialog.open();
    }
  },
  mounted() {
    new MDCTextField(this.$el.querySelector(".mdc-text-field"));
    // new MDCFloatingLabel(this.$el.querySelector(".mdc-floating-label"));
    new MDCRipple(this.$el.querySelector(".mdc-fab"));
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
@import "src/assets/shared";
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

  .mdc-text-field {
    width: 100%;
  }
  // .logo {
  // }
}

button {
  margin-top: 2rem;
  float: right;
  @include mdc-fab-accessible(map-get($palettes, red));
}

.faq {
  margin-top: 2rem;
  @include e(item) {
    margin: 0.75rem 0;
    &,
    a {
      color: #9e9e9e;
      font-size: 0.875rem;
      text-decoration: none
    }
  }
}

.powered {
  height: 1rem;
  margin: 0.5rem 0;
  text-align: center;
  font-size: 0.75rem;
  img {
    margin: 0 0.5rem;
  }
}
</style>
