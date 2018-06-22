import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)
const defaultAuth = {
  id: 0,
  name: "...",
  photo: "",
  role: "...",
};
export default new Vuex.Store({
  state: {
    currentMqtt: null,
    currentAuth: defaultAuth,

    currentDialog: null,
    currentSearch: null,

    currentStudentSession: null,

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

    nextDialog(state, nextVal) {
      state.currentDialog = nextVal
    },
    nextSearch(state, nextVal) {
      state.currentSearch = nextVal
    },

    nextStudentSession(state, nextVal) {
      state.currentStudentSession = nextVal
    },
    nextStats(state, nextVal) {
      state.currentStats = Object.assign(state.currentStats, nextVal)
    },

  },
  actions: {
    updateStats({commit, state}) {
      console.log('state.currentAuth', state.currentAuth.id);
      commit('nextStats', {})
    }

  }
})
