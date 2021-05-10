<template>
  <div class="center">
    <div class="card left" v-if="left">{{ left }}</div>
    <div class="card right" v-if="right">{{ right }}</div>
    <div class="card up" v-if="up">{{ up }}</div>
    <div class="card down" v-if="down">{{ down }}</div>
  </div>
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
// TODO: Dry with other card class
.card {
  background: coral;
  border: 5px solid darkred;

  width: 150px;

  min-width: 150px;
  height: 200px;
}

// TODO: Make this look nicer
.center {
  display: grid;
  grid-template-columns: repeat(3, 150px);
  grid-template-rows: 200px 0px 200px;
  gap: 15px;
  grid-template-areas:
    ". Top ."
    "Left . Right"
    ". Bottom .";
}

.up {
  grid-area: Top;
}
.left {
  grid-area: Left;
}
.down {
  grid-area: Bottom;
}
.right {
  grid-area: Right;
}
</style>
