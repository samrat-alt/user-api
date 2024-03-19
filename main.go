package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID         int    `json:"id"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Role       string `json:"role"`
}

var users = []Users{

	{ID: 1, First_Name: "John", Last_Name: "Doe", Email: "jdoe", Password: "123456", Phone: "0909090909", Role: "admin"},
	{ID: 2, First_Name: "Jane", Last_Name: "Doe", Email: "jdoe", Password: "123456", Phone: "0909090909", Role: "user"},
	{ID: 3, First_Name: "Jim", Last_Name: "Doe", Email: "jdoe", Password: "123456", Phone: "0909090909", Role: "admin"},
	{ID: 4, First_Name: "Jack", Last_Name: "Doe", Email: "jdoe", Password: "123456", Phone: "0909090909", Role: "admin"},
	{ID: 5, First_Name: "Jill", Last_Name: "Doe", Email: "jdoe", Password: "123456", Phone: "0909090909", Role: "user"},
}

func getusers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func UserbyID(c *gin.Context) {
	id := c.Param("id")

	// Convert id from string to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	users, err := getusersbyID(idInt)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func getusersbyID(id int) (*Users, error) {
	for i, user := range users {
		if user.ID == id {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func createUsers(c *gin.Context) {
	var newUser Users

	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func main() {
	router := gin.Default()
	router.GET("/users", getusers)
	router.POST("/users", createUsers)
	router.GET("/users/:id", UserbyID)
	router.Run("localhost:9000")
}
