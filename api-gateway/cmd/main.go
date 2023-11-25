package main

import (
	"log"

	"github.com/RohithER12/api-gateway/pkg/auth"
	"github.com/RohithER12/api-gateway/pkg/cart"
	"github.com/RohithER12/api-gateway/pkg/config"
	"github.com/RohithER12/api-gateway/pkg/order"
	"github.com/RohithER12/api-gateway/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	cart.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}