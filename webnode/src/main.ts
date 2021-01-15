//Vue
import Vue from 'vue'
import vuetify from './plugins/vuetify';
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import '@mdi/font/css/materialdesignicons.css'

//Globals
import VueRouter from 'vue-router'
Vue.use(VueRouter)

//API
import api from './apiHelper'
Vue.use(api)

import store from './Store'

//Modules
import App from './App.vue'
import Postbox from './Postbox.vue'
import Newsletter from './Newsletter.vue'


Vue.config.productionTip = false

const router = new VueRouter({
  mode: 'history',
  routes: [
    { path: '/', component: Postbox },
    { path: '/newsletter', component: Newsletter },
  ]
})

new Vue({
  router,
  vuetify,
  store,
  render: h => h(App)
}).$mount('#app')
