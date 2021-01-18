<template>
  <v-main>
    <v-container>
      <div class="messages">

        <new-message></new-message>

        <h2 class="mb-1 ml-4">Messages</h2>
        <v-list>
          <div v-for="msg in messages" :key="msg.id">
            <v-dialog
              max-width="1200"
              >
              <template v-slot:activator="{ on }">
                <v-list-item link="" v-on="on">
                  <v-list-item-content>
                    <v-list-item-title>{{msg.topic}}</v-list-item-title>
                    <v-list-item-subtitle>
                      <div><b>{{msg.name}}</b> - {{msg.content}}</div>
                    </v-list-item-subtitle>
                  </v-list-item-content>
                  <v-list-item-action>
                    <v-list-item-action-text >{{ msg.time }}</v-list-item-action-text>
                  </v-list-item-action>
                </v-list-item>
              </template>
              <v-card>
                <v-card-title class="headline grey lighten-2">
                  {{msg.topic}}
                </v-card-title>
                <v-card-text>
                  {{msg.content}}
                </v-card-text>
                <v-divider></v-divider>
              </v-card>
            </v-dialog>
          </div>
        </v-list>
      </div>
    </v-container>
  </v-main>
</template>

<script lang="ts">
import Vue from 'vue';
import * as api from './apiHelper';api;

import NewMessage from './components/NewMessage.vue';

export default Vue.extend({
  name: 'Postbox',

  components: {
    NewMessage,
  },

  data: () => ({
    messages: [] as api.Message[],

    reloadTimer: null as any,
  }),

  created: async function() {
    this.messages = await Vue.api.getMessages()
    this.reloadTimer = setInterval(async ()=>{
      this.messages = await Vue.api.getMessages()
    }, 10000)
  },

  beforeDestroy: function() {
    clearInterval(this.reloadTimer)
  },

});
</script>

<style scoped>

</style>
