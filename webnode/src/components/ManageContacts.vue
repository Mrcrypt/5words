<template>
  <v-dialog max-width="1000" v-model="open">
    <template v-slot:activator="{ on }">
      <v-btn
        depressed
        rounded
        text
        v-on="on"
        >Manage Contacts</v-btn
      >
    </template>
        <v-card>
          <v-card-title>
            <span class="headline">Manage Contacts</span>
          </v-card-title>
          <v-card-text>
            <v-list>
              <v-subheader>Contacts</v-subheader>
              <v-list-item-group
                color="primary"
                >
                <v-list-item
                  v-for="(contact, i) in contacts"
                  :key="i"
                  >
                  <v-dialog max-width="400" v-model="editopen[contact.words]">
                    <template v-slot:activator="{ on }">
                      <v-list-item-content class="d-flex flex-row">
                        <v-list-item-title>{{contact.alias}} -
                          <span class="caption grey--text" v-text="contact.words"></span>
                        </v-list-item-title>
                      </v-list-item-content>
                      <v-list-item-action>
                        <v-btn icon v-on="on">
                          <v-icon color="grey lighten-1">mdi-pencil</v-icon>
                        </v-btn>
                      </v-list-item-action>
                    </template>
                    <v-card>
                      <v-card-title>
                        <span class="headline">Edit Contact</span>
                      </v-card-title>
                      <v-card-text>
                        <v-text-field label="Name" v-model="contact.alias"></v-text-field>
                      </v-card-text>
                      <v-card-actions>
                        <v-dialog
                          v-model="sureDialog[contact.words]"
                          max-width="290"
                          >
                          <template v-slot:activator="{ on, attrs }">
                            <v-btn color="red darken-1" v-bind="attrs" v-on="on" text > Delete </v-btn>
                          </template>
                          <v-card>
                            <v-card-title class="headline">
                              Are you Sure?
                            </v-card-title>
                            <v-card-actions>
                              <v-spacer></v-spacer>
                              <v-btn
                                color="green darken-1"
                                text
                                @click="sureDialog[contact.words] = false"
                                >
                                No
                              </v-btn>
                                <v-btn
                                  color="green darken-1"
                                  text
                                  @click="sureDialog[contact.words] = false; deleteContact(contact)"
                                  >
                                  Yes
                                </v-btn>
                            </v-card-actions>
                          </v-card>
                        </v-dialog>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="editopen[contact.words] = false"> Close </v-btn>
                        <v-btn color="blue darken-1" text @click="editopen[contact.words] = false; updateContact(contact)"> Save </v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                </v-list-item>
              </v-list-item-group>
            </v-list>
            <new-contact></new-contact>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="blue darken-1" text @click="open = false"> Close </v-btn>
          </v-card-actions>
        </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import { mapState } from 'vuex';

import NewContact from './NewContact.vue';

export default Vue.extend({
  components: {
    NewContact,
  },

  props: {},

  computed: mapState(["contacts"]),

  data() {
    return {
      open: false,
      sureDialog: [],

      editopen: [],
    };
  },

  created: async function() {
  },

  methods: {
    updateContact: async function(contact : any) {
      await Vue.api.updateContact(contact);
      this.$store.dispatch("loadContacts")
    },
    deleteContact: async function(contact : any) {
      await Vue.api.deleteContact(contact);
      this.$store.dispatch("loadContacts")
    }
  },
});
</script>

<style scoped></style>
