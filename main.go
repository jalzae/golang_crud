package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"rest/response"
	"rest/routes"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
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
	r.Use(Cors())

	r.GET("/", func(c *gin.Context) {
		res.Res(c, http.StatusOK, true, "Connected", nil)
	})
	setupRoutesFromFiles(r)
	return r

}

// Cors handles cross-domain requests and supports options access
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT,DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//Release all OPTIONS methods
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// process request
		c.Next()
	}
}
