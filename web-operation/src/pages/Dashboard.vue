<template lang="pug">
section.row
	section.col.s3
		nav-side
	main.col.s9
		header
			nav
				div.nav-wrapper
					ul.row.right
						li.col.s3
							a(href="sass.html"): i.material-icons search
						li.col.s3
							a(href="badges.html"): i.material-icons view_module
						li.col.s3
							a(href="collapsible.html"): i.material-icons refresh
						li.col.s3
							a(href="mobile.html"): i.material-icons more_vert
		section main
</template>

<script>
import NavSide from '@/components/NavSide'
export default {
  name: 'dashboard',
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
$roboto-font-path: "~materialize-css/dist/fonts/roboto/";
@import "~materialize-css/sass/components/color";
$primary-color: color("grey", "base");
$secondary-color: color("grey", "lighten-1");
@import "~materialize-css/sass/materialize.scss";
nav {
    background-color: transparent;
    .material-icons {
        color: $primary-color;
    };
}
</style>
