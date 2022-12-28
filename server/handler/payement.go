package rest

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/lailacha/go-MarketAPI_esgi/server/payement"
	"github.com/lailacha/go-MarketAPI_esgi/server/broadcaster"
)


type GinRestAPI struct {
	payementService payement.Service
	broadcaster	 broadcast.Broadcaster
}


// interface
type RestAPI interface {
	Run()
	Stream()
}

func NewGinRestAPI(payementService payement.Service) *GinRestAPI {
	return &GinRestAPI{
		payementService: payementService,
	}
}

func (g *GinRestAPI) createPayement(ctx *gin.Context) {
	
	// créer un payement voir si ça rentre en bdd

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid id"})
		return
	}

	// créer un payement
	payementObject := payement.Payement{
		ProductID: id,
		PricePaid: "10",
	}
	

	// créer un payement
	g.payementService.Create(payementObject)

	// envoyer le payement à toutes les instances
	g.broadcaster.Submit(payementObject)
	
}

	
func (g *GinRestAPI) Run() {
	router := gin.Default()
	router.POST("/payement/:id", g.createPayement)
	router.Run(":8080")
}

// service de stream permettant d'envoyer à toutes les instances les modifications de la bdd
func (g *GinRestAPI) Stream() {
	
	// creation d'un broadcaster
	g.broadcaster = broadcast.NewBroadcaster(10)
}
