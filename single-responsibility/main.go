package main

import "log"

type Cart struct {
	Items []string
}

func (c *Cart) GetAll() []string {
	return c.Items
}

func (c *Cart) AddItem(item string) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(item string) {
	for i, val := range c.Items {
		if val == item {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
		}
	}
}

type Product struct {
	Products []string
}

func (p *Product) GetAll() []string {
	return p.Products
}

func (p *Product) AddProduct(product string) {
	p.Products = append(p.Products, product)
}

type UserService struct {
	Product *Product
	Cart    *Cart
}

func NewUserService(product *Product, cart *Cart) *UserService {
	return &UserService{
		Product: product,
		Cart:    cart,
	}
}

func (u UserService) GetAll() []string {
	return u.Cart.GetAll()
}

func (u UserService) AddToCart(item string) {
	for _, val := range u.Product.GetAll() {
		if val == item {
			u.Cart.AddItem(item)
		}
	}
}

func (u UserService) RemoveCart(item string) {
	u.Cart.RemoveItem(item)
}

func main() {
	var product Product
	product.AddProduct("T-Shirt")
	product.AddProduct("Jeans")
	product.AddProduct("Hat")

	userOne := NewUserService(&product, &Cart{})
	userOne.AddToCart("T-Shirt")
	userOne.AddToCart("Hat")
	userOne.RemoveCart("Hat")

	cartsOne := userOne.GetAll()

	userTwo := NewUserService(&product, &Cart{})
	userTwo.AddToCart("Jeans")
	userTwo.AddToCart("T-Shirt")

	cartsTwo := userTwo.GetAll()
	log.Println(cartsOne, cartsTwo)
}
