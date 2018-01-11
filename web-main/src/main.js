// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'

Vue.config.productionTip = false
require('normalize.css/normalize.css')

router.beforeEach((to, from, next) => {
  // ...
  const isLogin = localStorage.getItem('token');
  // console.log(isLogin, to.path,from.path, !isLogin && to.path !== '/sign');
  if (!isLogin && to.path !== '/sign') {
    next({
      path: '/sign',
      query: {
        redirect: to.path
      }
    })
  }
  next();
})
window.$ = window.jQuery = require('jquery')
require('materialize-css/dist/js/materialize.min.js')


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
