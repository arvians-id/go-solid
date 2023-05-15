package main

import "log"

type CategoryService struct {
}

func NewCategoryService() DataManipulation {
	return &CategoryService{}
}

func (c CategoryService) Create() {
	log.Println("Creating category")
}

func (c CategoryService) Delete() {
	log.Println("Deleting category")
}
