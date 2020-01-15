package bookmaker

import "github.com/statistico/statistico-bet-finder/internal/statistico"

type Service interface {
	Markets(query ServiceQuery) <-chan *Market
}

type ServiceQuery struct {
	BetTypes []string
	Fixtures []statistico.Fixture
}

//type Book struct {
//	FixtureID uint64   `json:"fixture_id"`
//	Bookmaker string   `json:"bookmaker"`
//	Markets   Market   `json:"markets"`
//}

type Market struct {
	ID        string   `json:"id"`
	FixtureID uint64   `json:"fixture_id"`
	Bookmaker string   `json:"bookmaker"`
	Name      string   `json:"name"`
	BetType   string   `json:"bet_type"`
	Runners   []Runner `json:"book"`
}

type Runner struct {
	Name        string  `json:"name"`
	SelectionID uint64  `json:"selection_id"`
	Back        []Price `json:"back"`
	Lay         []Price `json:"lay"`
}

type Price struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}