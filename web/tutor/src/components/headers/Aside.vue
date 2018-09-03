<template lang="pug">
  #header
    header
      .grid-x.align-middle
        .cell.shrink
          i.material-icons(@click="$router.push('/')") arrow_back
        .cell.auto
          .input
            input(type="text" v-model='q' placeholder='Search')
        .cell.shrink
          i.material-icons(@click="$router.push('/')") close

</template>

<script>
  import {mapState} from 'vuex'

  export default {
    components: {
    },
    computed: {
      ...mapState(['currentAuth', 'currentStats']),
    },
    data() {
      return {
        q: '',
      }
    },
    methods: {
    },
    destroyed() {
    },
    mounted() {
      this.$el.querySelector('input').focus()
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
  @import "../../assets/shared";
  @import "~sass-bem";

  header {
    position: relative;
    background-color: $mdc-theme-primary;
    overflow: hidden;
    /*height: 56px;*/
    padding:.875rem;
  }

  .material-icons{
    color: #fff;
  }
  .input{
    margin:0 1rem;
    input{
      color:#fff;
      background-color: transparent;
      border: none;
      width: 100%;
      height: 100%;
      padding: .5rem;
    }
  }
</style>
