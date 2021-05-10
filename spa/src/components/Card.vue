<template>
  <div class="cardWraper">
    <div class="card">
      <h2 class="name">{{ name }}</h2>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    name: String,
  },
});
</script>

<style scoped lang="scss">
.card {
  // https://stackoverflow.com/a/50109649/
  // TODO: Enable radius, margin and shadow
  //   margin-left: 5px;
  //   margin-right: 5px;
  //   border-radius: 10px;
  //   box-shadow: 0 14px 28px;

  width: 2.5in;
  min-width: 2.5in;
  height: 3.5in;

  background: coral;
  border: 5px solid darkred;

  position: relative;
  top: 0;
  transition: top ease 0.5s;
}

.cardWraper:not(hover) {
  overflow-x: hidden;
}

.card:hover {
  background: darken(coral, 10%);
  //   top: -30px;
  overflow-x: visible;
}
</style>
