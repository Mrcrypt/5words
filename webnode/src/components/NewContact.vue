<template>
  <v-dialog max-width="500" v-model="open">
    <template v-slot:activator="{ on }">
      <v-btn rounded small outlined elevation="1" class="" v-on="on"
        >+ New Contact</v-btn
      >
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">+ New Contact</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <words-autocomplete
                v-model="words"
              ></words-autocomplete>
              <v-text-field label="Name" v-model="addressAlias"></v-text-field>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="open = false"> Close </v-btn>
        <v-btn color="blue darken-1" text @click="newContact"> Add </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import Vue from "vue";
import WordsAutocomplete from "@/components/WordsAutocomplete";

export default Vue.extend({
  components: {
    WordsAutocomplete,
  },

  props: {},

  data() {
    return {
      open: false,

      words: [] as string[],
      addressAlias: "",
    };
  },

  methods: {
    newContact: async function () {
      this.open = false;
      await Vue.api.newContact(this.words, this.addressAlias);
      this.$store.dispatch("loadContacts")
    },
  },
});
</script>

<style scoped></style>
