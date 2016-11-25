package clientes

import (
   "fmt"
   "database/sql"
 _ "github.com/go-sql-driver/mysql"
 "github.com/gin-gonic/gin" 
)
type Cliente struct{
    Id int
    Nombre string
    Apellido string
}
func ObtenerClientes() []Cliente{
// tomas := Cliente{39466010, "Tomas", "Jaratz"}


    db, err := sql.Open("mysql", "root:tomyjara@/AlarmaSys")
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

// Open doesn't open a connection. Validate DSN data:
  rows, err := db.Query("SELECT id, nombre, apellido FROM cliente")
  clientes := make([]Cliente,0)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Fetch rows
    for rows.Next() {
        var cliente Cliente    

        err = rows.Scan(&cliente.Id, &cliente.Nombre, &cliente.Apellido )
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
        clientes = append(clientes,cliente)
        
    }
    if err = rows.Err(); err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    
    return clientes
}


// var dbmap  = initDb()

// func initDb() {

    // db, err := sql.Open("mysql", "root:tomyjara@/AlarmaSys")
    // if err != nil {
    //     panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    // }
    // defer db.Close()
// }

func CambiarUser(c *gin.Context, id string) Cliente{

    fmt.Println("AAAAAAAAAAAAAAAAA")
    db, err := sql.Open("mysql", "root:tomyjara@/AlarmaSys")
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    //Traer datos usuario en bd
    var cliente Cliente
    // err := dbmap.SelectOne(&cliente, "SELECT * FROM user WHERE id=?", id)
    rows, err := db.Query("SELECT id FROM cliente WHERE id = "+string(id))
    err = rows.Scan(&id)

    var json Cliente
    c.Bind(&json)

    // user_id := int(id)

    cliente = Cliente{
        Id:        int(id),
        Nombre:    json.Nombre,
        Apellido:  json.Apellido,
    }

    if cliente.Nombre != "" && cliente.Apellido != "" {
        //actualizar datos usuario en bd
        
        _, err = db.Query("UPDATE CLIENTE SET nombre =" + cliente.Nombre + "apellido=" + cliente.Apellido + "WHERE id=" + string(id))


    }
    return cliente        
}

