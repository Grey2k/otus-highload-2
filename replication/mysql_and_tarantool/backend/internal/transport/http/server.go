package http

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHTTPServer(addr string, endpoints *Endpoints) *http.Server {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"PUT", "POST"}

	router.Use(cors.New(config))

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/sign-up", endpoints.Auth.SignUp)
		authGroup.POST("/sign-in", endpoints.Auth.SignIn)
		authGroup.PUT("/token", endpoints.Auth.RefreshToken)
	}

	router.POST("/questionnaires", endpoints.Social.GetAllQuestionnaires)
	router.GET("/questionnaires", endpoints.Social.GetQuestionnairesByNameAndSurname)

	messengerGroup := router.Group("/messenger")
	{
		messengerGroup.POST("/chat", endpoints.Messenger.CreateChat)
		messengerGroup.GET("/chat", endpoints.Messenger.GetChat)
		messengerGroup.GET("/chats", endpoints.Messenger.GetChats)
		messengerGroup.POST("/messages", endpoints.Messenger.SendMessage)
		messengerGroup.GET("/messages", endpoints.Messenger.GetMessages)
	}

	router.GET("/ws", endpoints.Ws.Connect)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
