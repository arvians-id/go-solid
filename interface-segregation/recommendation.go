package main

import "log"

type RecommendationService struct {
}

func NewRecommendationService() Retrieve {
	return &RecommendationService{}
}

func (r RecommendationService) FindAll() {
	log.Println("Show all recommendations list")
}

func (r RecommendationService) FindByID() {
	log.Println("Show all recommendation by product ID")
}
