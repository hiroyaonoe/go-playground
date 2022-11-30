module github.com/hiroyaonoe/go-playground

go 1.18

replace github.com/hiroyaonoe/bcop-go => ../bcop-go

require (
	github.com/hiroyaonoe/bcop-go v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.11.1
)

require go.opentelemetry.io/otel/trace v1.11.1 // indirect
