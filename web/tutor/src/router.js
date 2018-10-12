import Vue from 'vue'
import Router from 'vue-router'
import axios from "axios";
import store from './store';

import Schedules from '@/views/Schedules/Schedules';
import Students from '@/views/Students/Students';
import ProgressClasses from '@/views/Progress/Classes/Progress';
import ProgressStudents from '@/views/Progress/Students/Progress';
import Search from '@/views/Search/Search';
import Sign from '@/views/Sign';

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: '/',
      component: Schedules,
      meta: { requiresAuth: true }
    },
    {
      path: '/students',
      component: Students,
      meta: { requiresAuth: true }
    },
    {
      path: '/progress',
      component: ProgressClasses,
      children: [
        {
          path: "classes",
          // component: () => import('@/views/Progress/Classes/Progress'),
        },
        {
          path: "students",
          component: ProgressStudents,
        },
      ],
      meta: { requiresAuth: true }
    },
    {
      path: '/search',
      component: Search,
      children: [
        {
          path: "",
          component: Search,
        },
        {
          path: "students",
          component: Search,
        },
      ],
      meta: { requiresAuth: true, isAside: true }
    },
    {
      path: '/sign',
      component: Sign,
    }
  ]
});

router.beforeEach((to, from, next) => {
  if (localStorage.getItem('token')) {
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  } else {
    axios.defaults.headers.common['Authorization'] = '';
  }
  // axios.defaults.headers.post['Content-Type'] = 'application/json; charset=utf-8';
  axios
    .get(`${process.env.VUE_APP_API}/auth`)
    .then(response => {
      // console.log(response.data);
      const currentAuth = response.data;
      store.commit('nextAuth', currentAuth);
    })
    .catch(error => {
      // console.log(error.response, to.path);
      if (to.matched.some(record => record.meta.requiresAuth)) {
        //   next({path: "/sign", query: {redirect: to.path}});
        next({ path: "/sign" });
      }
      // next({path: "/sign"});
    });
  next();
});

export default router
