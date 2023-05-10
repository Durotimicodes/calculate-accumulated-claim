package models

//type ProductTriangle represents the structure of the input data
type ProductTriangle struct {
	Product         string
	Origin          int
	DevelopmentYear int
	Value           float64
}

//type Triangle is a datastructure to store products with its corresponding data
type Triangle map[string]ProductTriangle

//type FinalValue represents the structure of the output data
type FinalValue struct {
	Origin           int
	AccumulatedClaim float64
}
