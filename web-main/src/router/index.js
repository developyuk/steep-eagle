import Vue from 'vue'
import Router from 'vue-router'
import Sign from '@/pages/Sign'
import Schedules from '@/pages/Schedules'
import Students from '@/pages/Students'
// import Hello from '@/components/Hello'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Schedules',
      component: Schedules,
      meta: { requiresAuth: true }
    },
    {
      path: '/students',
      name: 'Students',
      component: Students,
      meta: { requiresAuth: true }
    },
    {
      path: '/sign',
      name: 'Sign',
      component: Sign
    }
  ]
})
