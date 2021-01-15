<template>
  <v-dialog
    max-width="700"
    v-model="open"
    >
    <template v-slot:activator="{ on }">
      <v-btn rounded outlined elevation="1" v-on="on">+ Account</v-btn>
    </template>

    <v-card>
      <v-card-title>
        <span class="headline">+ Account</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">

              <v-text-field
                label="Name"
                v-model="alias"
                ></v-text-field>
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
            @click="newAccount"
            >
            Add
          </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  components: {
  },

  props: {
  },


  data () {
    return {
      open: false,

      words: [],
      alias: "",
    }
  },

  methods: {
    newAccount : async function() {
      this.open = false
      await Vue.api.createAccount({alias: this.alias})
      this.$store.dispatch("loadAccounts")
    },
  }
})
</script>

<style scoped>
</style>
