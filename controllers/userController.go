package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all users")
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get a user by id")
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "User signed up")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "User logged in")
	}
}

func HashPassword(password string) string {
	// Implement password hashing logic here
	return "hashed_" + password
}

func CheckPasswordHash(userPassword string, providePassword string) (bool, string) {
	// Implement password hash checking logic here
	return userPassword == "hashed_"+providePassword, "user_id_example"
}
