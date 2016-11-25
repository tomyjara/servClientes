package controller

import (
	"fmt"
    "github.com/gin-gonic/gin"
    "github.com/tomyjara/ServClientes/clientes"
)

func Usuarios(c *gin.Context) {
        listaClientes := clientes.ObtenerClientes()
        // content := gin.H{}
        // for index, value := range listaClientes{
        //     content[strconv.Itoa(index)] = value
        // }
        c.JSON(200, listaClientes)
}

func UpdateUser(c *gin.Context) {
	
	fmt.Println("ABAAAAAAAAAAAAAAAA")
	id := c.Params.ByName("id")
	cliente := clientes.CambiarUser(c,id)    
    c.JSON(200, cliente)
  	
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

