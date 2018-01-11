<template>
<router-view></router-view>
</template>

<script>
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

body,
html {
    height: 100%;
    min-height: 100%;
    &,
    .material-icons,
    a {
        color: color("grey", "darken-1");
    }
}
.material-icons {
    // color: $primary-color;
}
</style>
