<template>
  <div class="hand">
    <Card v-for="card in cards" :name="card" :key="card" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Card from "./components/Card.vue";

export default defineComponent({
  name: "App",
  components: {
    Card,
  },
  data() {
    return {
      cards: ["King", "Queen", "Ace", "jack"],
    };
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.hand {
  display: flex;
  flex-direction: row;
}
</style>
