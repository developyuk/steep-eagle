import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      component: () => import('@/pages/Schedules/Schedules'),
      meta: {requiresAuth: true}
    },
    {
      path: '/students',
      component: () => import('@/pages/Students/Students'),
      meta: {requiresAuth: true}
    },
    {
      path: '/progress',
      component: () => import('@/pages/Progress/Progress'),
      meta: {requiresAuth: true}
    },
    {
      path: '/sign',
      name: 'Sign',
      component: () => import('@/pages/Sign'),
    }
  ]
})
