import Vue from 'vue'
import Router from 'vue-router'
import axios from "axios";
import store from './store';

Vue.use(Router);

const router = new Router({
  routes: [
    {
      path: '/',
      component: () => import('@/views/Schedules/Schedules'),
      meta: { requiresAuth: true }
    },
    {
      path: '/students',
      component: () => import('@/views/Students/Students'),
      meta: { requiresAuth: true }
    },
    {
      path: '/progress',
      component: () => import('@/views/Progress/Classes/Progress'),
      children: [
        {
          path: "classes",
          // component: () => import('@/views/Progress/Classes/Progress'),
        },
        {
          path: "students",
          component: () => import('@/views/Progress/Students/Progress'),
        },
      ],
      meta: { requiresAuth: true }
    },
    {
      path: '/search',
      component: () => import('@/views/Search/Search'),
      children: [
        {
          path: "",
          component: () => import('@/views/Search/Search'),
        },
        {
          path: "students",
          component: () => import('@/views/Search/Search'),
        },
      ],
      meta: { requiresAuth: true, isAside: true }
    },
    {
      path: '/sign',
      component: () => import('@/views/Sign'),
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
