import Vue from 'vue';
import routes from './routes';
import axios from 'axios';
import VueRouter from 'vue-router';
import store from '../store/index'

Vue.use(VueRouter);

// configure router
const router = new VueRouter({
  routes, // short for routes: routes
  linkActiveClass: 'active'
})

router.beforeEach((to, from, next) => {
  if (localStorage.getItem('token')) {
    axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token');
  } else {
    axios.defaults.headers.common['Authorization'] = '';
  }
  // axios.defaults.headers.post['Content-Type'] = 'application/json; charset=utf-8';
  axios
    .get(`${process.env.API}/auth`)
    .then(response => {
      // console.log(response.data);
      const currentAuth = response.data;
      store.commit('nextAuth', currentAuth);
    })
    .catch(error => {
      // console.log(error.response, to.path);
      // if (to.matched.some(record => record.meta.requiresAuth)) {
      //   //   next({path: "/sign", query: {redirect: to.path}});
      //   next({path: "/sign"});
      // }
      // next({path: "/sign"});
    });
  next();
});

export default router;
