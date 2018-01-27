// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import axios from "axios";

Vue.config.productionTip = false;
require("normalize.css/normalize.css");

router.beforeEach((to, from, next) => {
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      // console.log(response.data);
      next();
    })
    .catch(error => {
      // console.log(error.response, to.path);
      if (to.path !== "/sign") {
        next({
          path: "/sign",
          query: {
            redirect: to.path
          }
        });
      }
      next();
    });
});
window.$ = window.jQuery = require("jquery");
require("materialize-css/dist/js/materialize.min.js");

/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  template: "<App/>",
  components: {
    App
  }
});
