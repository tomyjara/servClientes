package main

import ( "github.com/tomyjara/ServClientes/controller"
	    "github.com/gin-gonic/gin" )

func main() {
    r := gin.Default()

    r.GET("/ListaClientes", controller.Usuarios)
    r.POST("/ListaClientes/:id", controller.UpdateUser)
    r.Run() // listen and server on 0.0.0.0:8080
}