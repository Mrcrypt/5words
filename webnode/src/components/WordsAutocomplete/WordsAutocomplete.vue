<template>
  <div>
    <div class="words d-flex align-center">
      <WordsAutocompleteWord
        v-for="(word, index) in words"
        v-model="words[index]"
        :key="index"
        :ref="`word${index}`"
        :index="index"
        @done="wordCompleted(index)"
        @overflow="pushOverflowToNextWord"
      />
    </div>
  </div>
</template>
<script lang="ts">
import { Vue, Component, Prop, Ref, Watch } from "vue-property-decorator";
import WordsAutocompleteWord from "./WordsAutocompleteWord.vue";
@Component({
  components: {
    WordsAutocompleteWord,
  },
})
export default class WordsAutocomplete extends Vue {
  @Prop({ type: Array, default: [] }) value!: string[];
  @Ref("word1") word1: any;
  @Ref("word2") word2: any;
  @Ref("word3") word3: any;
  @Ref("word4") word4: any;
  @Ref("word5") word5: any;
  words: string[] = [];

  @Watch("words")
  wordsWatcher(newVal: any) {
    this.$emit("input", newVal)
  }

  created() {
    this.words = this.value;
    while (this.words.length < 5) this.words = this.words.concat([""]);
  }

  wordCompleted(index: number) {
    console.log("word" + (index + 1));
    let nextword = this.$refs["word" + (index + 1)] as any;
    if (nextword) nextword[0].focus();
  }

  pushOverflowToNextWord({ index , text } : {index: any, text: any}) {
    (this as any).$refs["word" + (index + 1)][0].insert(text);
  }
}
</script>

<style scoped>
.words {
}
</style>
