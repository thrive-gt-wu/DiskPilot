package main

import (
	"DiskPilot/bootstrap"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	bootstrap.SetupRoute(r)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}