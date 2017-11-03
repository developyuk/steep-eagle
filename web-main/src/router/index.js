import Vue from 'vue'
import Router from 'vue-router'
// import Hello from '@/components/Hello'
import Schedules from '@/pages/Schedules'
import ClassDetail from '@/pages/ClassDetail'

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
      name: 'ClassDetail',
      component: ClassDetail
    }
  ]
})
