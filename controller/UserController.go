package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"ilham/structs"
	"net/http"
)

// to get one data with {id}
func (idb *InDB) GetUser(c *gin.Context) {
	var (
		user structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": user,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// get all data in person
func (idb *InDB) GetUsers(c *gin.Context) {
	var (
		persons []structs.User
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"status": 200,
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// create new data to the database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
	)
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	person.First_Name = first_name
	person.Last_Name = last_name
	person.Email = email
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	person.Password = string(hashedPassword)
	idb.DB.Create(&person)
	result = gin.H{
		"status": 200,
		"result": person,
	}
	c.JSON(http.StatusOK, result)
}

// update data with {id} as query
func (idb *InDB) UpdateUser(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    structs.User
		newPerson structs.User
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	println(err)
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// delete data with {id}
func (idb *InDB) DeleteUser(c *gin.Context) {
	var (
		person structs.User
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"status": 404,
			"result": "data not found",
		}
	}
	println(id)
	err = idb.DB.Where("id = ?",id).Delete(&person).Error
	fmt.Println(err)
	if err != nil {
		result = gin.H{
			"status": 400,
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"status": 200,
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

func Index(c *gin.Context)  {
	c.JSON(200, gin.H{
		"status": 200,
		"message": "ok welcome to golang",
	})
}