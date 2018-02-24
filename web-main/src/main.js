// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from "axios";
// require("material-components-web/dist/material-components-web.min.js");
import * as mdc from 'material-components-web/dist/material-components-web';
// import EventBus from './event-bus.js';

Vue.config.productionTip = false;
require("normalize.css/normalize.css");
require("hammerjs");

window.mdc = mdc;

router.beforeEach((to, from, next) => {

  axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      // console.log(response.data);
      setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99*8);
      next();
    })
    .catch(error => {
      // console.log(error.response, to.path);
      if (to.matched.some(record => record.meta.requiresAuth)) {
        next({path: "/sign", query: {redirect: to.path}});
      } else {
        next();
      }
    });
});

Vue.prototype.$bus = new Vue({});
/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {App}
})
