package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/apodicticscott/oaas/graph/generated"
	"github.com/apodicticscott/oaas/graph/resolvers"
	"github.com/apodicticscott/oaas/internal/api"
	"github.com/apodicticscott/oaas/internal/persistence"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// getEnvOrDefault returns the value of the environment variable or the default value if not set
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	log.Println("Starting OaaS server...")

	// Try to load .env file from different possible locations
	envPaths := []string{".env", "../.env", "../../.env"}
	envLoaded := false

	for _, path := range envPaths {
		if _, err := os.Stat(path); err == nil {
			log.Printf("Loading environment from %s...", path)
			if err := godotenv.Load(path); err != nil {
				log.Printf("Error loading environment file: %v\n", err)
			} else {
				envLoaded = true
				log.Printf("Successfully loaded environment from %s", path)
				break
			}
		}
	}

	if !envLoaded {
		log.Println("No environment file found, using system environment variables")
	}

	// Get database connection parameters from environment variables or use defaults
	// Determine if we're running locally or in Docker
	isLocal := true
	_, err := net.LookupHost("db")
	if err == nil {
		// If we can resolve 'db', we're probably running in Docker
		isLocal = false
	}

	// Set database connection parameters based on environment
	var dbHost, dbPort string
	if isLocal {
		log.Println("Running in local environment, using localhost:5433")
		dbHost = "localhost"
		dbPort = "5433"
	} else {
		log.Println("Running in Docker environment, using db:5432")
		dbHost = "db"
		dbPort = "5432"
	}

	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "ontology")

	// Log the database connection parameters for debugging
	log.Printf("Database connection parameters: host=%s, port=%s, user=%s, dbname=%s", dbHost, dbPort, dbUser, dbName)

	// DSN for Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	log.Println("Connecting to PostgreSQL database...")
	db, err := persistence.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	log.Println("Database connected successfully")

	// Initialize API handler
	apiHandler := api.NewHandler(db)

	// Initialize GraphQL resolver
	resolver := &resolvers.Resolver{DB: db}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	// Setup Gin router
	router := gin.Default()

	// Health check
	router.GET("/health", apiHandler.HealthCheck)

	// REST API routes
	api := router.Group("/api/v1")
	{
		// Substances
		api.GET("/substances", apiHandler.GetSubstances)
		api.GET("/substances/:id", apiHandler.GetSubstance)
		api.POST("/substances", apiHandler.CreateSubstance)
		api.PUT("/substances/:id", apiHandler.UpdateSubstance)
		api.DELETE("/substances/:id", apiHandler.DeleteSubstance)

		// Kinds
		api.GET("/kinds", apiHandler.GetKinds)
		api.POST("/kinds", apiHandler.CreateKind)

		// Attributes
		api.GET("/attributes", apiHandler.GetAttributes)
		api.POST("/attributes", apiHandler.CreateAttribute)

		// Modes
		api.GET("/modes", apiHandler.GetModes)
		api.POST("/modes", apiHandler.CreateMode)

		// Causality
		api.GET("/substances/:id/causes", apiHandler.GetCauses)
		api.POST("/causes", apiHandler.AddCause)

		// Potentialities
		api.GET("/potentialities", apiHandler.GetPotentialities)
		api.POST("/potentialities", apiHandler.CreatePotentiality)
		api.POST("/potentialities/:id/actualize", apiHandler.ActualizePotentiality)
		api.GET("/potentialities/:id/conditions", apiHandler.CheckConditions)

		// Evolution
		api.GET("/substances/:id/evolution", apiHandler.GetSubstanceEvolution)
	}

	// GraphQL routes
	router.GET("/playground", gin.WrapF(playground.Handler("GraphQL playground", "/query")))
	router.POST("/query", gin.WrapH(srv))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OaaS API is running",
			"endpoints": gin.H{
				"rest_api":           "http://localhost:8080/api/v1/",
				"graphql_playground": "http://localhost:8080/playground",
				"graphql_endpoint":   "http://localhost:8080/query",
			},
		})
	})

	log.Println("🚀 Server running on http://localhost:8080/")
	log.Println("REST API available at http://localhost:8080/api/v1/")
	log.Println("GraphQL playground available at http://localhost:8080/playground")
	log.Println("GraphQL endpoint available at http://localhost:8080/query")
	log.Fatal(router.Run(":8080"))
}
