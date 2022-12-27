package main

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticket string) (float64, error) {
	return MockPriceFetcher(ctx, ticket)
}

var priceMocks = map[string]float64{
	"BTC": 20000.0,
	"ETH": 200.0,
	"GG":  100000.0,
}

func MockPriceFetcher(ctx context.Context, ticket string) (float64, error) {
	price, ok := priceMocks[ticket]

	if !ok {
		return price, fmt.Errorf("the given ticket (%s) is not supported", ticket)
	}

	return price, nil
}
