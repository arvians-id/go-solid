package main

import "log"

type ProductService struct {
}

func NewProductService() FullQuery {
	return &ProductService{}
}
func (p *ProductService) Create() {
	log.Println("Creating product")
}

func (p *ProductService) Delete() {
	log.Println("Deleting product")
}

func (p *ProductService) FindAll() {
	log.Println("Show all products list")
}

func (p *ProductService) FindByID() {
	log.Println("Show product by ID")
}
