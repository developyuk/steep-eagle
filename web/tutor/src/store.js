import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)
const defaultAuth = {
  id: 0,
  name: "",
  photo: "",
  role: "",
};
export default new Vuex.Store({
  state: {
    currentMqtt: null,
    currentAuth: defaultAuth,
    currentStats: {
      classes: 0,
      hours: 0,
      feedbacks: 0,
      ratings: 0,
      reviews: 0,
      attendances: 0,
    },
  },
  mutations: {
    nextAuth(state, nextVal) {
      state.currentAuth = Object.assign({}, defaultAuth, nextVal)
    },
    nextMqtt(state, nextVal) {
      state.currentMqtt = nextVal
    },

    nextStats(state, nextVal) {
      state.currentStats = Object.assign(state.currentStats, nextVal)
    },

  },
  actions: {
    updateStats({commit, state}) {
      console.log('state.currentAuth', state.currentAuth.id);
      axios.get(`${process.env.VUE_APP_DBAPI}/tutor_stats`)
        .then(response => {
          const data = {
            classes: response.data._items.classes_sum,
            hours: response.data._items.hours_sum,
            feedbacks: response.data._items.feedbacks_sum,
            ratings: response.data._items.ratings_avg,
            reviews: response.data._items.reviews_avg,
            attendances: response.data._items.attendances_avg,
          };
          commit('nextStats', data)
        })
        .catch(error => console.log(error))

    }

  }
})
