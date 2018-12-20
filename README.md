# go-gin-mvc-crud
    Simple API with go ,gin and mysql

# How to Run ?
## 1. Create The Database with name go_api
## 2. Config DB
    replace test with your db go_api
    db, err := gorm.Open("mysql", "root:@/go_api?charset=utf8&parseTime=True&loc=Local")
## Optional Install Glide
    1. Mac
       brew install glide
    2. Linux
       sudo apt-get install glide
    
    after install glide
    1. glide init
    2. glide install
    
## get library
you should import the library first in terminal

	go get github.com/gin-gonic/gin

	go get golang.org/x/crypto/bcrypt

	go get github.com/go-sql-driver/mysql
	go get github.com/kataras/go-sessions

### And here we go 
	go run main.go