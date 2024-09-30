package main

import (
	"github.com/fhorray/gingon/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
   r := gin.Default()

	 routes.SetupPostRoutes(r)
	 routes.SetupUserRoutes(r)

	 r.Run(":3000")
}
