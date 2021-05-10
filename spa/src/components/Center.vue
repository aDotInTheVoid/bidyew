<template>
  <div class="card" v-if="left">{{ left }}</div>
  <div class="card" v-if="right">{{ right }}</div>
  <div class="card" v-if="up">{{ up }}</div>
  <div class="card" v-if="down">{{ down }}</div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    left: String,
    right: String,
    up: String,
    down: String,
  },
});
</script>

<style scoped lang="scss">
.card {
  background: coral;
}
</style>