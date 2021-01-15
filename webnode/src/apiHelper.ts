import Vue from "vue";
import axios from "axios";

let APIPort = (window as any).APIPort;
if (APIPort === undefined) {
  APIPort = 9000;
}
axios.defaults.baseURL = "http://localhost:" + APIPort + "/";
Vue.prototype.$http = axios;

declare module "vue/types/vue" {
  interface VueConstructor {
    api: API;
  }
}

//Types
export interface Message {
  id: number;
  topic: string;
  name: string;
  content: string;
}


//API
export class API {
  async getWordList() {
    const resp = await axios.get("wordlist");
    return resp.data.wordList as string[];
  }

  async getMessages() {
    const resp = await axios.get("messages");
    return resp.data.messages as Message[];
  }

  //Contacts
  async getContacts() {
    const resp = await axios.get("contacts");
    return resp.data.contacts as any;
  }

  async deleteContact(contact : any) {
    await axios.delete("/contacts", {data: contact});
  }

  async updateContact(contact : any) {
    await axios.put("/contacts", contact);
  }

  async sendMessage(message: any) {
    await axios.post("/sendmessage", message);
  }

  async newContact(words: any, alias: string) {
    await axios.post("/newcontact", { Words: words.join(" "), Alias: alias });
  }

  async newFollowing(words: any, alias: string) {
    await axios.post("/newfollowing", { Words: words.join(" "), Alias: alias });
  }

  //Accounts
  async deleteAccount(account : any) {
    await axios.delete("/accounts", {data: account});
  }

  async updateAccount(account : any) {
    await axios.put("/accounts", account);
  }

  async createAccount(account : any) {
    await axios.post("/accounts", account);
  }

  async getNewAccounts() {
    return (await axios.get("/accounts/new")).data.accounts as any;
  }

  async getAccounts() {
    return (await axios.get("/accounts")).data.accounts as any;
  }
}

export default {
  install(Vue: any, options: any) {
    Vue.api = new API();
  },
};
