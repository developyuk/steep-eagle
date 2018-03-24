<template lang="pug">
  #sign.mdc-layout-grid
    main.mdc-layout-grid__inner
      section.left.mdc-layout-grid__cell.mdc-layout-grid__cell--span-4
        form.login
          img.logo(src="https://images.weserv.nl/?crop=110,95,88,107&url=dl.dropboxusercontent.com/s/psvta5uwq4s0m5y/logo2.jpg")

          .mdc-form-field
            .mdc-text-field(data-mdc-auto-init="MDCTextField")
              input#id.mdc-text-field__input(v-model.trim="username" name="username" type="text" required)
              label.mdc-text-field__label(for="id") Yourname
              .mdc-line-ripple
          .mdc-form-field
            button(type="submit").mdc-button.mdc-button--raised login
        .powered powered by
          img(src="https://images.weserv.nl/?h=10&url=dl.dropboxusercontent.com/s/htl2v26j5imxgxa/Group.png")
      section.right.mdc-layout-grid__cell.mdc-layout-grid__cell--span-8
        .container
          .container__text Scheduling never been this easy.
          .container__text Manage with M.
</template>

<script>
  export default {
    name: 'sign',
    data() {
      return {
        msg: 'Welcome to Your Vue.js PWA'
      }
    },
    mounted() {
      window.mdc.autoInit();
    },
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../assets/shared";

  #sign {
    padding: 0;
    height: 100vh;
  }

  main {
    height: 100%;
  }

  form {
  }

  .left {
    position: relative;
  }

  .right {
    background-image: url(//images.weserv.nl/?output=jpg&q=21&url=dl.dropboxusercontent.com/s/asuxh2qp5zh79uu/Rectangle-compressor.png);
    position: relative;
    .container {

      &__text {
        color: #fff;
        font-size: 3rem;
        font-weight: 900;
      }
      /*padding: 4rem;*/
      margin: 2rem;
      position: absolute;
      top: 50%;
      transform: translateY(-50%);
    }
  }
</style>
