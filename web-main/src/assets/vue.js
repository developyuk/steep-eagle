import axios from "axios";

export default {
  routerBeforeEach(router, to, from, next) {
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

        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 3);
        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 8);
        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 13);
        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 21);
        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 34);
        setTimeout(() => router.app.$bus.$emit('currentAuth', response.data), 99 * 55);
      })
      .catch(error => {
        // console.log(error.response, to.path);
        // if (to.matched.some(record => record.meta.requiresAuth)) {
        //   next({path: "/sign", query: {redirect: to.path}});
        // }
        next({path: "/sign"});
      });
    next();
  }

}
