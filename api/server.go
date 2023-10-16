package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	db "github.com/minhtri6179/manata/db/sqlc"
	"github.com/minhtri6179/manata/tokenprovider"
	tokenJWT "github.com/minhtri6179/manata/tokenprovider/jwt"
	"github.com/minhtri6179/manata/util"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker tokenprovider.TokenProvider
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := tokenJWT.NewJWTProvider(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cưannot create token: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	router.Use(cors.Default())

	router.GET("/ping", server.pong)
	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/register", server.registerUser)
			user.POST("/login", server.loginUser)

		}
		task := v1.Group("/tasks")
		{
			task.POST("/create", server.createTask)
			task.GET("/list", server.listTask)
			task.GET("/:id", server.getTask)
			task.PUT("/:id", server.updateTask)
			task.DELETE("/:id", server.deleteTask)
		}
		docs := v1.Group("/docs")
		{
			docs.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

	}
	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
