package controllers

import (
	"brenimsilva/auth/initializers"
	"brenimsilva/auth/models"
	"net/http"
	"os"
	"time"

    "github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
    var body struct {
        Email string
        Password string
    }

    if c.BindJSON(&body) != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Failed to read body",
        })
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

    if(err != nil) {
        c.JSON(http.StatusBadRequest, gin.H {
            "error": "Failed do hash password",
        })

        return
    }

    user := models.User{Name: "", Email: body.Email, Password: string(hash)}

    result := initializers.DB.Create(&user)

    if(result.Error != nil) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to create user",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": "User created",
    })

}

func Login(c *gin.Context) {
    var body struct {
        Email string
        Password string
    }

    if c.BindJSON(&body) != nil {
        c.IndentedJSON(http.StatusBadRequest, gin.H{
            "error": "Failed to read body",
        })
        return
    }

    var user models.User
    initializers.DB.First(&user, "email = ?", body.Email)

    if(user.ID == 0) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email or password",
        })
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

    if(err != nil) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email or password",
        })
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodPS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
    })

    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if(err != nil) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Failed to create token",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}
