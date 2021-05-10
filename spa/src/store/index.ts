import { createStore } from "vuex";
import { Card } from "@/card";

export interface State {
  hand: string[],
  center: {
    up: string,
    down: string,
    left: string,
    right: string,
  }
}

export default createStore<State>({
  state: {
    hand: ["King Of Clubs", "Jack of Diamonds"],
    center: {
      left: "Queen",
      right: "spade",
      up: "Joker",
      down: "evlen",
    },
  },
  mutations: {
    play(state, card: string) {
      let index = state.hand.indexOf(card);
      state.hand.splice(index, 1);
      state.center.down = card;
    },
    setDeck(state, newDeck) {
      state.hand = newDeck.map((x: any) => `${x.Value} of ${x.Suit}`)
    }
  },
  actions: {},
  modules: {},
});
