import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'
import Schedules from '@/pages/Schedules'
import ClassDetail from '@/pages/ClassDetail'
import ComingSoon from '@/components/ComingSoon'
import Sign from '@/pages/Sign'
import Profile from '@/pages/Profile'
import ProfileSettings from '@/pages/ProfileSettings'

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
      component: Schedules
    },
    {
      path: '/class/:id',
      component: ComingSoon
    },
    {
      path: '/activities',
      component: ComingSoon
    },
    {
      path: '/progress',
      component: ComingSoon
    },
    {
      path: '/profile',
      component: Profile
    },
    {
      path: '/profile/settings',
      component: ProfileSettings,
    },
    {
      path: '/notifications',
      component: ComingSoon,
    },
    {
      path: '/how-to-use',
      component: ComingSoon,
    },
    {
      path: '/sign',
      component: Sign
    }
  ]
})
