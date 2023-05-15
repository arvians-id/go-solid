package main

type DataManipulation interface {
	Create()
	Delete()
}

type Retrieve interface {
	FindAll()
	FindByID()
}

type FullQuery interface {
	DataManipulation
	Retrieve
}
