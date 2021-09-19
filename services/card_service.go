package services

import (
	"errors"
	"strconv"

	"github.com/Israel-Ferreira/challenge-conductor/models"
	"github.com/gin-gonic/gin"
)

var cards = []models.Card{
	{ID: 2490, Account: models.Account{Id: 1, Status: "active"}, Amount: 300.00},
	{ID: 2520, Account: models.Account{Id: 2, Status: "unactive"}, Amount: 0.00},
}

func findById(cards []models.Card, id uint) (models.Card, error) {

	var searchedCards []models.Card

	for _, card := range cards {
		if card.ID == id {
			searchedCards = append(searchedCards, card)
		}
	}

	if len(searchedCards) == 0 {
		return models.Card{}, errors.New("cartão não encontrado")
	} else {
		return searchedCards[0], nil
	}

}

func GetCardById(c *gin.Context) {
	pathVar := c.Params.ByName("id")
	id, err := strconv.ParseInt(pathVar, 10, 64)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erro ao converter o id",
		})
	}

	card, err := findById(cards, uint(id))

	if err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}else{
		c.JSON(200, card)
	}

	
}
