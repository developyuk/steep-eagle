<template lang="pug">
section
	nav-side
	h3 coming soon
</template>

<script>
import NavSide from '@/components/NavSide'
export default {
  name: 'coming-soon',
  components: {
    'nav-side': NavSide
  },
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
    text-transform: capitalize;
    text-align: center;
    position: absolute;
    top: 42%;
    transform: translateY(-50%);
    width: 100%;
}
</style>
