package controller
import(
	"log"
	//import library
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//import API packages
	"github.com/miraikeitai2020/michisirube_camp_backend/evaluation/pkg/server/model"
	"github.com/miraikeitai2020/michisirube_camp_backend/evaluation/pkg/evaluation"

)
var(
	eva 	model.Request.Evaluation
	emo		model.Request.Evaluation.Emotion
	state	model.Request.Location

)
type Controller struct {
	DB *gorm.DB
}
func(ctrl *Controller)/(context *gin.Context){
	context.JSON(500,gin.H{
		"status": evaluation.Evaluation(eva.Value,emo.Value,state.Id),
	})
}