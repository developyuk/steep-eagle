import Vue from 'vue'
import './pollyfills'
import VueNotify from 'vue-notifyjs'
import VeeValidate from 'vee-validate'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import App from './App.vue'

// Plugins
import GlobalComponents from './gloablComponents'
import GlobalDirectives from './globalDirectives'
import SideBar from './views/SidebarPlugin'

// router setup
import router from './routes/router'

// store setup
import store from './store/index'

// library imports

import './assets/sass/paper-dashboard.scss'
import './assets/sass/element_variables.scss'
import './assets/sass/demo.scss'

import sidebarLinks from './sidebarLinks'
// plugin setup
Vue.use(GlobalDirectives)
Vue.use(GlobalComponents)
Vue.use(VueNotify)
Vue.use(SideBar, {sidebarLinks: sidebarLinks})
Vue.use(VeeValidate)
locale.use(lang)


/* eslint-disable no-new */
new Vue({
  el: '#app',
  render: h => h(App),
  router,
  store
})
