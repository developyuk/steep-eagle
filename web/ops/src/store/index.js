import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);
const store = new Vuex.Store({
  state: {
    currentAuth: 0,
  },
  mutations: {
    nextAuth(state, nextVal) {
      state.currentAuth = Object.assign({}, state.currentAuth, nextVal);
    }
  }
});

export default store;
