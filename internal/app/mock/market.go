package mock

import (
	"github.com/statistico/statistico-bet-finder/internal/app"
	"github.com/statistico/statistico-bet-finder/internal/app/statistico"
	"github.com/stretchr/testify/mock"
)

type MarketBuilder struct {
	mock.Mock
}

func (m MarketBuilder) FixtureAndMarket(f *statistico.Fixture, market string) (*app.Market, error) {
	args := m.Called(f, market)
	return args.Get(0).(*app.Market), args.Error(1)
}
