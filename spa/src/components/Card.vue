<template>
  <div class="card">
    <h2>{{ name }}</h2>
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
  width: 2.5in;
  height: 3.5in;
  margin-left: 30px;
  margin-right: 30px;
  background: coral;
}
</style>
