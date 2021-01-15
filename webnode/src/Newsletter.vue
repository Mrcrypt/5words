<template>
  <v-main>
    <v-container>
      <div class="messages">
        <v-btn rounded outlined elevation="1" class="writebtn">+ Publish</v-btn>
        <h2 class="mb-1 ml-4">New</h2>
        <v-list>
          <div v-for="n in 5" :key="n">
            <v-list-item :key="n" link="">
              <v-list-item-content>
                <v-list-item-title>Test Message Title #{{n}}</v-list-item-title>
                <v-list-item-subtitle>
                  <div><b>Max Mustermann</b> - Hey Bob I just wanted to write you a message because ...</div>
                </v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-action>
                <v-list-item-action-text >Tue 13:20</v-list-item-action-text>
              </v-list-item-action>
            </v-list-item>
          </div>
        </v-list>
        <h2 class="mb-1 ml-4">Seen</h2>
        <v-list>
          <div v-for="n in 5" :key="n">
            <v-list-item :key="n" link="">
              <v-list-item-content>
                <v-list-item-title>Test Message Title #{{n}}</v-list-item-title>
                <v-list-item-subtitle>
                  <div><b>Max Mustermann</b> - Hey Bob I just wanted to write you a message because ...</div>
                </v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-action>
                <v-list-item-action-text >Tue 13:20</v-list-item-action-text>
              </v-list-item-action>
            </v-list-item>
          </div>
        </v-list>
      </div>
    </v-container>
  </v-main>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Newsletter',

  components: {
  },

  data: () => ({
    //
  }),
});
</script>

<style scoped>
.writebtn {
  position: absolute !important;
  margin-top: -60px;
}

</style>
