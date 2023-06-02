package todos

import (
	"net/http"

	"example.com/restful-api-example/internal/controllers/todos/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodoByID(data_todos *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		todo, err := utils.GetTodoByID(data_todos, id)

		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": err})
			return
		}

		ctx.IndentedJSON(http.StatusOK, todo)
	}
}

func GetTodos(data_todos *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name, err := ctx.Get("uun")
		if !err {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"Message": "User doesn't exist"})
			return
		}

		todo, err2 := utils.GetTodoByOwner(data_todos, name.(string))

		if err2 != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": err2})
			return
		}

		ctx.IndentedJSON(http.StatusOK, todo)
	}
}
