package routes

import (
	"example.com/restful-api-example/configs"
	"example.com/restful-api-example/internal/controllers/todos"
	"example.com/restful-api-example/internal/controllers/token"
	"example.com/restful-api-example/internal/controllers/user"
	"example.com/restful-api-example/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	route := gin.Default()

	api := route.Group("/api")
	{
		api.POST("/token", token.GenerateToken)
		api.POST("/user/register", user.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", token.Ping)

			secured.GET("/todos", todos.GetTodos(configs.Db_todos))

			secured.GET("/todo/:id", todos.GetTodoByID(configs.Db_todos))

			secured.POST("/todos", todos.CreateTodo(configs.Db_todos))

			secured.PATCH("/todo_change/:id", todos.ChangeTodo(configs.Db_todos))

			secured.DELETE("delete_todo/:id", todos.DeleteTodo(configs.Db_todos))
		}
	}

	return route

}
