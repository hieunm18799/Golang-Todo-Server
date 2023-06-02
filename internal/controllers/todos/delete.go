package todos

import (
	"net/http"

	"example.com/restful-api-example/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DeleteTodo(data_todos *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		var todo models.Todo
		name, err := ctx.Get("uun")
		if !err {
			ctx.IndentedJSON(http.StatusConflict, gin.H{"Message": "User doesn't exist"})
			return
		}

		// clause is needed for returnning data
		res := data_todos.Clauses(clause.Returning{}).Where("id = ?", id).Where("owner = ?", name).Delete(&todo)

		if res.Error != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": res.Error.Error()})
			return
		}

		if res.RowsAffected <= 0 {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, &todo)
	}
}
