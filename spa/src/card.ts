export interface Card {
  suit: Suit;
  position: Position;
}

export type Suit = "Clubs" | "Diamonds" | "Spades " | "Hearts";
export type Position = "Ace" | "King" | "Queen" | "Jack" | 10 | 9;
