package main

import (
	_"XCHANGO_APIS_CRUD/routers"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// 🔹 Cargar .env
	err := godotenv.Load()
	if err != nil {
		panic("Error cargando archivo .env")
	}

	// 🔹 Leer variables XCHANGO
	pgUser := os.Getenv("XCHANGO_PGUSER")
	pgPass := os.Getenv("XCHANGO_PGPASS")
	pgHost := os.Getenv("XCHANGO_PGHOST")
	pgPort := os.Getenv("XCHANGO_PGPORT")
	pgDb := os.Getenv("XCHANGO_PGDB")
	pgSchema := os.Getenv("XCHANGO_PGSCHEMA")

	runMode := os.Getenv("XCHANGO_RUN_MODE")
	httpPort := os.Getenv("XCHANGO_HTTP_PORT")

	// 🔹 Validación básica
	if pgUser == "" || pgPass == "" || pgHost == "" || pgPort == "" || pgDb == "" {
		panic("Variables .env incompletas")
	}

	// 🔹 Modo de ejecución
	if runMode != "" {
		beego.BConfig.RunMode = runMode
	}

	// 🔹 Puerto HTTP
	if httpPort != "" {
		if port, err := strconv.Atoi(httpPort); err == nil {
			beego.BConfig.Listen.HTTPPort = port
		}
	}

	// 🔹 Conexión PostgreSQL
	sqlConn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s",
		pgUser, pgPass, pgHost, pgPort, pgDb, pgSchema,
	)

	fmt.Println("PostgreSQL:", sqlConn)

	// 🔹 Registrar DB
	orm.RegisterDataBase("default", "postgres", sqlConn)

	// 🔹 Swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 🔹 CORS
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 🔹 Iniciar servidor
	beego.Run()
}