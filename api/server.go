package api

import (
	"fmt"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/liquiddev99/mantra-interview-be/util"
)

type Server struct {
	router *gin.Engine
	config util.Config
}

func NewServer(config util.Config) (*Server, error) {
	server := &Server{config: config}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	fmt.Println("OriginAllowed", server.config.OriginAllowed)

	router.SetTrustedProxies(
		[]string{
			"103.21.244.0/22",
			"103.22.200.0/22",
			"103.31.4.0/22",
			"104.16.0.0/13",
			"104.24.0.0/14",
			"108.162.192.0/18",
			"131.0.72.0/22",
			"141.101.64.0/18",
			"162.158.0.0/15",
			"172.64.0.0/13",
			"173.245.48.0/20",
			"188.114.96.0/20",
			"190.93.240.0/20",
			"197.234.240.0/22",
			"198.41.128.0/17",
		},
	)

	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = strings.Split(server.config.OriginAllowed, ",")
	corsConf.AllowBrowserExtensions = true
	corsConf.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"}
	corsConf.AllowCredentials = true
	corsConf.AllowHeaders = []string{
		"Content-Type",
		"Authorization",
		"Accept",
		"X-Requested-With",
		"Origin",
		"Access-Control-Request-Headers",
	}

	router.Use(cors.New(corsConf))

	apiRouter := router.Group("/server")
	authRoutes := apiRouter.Group("/")

	apiRouter.GET("/healthcheck", server.healthCheck)

	authRoutes.POST("/upload", server.upload)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func responseError(err error) gin.H {
	return gin.H{"message": err.Error()}
}
