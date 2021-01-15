<template>
  <v-dialog
    max-width="1200"
    v-model="open"
    >
    <template v-slot:activator="{ on }">
      <v-btn rounded outlined elevation="1" class="writebtn" v-on="on">+ Write</v-btn>
    </template>

    <v-card>
      <v-card-title>
        <span class="headline">Write</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12" sm="6" md="4" >
              <new-contact></new-contact>
              <br>
              <br>
              <v-autocomplete
                label="Reciever"
                hide-selected
                multiple
                dense
                chips
                deletable-chips
                :items="contacts"
                :search-input.sync="recieverSearch"
                @change="recieverSearch = ''"
                v-model="newmessage.reciever"
                ></v-autocomplete>
            </v-col>
            <v-col cols="12">
              <v-text-field
                label="Title"
                v-model="newmessage.title"
                ></v-text-field>
              <v-textarea
                label="Text"
                v-model="newmessage.content"
                ></v-textarea>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn
          color="blue darken-1"
          text
          @click="open = false"
          >
          Close
        </v-btn>
          <v-btn
            color="blue darken-1"
            text
            @click="sendMessage"
            >
            Send
          </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from 'vue'

import NewContact from './NewContact.vue'

export default Vue.extend({
  components: {
    NewContact
  },

  props: {
  },

  data () {
    return {
      open: false,

      contacts: [] as any,
      recieverSearch: "",

      newmessage: {
        reciever: [] as string[],
        content: "",
        title: "",
      },

      reloadTimer: null as any,
    }
  },

  created: function() {
    this.reloadTimer = setInterval(async ()=>{
      this.contacts = await Vue.api.getContacts()
      this.contacts = this.contacts.map((old :any) => old.alias)
    }, 10000)
  },

  beforeDestroy: function() {
    clearInterval(this.reloadTimer)
  },

  methods : {
    sendMessage : async function() {
      this.open = false
      await Vue.api.sendMessage(this.newmessage)
    },
  },
})
</script>

<style scoped>
.writebtn {
  position: absolute !important;
  margin-top: -60px;
  margin-left: 70px;
}
</style>
