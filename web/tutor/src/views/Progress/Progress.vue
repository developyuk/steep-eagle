<template lang="pug">
  template-main#progress
    .empty
      img(src="https://images.weserv.nl/?il&crop=400,100,0,80&url=dl.dropboxusercontent.com/s/n7l2woy0sqh09f8/no-progress.png")
      h1.mdc-typography--headline6 No record
      p.mdc-typography--body2 It seems you haven't rated your students from your class. Please make sure to do so
      router-link(to="/students").mdc-button.mdc-button--raised Rate Students
</template>

<script>
  import TemplateMain from '@/components/views/Main';

  export default {
    components: {
      TemplateMain,
    },
    data() {
      return {
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "../../assets/shared";

  .empty {
    position: absolute;
    top: 50%;
    /*left: 50%;
    transform: translate(-50%, -50%);*/
    transform: translateY(-60%);
    margin: 2rem;
    width: calc(100% - 4rem);
    text-align: center;

    h1 {
      color: #58595B;
    }
    p {
      color: #BDBDBD;
    }
    img {
      width: 100%;
    }
  }
  .mdc-button {
    background-color: map-get($palettes, green);
    width: 100%;
  }
</style>
