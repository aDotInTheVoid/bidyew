<template>
  <div class="table">
    <Center right="9" up="foo" down="bar" />
    <div class="hand">
      <Card
        v-for="card in cards"
        :name="card"
        :key="card"
        :isClickable="isClickable"
        @click="remove(card)"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Card from "@/components/Card.vue";
import Center from "@/components/Center.vue";

export default defineComponent({
  name: "App",
  components: {
    Card,
    Center,
  },
  data() {
    return {
      cards: Array(8).fill("Some Card"),
      isClickable: true,
    };
  },
  methods: {
    remove(name: string) {
      if (this.isClickable) {
        let index = this.cards.indexOf(name);
        this.cards.splice(index, 1);
        this.isClickable = false;

        // Kind of hacky solution, be proper when theirs a server
        setInterval(() => {
          this.isClickable = true;
        }, 1000);
      }
    },
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  // text-align: center;
  color: #2c3e50;
}
</style>

<style scoped lang="scss">
.table {
  position: relative;
  height: 100%;
}

.hand {
  position: fixed;
  display: flex;
  align-content: center;
  max-width: 100vw;
  bottom: 0;
}
</style>