import { createStore } from "vuex";
import { Card } from "@/card";

export interface State {
  hand: string[],
  canPlay: boolean,
  center: {
    up: string | null,
    down: string | null,
    left: string | null,
    right: string | null,
  },
}

function cardName(card: any): string {
  return `${card.Value} of ${card.Suit}`
}

export default createStore<State>({
  state: {
    hand: [],
    canPlay: false,
    center: {
      left: null,
      right: null,
      up: null,
      down: null,
    },
  },
  mutations: {
    play(state, card: string) {
      let index = state.hand.indexOf(card);
      state.hand.splice(index, 1);
      state.center.down = card;
    },
    setDeck(state, newDeck) {
      state.hand = newDeck.map(cardName);
    },
    clearCenter(state) {
      state.center = {
        left: null,
        right: null,
        up: null,
        down: null,
      }
    },
    setLeft(state, card) {
      state.center.left = cardName(card);
    },
    setRight(state, card) {
      state.center.right = cardName(card);
    },
    setUp(state, card) {
      state.center.up = cardName(card);
    },
    setDown(state, card) {
      state.center.down = cardName(card);
    }
  },
  actions: {},
  modules: {},
});
