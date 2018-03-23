import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'
import Dashboard from '@/pages/Dashboard'
import Schedules from '@/pages/Schedules'
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
      component: Dashboard
    },
    {
      path: '/schedules',
      component: Schedules
    },
    {
      path: '/activities',
      component: ComingSoon
    },
    {
      path: '/modules',
      component: ComingSoon
    },
    {
      path: '/branches',
      component: ComingSoon
    },
    {
      path: '/classes',
      component: ComingSoon
    },
    {
      path: '/students',
      component: ComingSoon
    },
    {
      path: '/tutors',
      component: ComingSoon
    },
    {
      path: '/Sign',
      component: Sign
    }
  ]
})
