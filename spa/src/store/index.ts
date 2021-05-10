import { createStore } from "vuex";
// import { Card } from "@/card";

export default createStore({
  state: {
    hand: ["King Of Clubs", "Jack of Diamonds"],
    center: {
      left: "Queen",
      right: "spade",
      up: "Joker",
      down: "evlen",
    },
  },
  mutations: {},
  actions: {},
  modules: {},
});
