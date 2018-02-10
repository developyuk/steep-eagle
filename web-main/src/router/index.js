import Vue from 'vue'
import Router from 'vue-router'
import Sign from '@/pages/Sign'
import Schedules from '@/pages/Schedules'
import Profile from '@/pages/Profile'
// import Hello from '@/components/Hello'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Schedules',
      component: Schedules
    },
    {
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/sign',
      name: 'Sign',
      component: Sign
    }
  ]
})
