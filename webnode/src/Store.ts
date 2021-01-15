import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    messages: [] as any,
    accounts: [] as any,
    contacts: [] as any,

    wordList: [] as string[],
  },

  mutations: {},

  actions: {
    async loadWordList() {
      this.state.wordList = await Vue.api.getWordList()
    },
    async loadAccounts() {
      this.state.accounts = await Vue.api.getAccounts()
    },
    async loadContacts() {
      this.state.contacts = await Vue.api.getContacts()
    },
  },
});
