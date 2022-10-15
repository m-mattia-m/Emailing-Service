package api

import (
	"emailing-service/database"
	"emailing-service/helpers"
	"emailing-service/models"
	"emailing-service/pkg/emailing"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func Router() {
	port, err := helpers.GetEnv("PORT")
	if err != nil {
		log.Println("[GET-ENV]: Get Port was failed \t\t -> Error: " + err.Error())
		return
	}

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/email", auth, SendMail)
		v1.POST("/register", adminAuth, Register)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":" + port)
}

func auth(c *gin.Context) {
	APIKey := c.Request.Header.Get("X-API-Key")

	status, err := database.Auth(APIKey)
	if err != nil {
		log.Println("[API-AUTH]: An error occurred during authentication \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if status {
		log.Println("[API-AUTH]: Successfully Loged in")
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
}

func adminAuth(c *gin.Context) {
	requestUsername, requestPassword, hasAuth := c.Request.BasicAuth()
	envUsername, err := helpers.GetEnv("ADMIN_USERNAME")
	if err != nil {
		log.Println("[API-AUTH]: An error occurred during authentication \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	envPassword, err := helpers.GetEnv("ADMIN_PASSWORD")
	if err != nil {
		log.Println("[API-AUTH]: An error occurred during authentication \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if !hasAuth {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	if requestUsername != envUsername {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	if requestPassword != envPassword {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
	}
	log.Println("[ADMIN-AUTH]: Successfully Loged in")
}

// SendMail 					godoc
// @title           			SendMail
// @description     			Send a email to the receiver.
// @Router  					/email [post]
// @Tags 						Emailing-Service
// @Accept 						json
// @Produce						json
// @Security					ApiKeyAuth
// @Param						ReceiverRequst body models.ReceiverRequst true "ReceiverRequst"
// @Success      				200  {object} models.ReceiverRequst
// @Failure      				400  {object} models.HttpError
// @Failure      				404  {object} models.HttpError
// @Failure      				500  {object} models.HttpError
func SendMail(c *gin.Context) {
	receiver := models.ReceiverRequst{}
	if err := c.BindJSON(&receiver); err != nil {
		log.Println("[API-MAIL]: There is no application registered with this keyword \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	sender := models.Sender{}
	err := helpers.DB.First(&sender, "keyword = ?", receiver.Keyword).Error
	if err != nil {
		log.Println("[API-MAIL]: There is no application registered with this keyword \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	emailing.Mail(sender, receiver)
	c.JSON(http.StatusOK, sender.SenderRequest)
}

// Register 					godoc
// @title           			Register
// @description     			Register a new Application-Email
// @Router  					/register [post]
// @Tags 						Emailing-Service
// @Accept 						json
// @Produce						json
// @Security					ApiKeyAuth
// @Param						SenderRequest body models.SenderRequest true "SenderRequest"
// @Success      				200  {object} models.SenderRequest
// @Failure      				400  {object} models.HttpError
// @Failure      				404  {object} models.HttpError
// @Failure      				500  {object} models.HttpError
func Register(c *gin.Context) {
	senderRequest := models.SenderRequest{}
	if error := c.BindJSON(&senderRequest); error != nil {
		c.JSON(http.StatusInternalServerError, error.Error())
		return
	}

	token, err := helpers.GenerateRandomStringURLSafe(128)
	if err != nil {
		log.Println("[REGISTER]: Error while generating a token \t\t -> Error: " + err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	var sender models.Sender
	sender.SenderRequest = senderRequest
	sender.Token = token
	response := helpers.DB.Create(&sender)
	if response.Error != nil {
		c.JSON(http.StatusInternalServerError, response.Error)
		return
	}
	c.JSON(http.StatusOK, sender)
}
