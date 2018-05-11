import Vue from 'vue'
import Vuex from 'vuex'
// import axios from 'axios'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    currentMqtt: null,
    currentAuth: {
      id: 0,
      name: "...",
      photo: "",
      role: "...",
    },

    currentDialog: null,
    currentSearch: null,

    currentStudentSession: null,
  },
  mutations: {
    nextAuth(state, nextVal) {
      let authPhoto = !!nextVal.photo ? nextVal.photo : "data:image/gif;base64,R0lGODdhAQABAPAAAMPDwwAAACwAAAAAAQABAAACAkQBADs=";
      if (authPhoto.indexOf('data:image/gif') < 0) {
        authPhoto = authPhoto.replace('https://', '').replace('http://', '');
        authPhoto = `//images.weserv.nl/?il&q=100&w=64&h=64&t=square&shape=circle&url=${authPhoto}`;
      }
      nextVal.photo = authPhoto;
      state.currentAuth = nextVal
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
  },
  actions: {},
})
