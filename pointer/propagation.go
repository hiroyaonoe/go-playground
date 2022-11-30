package main

import (
	"context"
	"fmt"

	bcopprop "github.com/hiroyaonoe/bcop-go/propagation"
	"github.com/hiroyaonoe/bcop-go/protocol/header"
	"go.opentelemetry.io/otel/baggage"
	otelprop "go.opentelemetry.io/otel/propagation"
)

func main() {

	propagator := otelprop.Baggage{}
	ctx := context.Background()

	h1 := header.NewV1("test=testvalue")
	ctx = propagator.Extract(ctx, bcopprop.NewBCoPCarrier(h1))
	fmt.Println(h1)

	fmt.Printf("%s\n", baggage.FromContext(ctx))

	h2 := header.NewV1("")
	propagator.Inject(ctx, bcopprop.NewBCoPCarrier(h2))
	fmt.Println(h2)

	fmt.Printf("%s\n", baggage.FromContext(ctx))
}
