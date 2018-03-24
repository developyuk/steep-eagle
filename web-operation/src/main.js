// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";
import axios from "axios";
import * as mdc from 'material-components-web/dist/material-components-web';

Vue.config.productionTip = false;
require("normalize.css/normalize.css");
window.mdc = mdc;

router.beforeEach((to, from, next) => {
  axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      next();
    })
    .catch(error => {
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

Vue.prototype.$bus = new Vue({});
/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  template: "<App/>",
  components: {
    App
  }
});
