package main

import (
	"fmt"

	"github.com/elanticrytp0/go4it/go4it"
)

func main() {

	app := go4it.NewApp("")

	// app := core.LoadAppConfig()

	fmt.Printf("Config: \n %v \n", app)
	fmt.Printf("Config: \n %v \n", app.Config.App_name)
	fmt.Printf("Config: \n %v \n", app.Config.App_version)

	fmt.Printf("Conexiones: \n %v \n", app.Config.DB)
	app.Connect2Db("local")
	fmt.Printf("Config: \n %v \n", app.DB)
	fmt.Printf("Config: \n %v \n", app.DB[0])
	fmt.Printf("Config: \n %v \n", app.DB[0].Conn.Debug().Statement.TableExpr.Vars...)

	// fmt.Printf("Conexiones: \n %v \n", app.Config.DB["conn1"])

}
