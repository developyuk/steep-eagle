<template>
<section>
  <main>
    <h3>coming soon</h3>
  </main>
  <footer>
    <nav-bottom></nav-bottom>
  </footer>
</section>
</template>

<script>
import NavBottom from '@/components/NavBottom'

export default {
  components: {
    'nav-bottom': NavBottom
  },
  name: 'coming-soon',
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
