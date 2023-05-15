package main

func main() {
	productService := NewProductService()
	categoryService := NewCategoryService()
	recommendationService := NewRecommendationService()

	productService.FindAll()
	productService.FindByID()
	productService.Create()
	productService.Delete()

	categoryService.Create()
	categoryService.Delete()

	recommendationService.FindAll()
	recommendationService.FindByID()
}
