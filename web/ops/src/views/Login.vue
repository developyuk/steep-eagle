<template lang="pug">
  <!--nav.navbar.navbar-transparent.navbar-absolute-->
  <!--.container-->
  <!--.navbar-header-->
  <!--button.navbar-toggle(type='button' data-toggle='collapse' data-target='#navigation-example-2' @click='toggleNavbar')-->
  <!--span.sr-only Toggle navigation-->
  <!--span.icon-bar-->
  <!--span.icon-bar-->
  <!--span.icon-bar-->
  <!--router-link.navbar-brand(to='/admin') Paper Dashboard PRO-->
  <!--.collapse.navbar-collapse-->
  <!--ul.nav.navbar-nav.navbar-right-->
  <!--router-link(to='/register' tag='li')-->
  <!--a Register-->
  <!--router-link(to='/admin/overview' tag='li')-->
  <!--a Dashboard-->
  .wrapper.wrapper-full-page
    .full-page.login-page(data-color='' data-image='https://images.weserv.nl/?il&q=18&url=dl.dropboxusercontent.com/s/y91mai1ns2bchvh/M-Ops-Login.jpg')
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
                      //label Email address
                      input.form-control.input-no-border(type='text' name='username' placeholder='Enter email')
                    .form-group
                      //label Password
                      input.form-control.input-no-border(type='password' name='password' placeholder='Password please')
                  .card-footer.text-center
                    button.btn.btn-fill.btn-wd(type='submit') Let&apos;s go
                    <!--.forgot-->
                    <!--router-link(to='/register')-->
                    <!--| Forgot your password?-->
      footer.footer.footer-transparent
        <!--.container-->
        <!--.copyright-->
        <!--| &copy; Coded with-->
        <!--i.fa.fa-heart.heart-->
        <!--|  by-->
        <!--a(href='https://github.com/cristijora' target='_blank') Cristi Jora-->
        <!--| .-->
        <!--|               Designed by-->
        <!--a(href='https://www.creative-tim.com/?ref=pdf-vuejs' target='_blank') Creative Tim-->
        <!--| .-->
      .full-page-background(style='background-image: url(https://images.weserv.nl/?il&q=18&url=dl.dropboxusercontent.com/s/y91mai1ns2bchvh/M-Ops-Login.jpg) ')
  <!--.collapse.navbar-collapse.off-canvas-sidebar-->
  <!--ul.nav.nav-mobile-menu-->
  <!--router-link(to='/register' tag='li')-->
  <!--a Register-->
  <!--router-link(to='/admin/overview' tag='li')-->
  <!--a Dashboard-->
</template>

<script>
  import axios from 'axios';

  export default {
    methods: {
      toggleNavbar() {
        document.body.classList.toggle('nav-open')
      },
      closeMenu() {
        document.body.classList.remove('nav-open')
        document.body.classList.remove('off-canvas-sidebar')
      },
      onSubmit(e) {
        const formData = new FormData(e.target);

        axios.post(`${process.env.DBAPI}/sign`, formData)
          .then(response => {
//            console.log(response);
            localStorage.setItem('token', response.data.token);
            // console.log(this.$router);
            this.$router.push('/');
          })
          .then(error => {
            console.log(error);
          })
      }
    },
    beforeDestroy() {
      this.closeMenu()
    }
  }
</script>

<style scoped lang="scss">
  @import "../assets/scss/shared";

  .card {
    background-color: #F2F2F2;
  }

  .card-header {
    h1 {
      text-align: center;
    }
  }

  .full-page[data-image]:after, .full-page.has-image:after {
    opacity: 0;
  }
</style>
