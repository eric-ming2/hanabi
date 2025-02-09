package state

type GameState struct {
	Started         bool
	NotStartedState NotStartedGameState
	StartedState    StartedGameState
}

type NotStartedGameState struct {
	Players []NotStartedPlayer
}

type NotStartedPlayer struct {
	Name  string
	Id    string
	Ready bool
}

type StartedGameState struct {
	MyHand       []UnknownCard
	OtherPlayers []StartedPlayer
	Turn         uint8
	Deck         []Card
	DiscardPile  []Card
	Hints        uint8
	Bombs        uint8
	Fireworks    map[CardColor]uint8
}

type CardColor int

const (
	White  CardColor = iota // 0
	Yellow                  // 1
	Green                   // 2
	Blue                    // 3
	Red                     // 4
)

type Card struct {
	Color CardColor
	Num   uint8
}

type UnknownCard struct {
	ColorKnown bool
	Color      CardColor
	NumKnown   bool
	Num        uint8
}

type StartedPlayer struct {
	Name         string
	Id           string
	Cards        []Card
	UnknownCards []UnknownCard
}
