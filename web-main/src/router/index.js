import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'
import Schedules from '@/pages/Schedules'
import ClassDetail from '@/pages/ClassDetail'
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
      name: 'Schedules',
      component: Schedules
    },
    {
      path: '/class/:id',
      name: 'ComingSoon',
      component: ComingSoon
    },
    {
      path: '/feedbacks',
      name: 'ComingSoon',
      component: ComingSoon
    },
    {
      path: '/progress',
      name: 'ComingSoon',
      component: ComingSoon
    },
    {
      path: '/profile',
      name: 'ComingSoon',
      component: ComingSoon
    }
  ]
})
