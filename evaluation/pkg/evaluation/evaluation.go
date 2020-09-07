package evaluation

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/doalects/mysql"
	"github.com/miraikeitai2020/michisirube_camp_backend/evaluation/pkg/config"
	"github.com/miraikeitai2020/michisirube_camp_backend/evaluation/pkg/server/model"


)

type Locatetable struct{
	gorm.Model
	
	locate 		string
	latitude	float32
	longitude/	float32
}
type Emotionaltable struct{
	gorm.Model
	
	glad		int
	anger		int
	sad			int
	pleasant	int

}
type Logtable	struct{
	gorm.Model
	
	count		int
	timestamp	string
	emotion		int
	value		int
}
func DBinit(){
	token:=GetConnectionToken()
	db,err:=gorm.Open(mysql,token)
	if err !=nil{
		log.Fatal(err)
	}
	defer db.Close()
	do.Set("gorm:table_options","ENGINE"=InnoDB).AutoMigrate(&Locatetable{},&Emotionaltable{},&Logtable{})
return db
}
func Evaluation( int)bool{

	db:=DBinit()
}