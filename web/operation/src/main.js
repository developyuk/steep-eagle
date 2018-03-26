// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from "vue";
import App from "./App";
import router from "./router";

import sharedVue from "./assets/vue";
import * as mdc from 'material-components-web/dist/material-components-web';
window.mdc = mdc;
require("normalize.css/normalize.css");

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => sharedVue.routerBeforeEach(router, to, from, next));

Vue.prototype.$bus = new Vue({});
/* eslint-disable no-new */
new Vue({
  el: "#app",
  router,
  template: "<App/>",
  components: {App}
});
