<template>
  <v-app>
    <div class="menubar d-flex justify-center align-center">
      <mainmenu></mainmenu>
      <v-tabs class="d-flex justify-center align-center">
        <v-tab to="/">Postbox</v-tab>
        <!--<v-tab to="/newsletter">Newsletter</v-tab>-->
      </v-tabs>
    </div>

    <router-view></router-view>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import mainmenu from './components/Menu.vue';

export default Vue.extend({
  name: 'App',

  components: {
    mainmenu,
  },

  data: () => ({
    //
  }),

  created: function() {
    this.$store.dispatch("loadWordList")
    this.$store.dispatch("loadAccounts")
    this.$store.dispatch("loadContacts")
  },
});
</script>

<style>
html { zoom: 1.5; }

.menubar {
  height: 60px;
}

</style>
