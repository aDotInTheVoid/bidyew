<template>
  <div class="table">
    <Center
      :right="center.right"
      :left="center.left"
      :up="center.up"
      :down="center.down"
    />
    <div class="hand">
      <Card
        v-for="card in hand"
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
import { mapState } from "vuex";

import Card from "@/components/Card.vue";
import Center from "@/components/Center.vue";
import { handleEvent } from "@/conn";

export default defineComponent({
  name: "App",
  components: {
    Card,
    Center,
  },
  data() {
    return {
      // cards: Array(8).fill("Some Card"),
      isClickable: true,
      conn: null,
    };
  },
  computed: mapState(["hand", "center"]),
  created() {
    const conn = new WebSocket("ws://localhost:9843/ws");
    conn.onmessage = (event) => handleEvent(event, this.$store);
  },
  methods: {
    remove(name: string) {
      this.$store.commit("play", name);
      // TODO: handle isClickable
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
