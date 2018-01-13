import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'
import Classes from '@/pages/Classes'
import Sign from '@/pages/Sign'
import ComingSoon from '@/components/ComingSoon'

Vue.use(Router)

export default new Router({
  routes: [
    // {
    //   path: '/',
    //   name: 'Hello',
    //   component: Hello
    // }

    {
      path: '/',
      component: Classes
    },
    {
      path: '/Sign',
      component: Sign
    }
  ]
})
