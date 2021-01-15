<template>
  <v-dialog
    max-width="700"
    v-model="open"
    >
    <template v-slot:activator="{ on }">
      <v-btn rounded outlined elevation="1" class="followbtn" v-on="on">+ Follow</v-btn>
    </template>

    <v-card>
      <v-card-title>
        <span class="headline">+ Follow</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <!--<v-autocomplete
                label="5 Words"
                hide-selected
                multiple
                dense
                chips
                deletable-chips
                auto-select-first
                :items="wordList"
                :search-input.sync="wordSearch"
                @change="wordSearch = ''"
                v-model="words"
                ></v-autocomplete>-->
              <words-autocomplete
                v-model="words"
              ></words-autocomplete>

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
            @click="newFollowing"
            >
            Add
          </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>

</template>

<script lang="ts">
import Vue from 'vue'
import { mapState } from 'vuex'
import WordsAutocomplete from "@/components/WordsAutocomplete";

export default Vue.extend({
  components: {
    WordsAutocomplete,
  },

  props: {
  },

  computed: mapState([
    'wordList'
  ]),

  data () {
    return {
      open: false,
      alias: "",
      words: [],

      wordSearch: "",
    }
  },

  methods: {
    newFollowing : async function() {
      this.open = false
      await Vue.api.newFollowing(this.words, this.alias)
    },
  }
})
</script>

<style scoped>
.followbtn {
  position: absolute !important;
  margin-top: -60px;
  margin-left: 190px;
}
</style>
