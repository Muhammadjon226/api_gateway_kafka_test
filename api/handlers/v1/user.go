package v1

import (
	"net/http"

	"github.com/Muhammadjon226/api_gateway/api/models"
	l "github.com/Muhammadjon226/api_gateway/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//CreateUser method for create new user
// @Summary Create User
// @Description CreateUser API is for crete new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param create_user body models.CreateUserModel true "create_user"
// @Success 200 {object} models.UserModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/create-user/ [post]
func (h *HandlerV1) CreatUser(c *gin.Context) {

	var (
		body models.CreateUserModel
	)
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	userID, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error while generating uuid")
		h.log.Error("error while generate uuid", l.Error(err))
		return
	}
	user := &models.UserModel{
		Name: body.Name,
		Age:  body.Age,
		ID:   userID.String(),
	}

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource(c.FullPath())
	eventType := "v1.user.created" // ...user.created
	e.SetType(eventType)
	err = e.SetData(cloudevents.ApplicationJSON, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, "error while setting event data")
		h.log.Error("error while setting event data", l.Error(err))
		return
	}
	err = h.kafka.Push(eventType, e)
	if err != nil {
		c.JSON(http.StatusBadRequest, "[pub] failed: failed to publish event ")
		h.log.Error("[pub] failed: failed to publish event ", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, "ok")
}
