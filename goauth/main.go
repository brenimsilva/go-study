package main

import (
	"brenimsilva/auth/controllers"
	"brenimsilva/auth/initializers"

	"github.com/gin-gonic/gin"
)

const resource string = "/auth"
const uri string = "localhost"
const port string = "9090"
const URL string = uri + ":" + port

func init() {
    initializers.LoadEnvVariables()
    initializers.ConnectToDb()
    initializers.Migration()
}

func main() {
    router := gin.Default()
    router.GET(resource, controllers.Login)
    router.POST(resource, controllers.Signup)
    router.Run()
}



