<template>
  <div class="card" :class="{ clickable: isClickable }">
    <h2 class="name">{{ name }}</h2>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    name: String,
    isClickable: Boolean,
  },
  data() {
    return;
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

  width: 150px;

  min-width: 150px;
  height: 200px;

  background: lighten(coral, 10%);
  border: 5px solid darkred;

  position: relative;
  top: 0;
  transition: ease 0.5s;
}

.card.clickable {
  background: coral;
}

.clickable.card:hover {
  background: darken(coral, 10%);
  top: -30px;
  overflow-x: visible;
}
</style>
