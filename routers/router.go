package routers

import (
	"github.com/codepnw/go-mongo-todos/databases"
	"github.com/codepnw/go-mongo-todos/handlers"
	"github.com/codepnw/go-mongo-todos/services"
	"github.com/gin-gonic/gin"
)

// func NewRouter() *chi.Mux {
// 	router := chi.NewRouter()

// 	router.Use(cors.Handler(cors.Options{
// 		AllowedOrigins:   []string{"https://*", "http://*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	}))

// router.Route("/api", func(r chi.Router) {
// 	r.Route("/v1", func(router chi.Router) {
// 		router.Get("/healthcheck", handlers.HealthCheck)
// 	})
// })

// 	return router
// }

func NewRouter(router *gin.Engine) *gin.Engine {
	service := services.NewTodoService(databases.GetTodosCollection())
	handler := handlers.NewTodoHandler(service)

	group := router.Group("/api/v1")

	group.GET("/", handlers.HealthCheck)

	group.POST("/todos", handler.CreateTodo)
	group.GET("/todos", handler.GetTodos)
	group.GET("/todos/:id", handler.GetTodo)
	group.PATCH("/todos/:id", handler.Update)
	group.DELETE("/todos/:id", handler.DeleteTodo)

	return router
}
