package adapter

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lailacha/go-MarketAPI_esgi/server/broadcaster"
	"github.com/lailacha/go-MarketAPI_esgi/server/payement"
	//"github.com/lailacha/go-MarketAPI_esgi/server/product"
)

type GinAdapter interface {
	Stream(c *gin.Context)
	Submit(c *gin.Context)
	CreatePayement(c *gin.Context)
}

type ginAdapter struct {
	broadcaster broadcast.Broadcaster
	//productService product.Service
	payementService payement.Service
}

type Message struct
{
	UserId string
	Text string
}


func NewGinAdapter(broadcaster broadcast.Broadcaster, payementService payement.Service) *ginAdapter {
	return &ginAdapter{
		broadcaster: broadcaster,
		payementService: payementService,
	}
}

// Stream is the handler for the stream endpoint
func (adapter *ginAdapter) Stream(c *gin.Context) {
	

	//create a new channel to handle the stream
	listener := make(chan interface{})

	// get the broadcaster

	adapter.broadcaster.Register(listener)

	//close the channel when error message or client is gone
	defer adapter.broadcaster.Unregister(listener)

	clientGone := c.Request.Context().Done()

	c.Stream(func(w io.Writer) bool {
		select {
		case <-clientGone:
			return false
		case message := <-listener:
			serviceMsg, ok := message.(Message)
			if !ok {
				fmt.Println("not a message")
				c.SSEvent("message", message)
				return false
			}
			c.SSEvent("message", " "+serviceMsg.UserId+" → "+serviceMsg.Text)
			return true
		}
	})



	fmt.Println("stream is OK")
}

func (adapter *ginAdapter) CreatePayement (c *gin.Context) {
	
	//get POST data


	fmt.Println("create payement", c.PostForm("id"))

	id, err := strconv.Atoi(c.PostForm("id"))


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		fmt.Println(err)
		return
	}
		price := c.PostForm("price")


		// get the broadcaster
		b := adapter.broadcaster

		// save the payement
		adapter.payementService.Create(id, price);


		b.Submit(Message{
			UserId: "1",
			Text: "Payement is created",
		})

		c.JSON(http.StatusOK, gin.H{"message": "Payement is created"})
	
		fmt.Println("submit is OK")

}


func (adapter *ginAdapter) Submit(c *gin.Context) {
	
	// get the broadcaster
	b := adapter.broadcaster


	b.Submit("testMessag")

	fmt.Println("submit is OK")

}