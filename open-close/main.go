package main

import (
	"log"
)

type Shipping interface {
	Ship() string
}

type StandardShipping struct {
}

func (s *StandardShipping) Ship() string {
	return "Using standard shipping"
}

type ExpressShipping struct {
}

func (e *ExpressShipping) Ship() string {
	return "Using standard shipping"
}

type CargoShipping struct {
}

func (c *CargoShipping) Ship() string {
	return "Using standard shipping"
}

func ProcessOrder(shipping Shipping) {
	log.Println("Processing an order...")
	log.Println(shipping.Ship())
	log.Println("Order has been processed")
}

func main() {
	var standardShipping StandardShipping
	var expressShippinng ExpressShipping
	var cargoShipping CargoShipping

	ProcessOrder(&standardShipping)
	ProcessOrder(&expressShippinng)
	ProcessOrder(&cargoShipping)
}
