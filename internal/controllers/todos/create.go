package todos

import (
	"net/http"

	"example.com/restful-api-example/internal/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func CreateTodo(data_todos *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var new_todo *models.Todo

		if err := ctx.BindJSON(&new_todo); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "No create data found"})
			return
		}

		name, err := ctx.Get("uun")
		if !err {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"Message": "User doesn't exist"})
			return
		}

		new_todo.Owner = name.(string)

		if err := data_todos.Create(&new_todo).Error; err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "JSON type is wrong!"})
			return
		}
		ctx.IndentedJSON(http.StatusCreated, new_todo)
	}
}
