<template>
<nav>
  <!-- {{$route.path}} -->
  <ul class="row">
    <li :class="{col:true, s3:true, 'is-active': isActivePath('/')}">
      <router-link to="/"><i class="material-icons">calendar_today</i></router-link>
    </li>
    <li :class="{col:true, s3:true, 'is-active': isActivePath('/feedbacks')}">
      <router-link to="/feedbacks"><i class="material-icons">assignment</i></router-link>
    </li>
    <li :class="{col:true, s3:true, 'is-active': isActivePath('/progress')}">
      <router-link to="/progress"><i class="material-icons">assessment</i></router-link>
    </li>
    <li :class="{col:true, s3:true, 'is-active': isActivePath('/profile')}">
      <router-link to="/profile"><i class="material-icons">person</i></router-link>
    </li>
  </ul>
</nav>
</template>

<script>
export default {
  name: 'nav-bottom',
  data() {
    return {
      msg: 'Welcome to Your Vue.js PWA'
    }
  },
  methods: {
    isActivePath(path) {
      return this.$route.path === path
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
@import "~materialize-css/sass/components/color";
nav {
    position: fixed;
    bottom: 0;
    background-color: #fff;
    .row {
        margin-bottom: 0;
        .col {
            padding: 0;
        }
    }
    li {
        text-align: center;
        &.is-active {
            position: relative;
            // border-top: 0.25rem solid color("grey", "base");
            &:before {
                content: '';
                position: absolute;
                top: -.25rem;
                left: 0;
                height: 0.25rem;
                width: 100%;
                background-color: color("grey", "darken-1");
            }
        }
    }
    a {
        padding: 0;
    }
    .material-icons {
        font-size: 2.5rem;
    }
}
</style>
