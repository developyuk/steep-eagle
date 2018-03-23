<template lang="pug">
.sidebar
	ul
		li
			router-link(to="/")
				i.material-icons assignment
				span Dashboard
		li
			router-link(to="/schedules")
				i.material-icons assignment
				span Schedules
		li
			router-link(to="/activities")
				i.material-icons assignment
				span Activities
		<!--li-->
			<!--i.material-icons assignment-->
			<!--span Data Space-->
			<!--ul-->
				<!--li-->
					<!--router-link(to="/branches")-->
						<!--i.material-icons assignment-->
						<!--span Branches-->
				<!--li-->
					<!--router-link(to="/modules")-->
						<!--i.material-icons assignment-->
						<!--span Modules-->
				<!--li-->
					<!--router-link(to="/classes")-->
						<!--i.material-icons assignment-->
						<!--span Classes-->
		<!--li-->
			<!--i.material-icons assignment-->
			<!--span User Space-->
			<!--ul-->
				<!--li-->
					<!--router-link(to="/students")-->
						<!--i.material-icons assignment-->
						<!--span Students-->
				<!--li-->
					<!--router-link(to="/tutors")-->
						<!--i.material-icons assignment-->
						<!--span Tutors-->
</template>

<script>
export default {
  name: "nav-side",
  data() {
    return {
      msg: "Welcome to Your Vue.js PWA"
    };
  },
  methods: {
    isActivePath(path) {
      return this.$route.path === path;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.sidebar {
  ul {
    padding-left: 1rem;
  }
  .material-icons {
    vertical-align: middle;
  }
}
</style>
