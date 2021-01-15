<template>
  <v-menu
    bottom
    min-width="200px"
    rounded
    offset-y
    :close-on-content-click="false"
    >
    <template v-slot:activator="{ on }">
      <v-btn
        icon
        x-large
        v-on="on"
        class="ma-1 ml-5"
        >
        <v-avatar
          color="teal"
          size="48"
          >
          <span class="white--text headline"></span>
        </v-avatar>
      </v-btn>
    </template>
    <v-card>
      <v-list-item-content class="justify-center">
        <div class="mx-auto text-center">
          <div v-for="account in accounts" :key="account.address">
            <h3>{{account.alias}}</h3>
            <p class="caption mt-1 my-1 px-2">
            <copyable>
            {{account.address}}
            </copyable>
            </p>
            <v-divider class="my-2"></v-divider>
          </div>
          <manage-accounts>
          </manage-accounts>
          <v-divider class="my-2"></v-divider>
          <manage-contacts>
          </manage-contacts>
          <v-divider class="my-2"></v-divider>
          <v-btn
            depressed
            rounded
            text
            >
            Quit
          </v-btn>
        </div>
      </v-list-item-content>
    </v-card>
  </v-menu>
</template>

<script lang="ts">
import Vue from 'vue'
import { mapState } from 'vuex';

import Copyable from './Copyable.vue';
import ManageAccounts from './ManageAccounts.vue';
import ManageContacts from './ManageContacts.vue';

export default Vue.extend({
  components: {
    Copyable,
    ManageAccounts,
    ManageContacts,
  },

  props: {
  },

  computed: mapState(["accounts"]),

  data () {
    return {
    }
  },

  methods: {
  }
})
</script>

<style scoped>
</style>
