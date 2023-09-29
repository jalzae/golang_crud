package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
	"rest/config"
	"rest/helper"
	"rest/routes"
	"rest/service/response"
	"rest/system"
	"strings"
)

func main() {

	//CLI
	if len(os.Args) > 3 {
		command := os.Args[1]
		switch command {
		case "migrate:down":
			if len(os.Args) != 3 {
				fmt.Println("Usage: go run main.go migrate:down <filename>")
				return
			}

			db := config.InitDb()
			filename := os.Args[2]
			if err := system.RollbackMigration(db, filename); err != nil {
				fmt.Printf("Error rolling back migration '%s': %v\n", filename, err)
			} else {
				fmt.Printf("Migration '%s' has been rolled back\n", filename)
			}
		}
	}

	//main program
	gin.SetMode(gin.DebugMode)
	system.Migrate()
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRoutesFromFiles(r *gin.Engine) {
	routeDir := "./routes" // Modify this path as needed

	routeSetupFunctions := map[string]func(*gin.Engine){
		"auth":   func(r *gin.Engine) { routes.SetupAuthRoutes("auth", r) },
		"user":   func(r *gin.Engine) { routes.SetupUserRoutes("user", r) },
		"barang": func(r *gin.Engine) { routes.SetupBarangRoutes("barang", r) },
		// Add more route setup functions and filenames as needed
	}

	// Read the list of files in the routes directory
	files, err := filepath.Glob(filepath.Join(routeDir, "*.go"))
	if err != nil {
		fmt.Printf("Error reading route directory: %v\n", err)
		os.Exit(1)
	}

	for _, filePath := range files {
		// Load and run route setup functions from each file
		loadAndRunRouteFile(filePath, r, routeSetupFunctions)
	}
}

func loadAndRunRouteFile(filePath string, r *gin.Engine, routeSetupFunctions map[string]func(*gin.Engine)) {
	// Extract the filename without extension
	routeName := strings.TrimSuffix(filepath.Base(filePath), ".go")

	// Check if a route setup function exists for this filename
	if routeSetupFunc, ok := routeSetupFunctions[routeName]; ok {
		// Call the route setup function
		routeSetupFunc(r)
		fmt.Printf("Registered routes from: %s\n", filePath)
	} else {
		fmt.Printf("Route setup function not found for: %s\n", filePath)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(helper.Cors())

	r.GET("/", func(c *gin.Context) {
		res.Res(c, http.StatusOK, true, "Connected", nil)
	})
	setupRoutesFromFiles(r)
	return r

}
