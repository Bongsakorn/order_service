package types

// for more information https://github.com/go-validator/validator/tree/v2.0.1
type ExampleRequest struct {
	Foo string `json:"foo",validate:"required"`
	Bar string `json:"bar",validate:"required"`
}
