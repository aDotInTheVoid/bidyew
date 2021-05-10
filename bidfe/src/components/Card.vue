<template>
  <div class="card">
    <h2 v-if="!fliped">{{ name }}</h2>
    <h2 v-else>Some Card</h2>
    <button @click="fliped = !fliped">Flip</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "Card",
  props: {
    name: {
      type: String,
      required: true,
    },
  },
  data() {
    return {
      fliped: false,
    };
  },
});
</script>

<style scoped>
.card {
  width: 100px;
  height: 300px;
  margin: 0px 10px 0px 10px;
  background: coral;
}
</style>