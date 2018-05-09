// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store/index'
import axios from "axios";

import * as mdc from 'material-components-web/dist/material-components-web';

window.mdc = mdc;
require("normalize.css/normalize.css");

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => {
  if (localStorage.getItem('token')) {
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  } else {
    axios.defaults.headers.common['Authorization'] = '';
  }
  // axios.defaults.headers.post['Content-Type'] = 'application/json; charset=utf-8';
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      // console.log(response.data);
      const currentAuth = response.data;
      store.commit('nextAuth', currentAuth);
    })
    .catch(error => {
      // console.log(error.response, to.path);
      if (to.matched.some(record => record.meta.requiresAuth)) {
        //   next({path: "/sign", query: {redirect: to.path}});
        next({path: "/sign"});
      }
      // next({path: "/sign"});
    });
  next();
});

Vue.prototype.$bus = new Vue({});
/* eslint-disable no-new */

new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {App}
})
