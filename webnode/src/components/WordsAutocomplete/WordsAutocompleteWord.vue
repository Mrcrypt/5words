<template>
  <div class="word mx-1 px-2 d-flex align-center">
    <v-text-field
      ref="input"
      v-model="word"
      @input="inputCheck"
      @blur="update"
      @keydown="checkKeydown"
      v-show="!done"
      dense
      hide-details
      class="py-1"
      />
      <span v-show="done" @click="editField">{{ word }}</span>
      <v-card
        class="autocomplete-modal"
        v-if="filteredItems.length !== 0 && word.length !== 0 && !done"
        >
        <v-list dense>
          <v-list-item
            v-for="(item, index) in filteredItems"
            :class="{ active: index == suggestIndex }"
            @click="selectWord(item)"
            :key="item"
            >{{ item }}</v-list-item
          >
        </v-list>
      </v-card>
  </div>
</template>

<script lang="ts">
import { Vue, Component, Prop, Ref } from "vue-property-decorator";

@Component({})
export default class WordsAutocompleteWord extends Vue {
  @Prop({ type: String, default: "" }) value!: string;
  @Prop({ type: Number, default: 0 }) index!: number;
  @Ref("input") input!: HTMLInputElement;

  word = "";
  editing = true;
  suggestIndex = 0;

  get done() {
    return this.valid && !this.editing;
  }
  get valid() {
    let wordlist = this.$store.state.wordList;
    return wordlist.includes(this.word);
  }
  get filteredItems() {
    let wordlist = this.$store.state.wordList;
    return wordlist
      .filter((b : any) => b.startsWith(this.word))
      .filter((b : any, i : any) => i < 3);
  }

  created() {
    this.word = this.value;
  }
  update() {
    this.$emit("input", this.word);
  }
  focus() {
    this.input.focus();
  }
  insert(text: string) {
    this.word = text;
    this.validate();
  }
  inputCheck(input: string) {
    //check for autocomplete stuff here
    console.log(input);
  }
  editField() {
    this.editing = true;
    this.$nextTick(() => {
      this.focus();
    });
  }
  selectWord(word: string) {
    this.word = word;
    this.$nextTick(() => {
      this.validate();
    });
  }
  suggestDown() {
    this.suggestIndex++;
  }
  suggestUp() {
    this.suggestIndex--;
  }
  validate() {
    this.word = this.word.toLowerCase();
    this.word = this.word.trim();
    // validate the word here
    let wordlist = this.$store.state.wordList;
    // check if pasted multiple words
    let multiple = this.word.split(" ");
    if (multiple.length > 1) {
      // got pasted
      this.word = multiple[0];
      multiple.shift();
      let multiplestr = multiple.join(" ");
      if (multiplestr !== " ")
        this.$emit("overflow", { text: multiplestr, index: this.index });
    }
    console.log(wordlist, this.word);
    if (wordlist.includes(this.word)) {
      this.$emit("input", this.word);
      this.$emit("done");
      this.editing = false;
    }
  }

  checkKeydown(e: any) {
    console.log(e.which);
    switch (e.which) {
      case 13:
      case 9:
        e.preventDefault();
          this.word = this.filteredItems[this.suggestIndex];
          this.$nextTick(() => {
            this.validate();
          });
        break;
      case 32:
        this.validate();
        break;
      case 38:
        this.suggestUp();
        break;
      case 40:
        this.suggestDown();
        break;
    }
  }
}
</script>

<style lang="scss" scoped>
@import "../../styles/variables.scss";
.word {
  background: $colors-1-rgba;
  border-radius: 5px;
  position: relative;
}
.autocomplete-modal {
  position: absolute;
  top: 100%;
  width: 140px;
  max-width: inherit;
  z-index: 5;
  overflow: visible;
  .active {
    background: $colors-1-rgba;
  }
}
</style>
