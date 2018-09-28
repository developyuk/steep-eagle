<template lang="pug">
  .wrapper.wrapper-full-page
    .full-page.login-page(data-color='' data-image='https://images.weserv.nl/?il&w=1024&h=768&t=square&url=dl.dropboxusercontent.com/s/y91mai1ns2bchvh/M-Ops-Login.jpg')
      // you can change the color of the filter page using: data-color="blue | azure | green | orange | red | purple"
      .content
        .container
          .row
            .col-md-4.col-sm-6.col-md-offset-4.col-sm-offset-3
              form(method='#' action='#' @submit.prevent='onSubmit')
                .card(data-background='color' data-color='blue')
                  .card-header
                    //h3.card-title Login
                    h1.logo.card-title M

                  .card-content
                    .form-group
                      label Username or Email
                      input.form-control.input-no-border(type='text' name='username' placeholder='Enter username or email' v-model='username')
                    .err(v-if="errMsg") {{errMsg}}
                    .success(v-if="successMsg") {{successMsg}}
                  .card-footer.text-center
                    router-link.btn.btn-simple(to='/sign') Back
                    | &nbsp;&nbsp;&nbsp;
                    button.btn.btn-fill.btn-wd(type='submit') Submit
      //- .full-page-background(style='background-image: url(https://images.weserv.nl/?il&w=1024&h=768&t=square&url=dl.dropboxusercontent.com/s/y91mai1ns2bchvh/M-Ops-Login.jpg) ')
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      errMsg: null,
      successMsg: null,
      username: null,
      password: null
    };
  },
  methods: {
    toggleNavbar() {
      document.body.classList.toggle("nav-open");
    },
    closeMenu() {
      document.body.classList.remove("nav-open");
      document.body.classList.remove("off-canvas-sidebar");
    },
    onSubmit(e) {
      this.errMsg = null;
      const formData = new FormData(e.target);
      const data = {
        username: this.username,
        role: ["operation"]
      };
      axios
        .post(`${process.env.API}/forgot_password`, data)
        .then(response => {
          this.successMsg =
            "email has been sent to. You will be redirected to login page";
          setTimeout(() => {
            this.$router.push("/sign");
          }, 1000);
        })
        .catch(error => {
          console.log(error, error.response);
          this.errMsg = error.response.data._error.message;
        });
    }
  },
  beforeDestroy() {
    this.closeMenu();
  }
};
</script>

<style scoped lang="scss">
@import "src/assets/scss/shared";

// .card {
//   background-color: #f2f2f2;
// }

.card-header {
  h1 {
    text-align: center;
  }
}

// .form-control {
//   &,
//   &:focus {
//     background-color: #c4c4c4;
//     color: #fff;
//   }
// }

// ::placeholder {
//   color: #fff;
//   font-weight: bold;
//   text-transform: lowercase;
// }
.full-page {
  background-color: #ee215b;
  height: 100vh;
}
.full-page[data-image]:after,
.full-page.has-image:after {
  opacity: 0;
}
.err {
  background-color: pink;
  color: red;
  padding: 1rem;
  text-transform: capitalize;
}
</style>
