<template>
<main id="app">
  <router-view></router-view>
</main>
</template>

<script>
import axios from 'axios'
axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');

export default {
  name: 'app'
}
</script>

<style lang="scss">
$roboto-font-path: "~materialize-css/dist/fonts/roboto/";
@import "~materialize-css/sass/components/color";
$primary-color: color("grey", "base");
$secondary-color: color("grey", "lighten-1");
@import "~materialize-css/sass/materialize.scss";

#app,
body,
html {
    height: 100%;
    min-height: 100%;
}
body,
html {
    &,
    .material-icons,
    a {
        color: color("grey", "darken-1");
    }
}
body {
    background-color: color('grey','lighten-5');
}
header {
    background-color: #fff;
    overflow: auto;
}
#app {
    margin: 0 auto;
    // border: thin solid color('grey', 'lighten-1');
    &,
    * {
        max-width: 678px;
    }
}
.modal {
    left: 50%;
    transform: translateX(-50%);
}
</style>
