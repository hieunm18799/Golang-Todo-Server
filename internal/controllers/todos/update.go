package todos

import (
	"net/http"

	"example.com/restful-api-example/internal/controllers/todos/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ChangeTodo(data_todos *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		change_todo, err := utils.GetTodoByID(data_todos, id)

		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": err})
			return
		}

		if err := ctx.BindJSON(change_todo); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "No data change found"})
			return
		}

		name, err2 := ctx.Get("uun")
		if !err2 {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"Message": "User doesn't exist"})
			return
		}

		if name != change_todo.Owner {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"Message": "You can't change this todo"})
			return
		}

		data_todos.Save(&change_todo)
		ctx.IndentedJSON(http.StatusOK, change_todo)
	}
}
