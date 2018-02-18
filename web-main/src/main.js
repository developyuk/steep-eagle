// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import * as mdc from 'material-components-web';
import axios from "axios";
// import EventBus from './event-bus.js';

Vue.config.productionTip = false;
require("normalize.css/normalize.css");
require("material-components-web/dist/material-components-web.min.css");
require("material-components-web/dist/material-components-web.min.js");
window.mdc = mdc;

router.beforeEach((to, from, next) => {

  axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      // console.log(response.data);
      setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 144);
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
