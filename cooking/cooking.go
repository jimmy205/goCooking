package cooking

import (
	"goPra/cooking/world"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

// Cooking 煮菜
func Cooking() {
	r := gin.Default()
	m := melody.New()

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {

		if string(msg) == "customer" {

			c := world.CreatCustomer()
			go func() {
				world.DefaultCity.JoinRestaurant(c)
				world.DefaultWaiter.OrderSomething("buger", c)
			}()
		}

	})

	r.Run(":8000")
}
