package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lailacha/go-MarketAPI_esgi/server/broadcaster/broadcast"
)

type GinRestAPI struct {
	suscriberManager broadcast.Broadcaster
}


func NewGinRestAPI(suscriberManager service.Manager) *GinRestAPI {
	return &GinRestAPI{
		suscriberManager: suscriberManager,
	}
}

func (g *GinRestAPI) createPayement(ctx *gin.Context) {
	
}

	