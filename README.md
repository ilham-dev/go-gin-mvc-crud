# API-login-go-mysql
login API with bcrypt

# How to Run ?
##1. Create The Database with name go_api
##2. Config DB
      replace test with your db name
      db, err := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
##get library
you should import the library first in terminal

	go get github.com/gin-gonic/gin

	go get golang.org/x/crypto/bcrypt

	go get github.com/go-sql-driver/mysql
	go get github.com/kataras/go-sessions

### And here we go 
	go run main.go