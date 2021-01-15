<template>
  <v-dialog max-width="1000" v-model="open">
    <template v-slot:activator="{ on }">
      <v-btn
        depressed
        rounded
        text
        v-on="on"
        >Edit Accounts</v-btn
      >
    </template>
        <v-card>
          <v-card-title>
            <span class="headline">Edit Accounts</span>
          </v-card-title>
          <v-card-text>
            <v-list>
              <v-subheader>Accounts</v-subheader>
              <v-list-item-group
                color="primary"
                >
                <v-list-item
                  v-for="(account, i) in accounts"
                  :key="i"
                  >
                  <v-dialog max-width="400" v-model="editopen[account.address]">
                    <template v-slot:activator="{ on }">
                      <v-list-item-content class="d-flex flex-row">
                        <v-list-item-title>{{account.alias}} -
                          <span class="caption grey--text" v-text="account.address"></span>
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
                        <span class="headline">Edit Account</span>
                      </v-card-title>
                      <v-card-text>
                        <v-text-field label="Name" v-model="account.alias"></v-text-field>
                      </v-card-text>
                      <v-card-actions>
                        <v-dialog
                          v-model="sureDialog[account.address]"
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
                                @click="sureDialog[account.address] = false"
                                >
                                No
                              </v-btn>
                                <v-btn
                                  color="green darken-1"
                                  text
                                  @click="sureDialog[account.address] = false; deleteAccount(account)"
                                  >
                                  Yes
                                </v-btn>
                            </v-card-actions>
                          </v-card>
                        </v-dialog>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="editopen[account.address] = false"> Close </v-btn>
                        <v-btn color="blue darken-1" text @click="editopen[account.address] = false; updateAccount(account)"> Save </v-btn>
                      </v-card-actions>
                    </v-card>
                  </v-dialog>
                </v-list-item>
              </v-list-item-group>
            </v-list>
            <new-account></new-account>
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

import NewAccount from './NewAccount.vue';

export default Vue.extend({
  components: {
    NewAccount,
  },

  props: {},

  computed: mapState(["accounts"]),

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
    updateAccount: async function(account : any) {
      await Vue.api.updateAccount(account);
      this.$store.dispatch("loadAccounts")
    },
    deleteAccount: async function(account : any) {
      await Vue.api.deleteAccount(account);
      this.$store.dispatch("loadAccounts")
    }
  },
});
</script>

<style scoped></style>
