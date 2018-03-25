import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Schedules',
      component: () => import('@/pages/Schedules'),
      meta: {requiresAuth: true}
    },
    {
      path: '/students',
      name: 'Students',
      component: () => import('@/pages/Students/Students'),
      meta: {requiresAuth: true}
    },
    {
      path: '/sign',
      name: 'Sign',
      component: () => import('@/pages/Sign'),
    }
  ]
})
