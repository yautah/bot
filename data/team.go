package data

type Team struct {
	Tag      string `json:tag`
	Name     string `json:name`
	DeckLink string `json:deckLink`
	Deck     []Card `json:deck`
}
