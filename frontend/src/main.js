import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

import App from "./App.vue";

Vue.config.productionTip = false;

Vue.use(Vuex);

let store = new Vuex.Store({
  state: {
      companies: []
  },

  mutations: {
    setCompanies(state, payload) {
    }
  },

  actions: {
    refreshCompanies(context) {
      return new Promise(resolve => {
        axios.get("http://localhost:8000/companies").then(response => {
          resolve();
        });
      });
    }
  }
});

new Vue({
  render: h => h(App),
  store
}).$mount("#app");
