<template lang="pug">
  #empty.empty
    //img(src="https://images.weserv.nl/?il&crop=400,300,50,100&url=dl.dropboxusercontent.com/s/zlm9ycfk78jh20w/no-student.png")
    h1.mdc-typography--headline6 No class started
    p.mdc-typography--body2 Make sure to fill your student attendance by start or activate your schedule first. Also don't forget to rate and review your student as well
    router-link(to="/").mdc-button.mdc-button--raised Start Class
</template>

<script>
  import {MDCRipple} from '@material/ripple';

  export default {
    mounted() {
      new MDCRipple(this.$el.querySelector('.mdc-button'));
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
  @import "src/assets/shared";

  .empty {
    text-align: center;
    // position: absolute;
    // top: 50%;
    // transform: translateY(-60%);
    margin: 2rem;
    width: calc(100% - 4rem);

    h1 {
      color: #58595B;
    }
    p {
      color: #BDBDBD;
    }
    img{
      width: 100%;
    }
  }

  .mdc-button {
    background-color: map-get($palettes, green);
    width: 100%;
  }

</style>
