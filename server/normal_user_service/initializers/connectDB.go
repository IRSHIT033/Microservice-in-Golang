package initializers

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB  
  
func Connect_DB(){  
  var err error	
  dsn := os.Getenv("DB_ACCESS_INIT")
  //fmt.Println(dsn)

  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
     panic("failed to conenct to db")  
  }
  
}  