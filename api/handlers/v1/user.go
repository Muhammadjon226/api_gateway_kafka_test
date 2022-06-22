package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Muhammadjon226/api_gateway/api/models"
	l "github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/Muhammadjon226/api_gateway/pkg/utils"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

//CreateUser method for create new user
// @Summary Create User
// @Description CreateUser API is for crete new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param create_user body models.CreateUserModel true "create_user"
// @Success 200 {object} models.User
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
	user := &models.User{
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

//ListUsers method for get list of users
// @Summary List Users
// @Description ListUsers API is for get list of users
// @Tags user
// @Accept  json
// @Produce  json
// @Param page query int false "page"
// @Param limit query int false "limit"
// @Success 200 {object} models.ListUserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/list-users/ [get]
func (h *HandlerV1) ListUsers(c *gin.Context) {

	var response *models.ListUserResponse

	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)

	if errStr != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	URL := "http://" + h.cfg.UserServiceURL + "/v1/user/list-users/"
	// httpClient := &http.Client{}
	// req, err := http.NewRequest(
	// 	"GET", URL, nil,
	// )
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// 	c.JSON(http.StatusInternalServerError, err)
	// }
	// req.Header.Add("Accept", "application/json")
	// q := req.URL.Query()
	// q.Add("limit", strconv.Itoa(int(params.Limit)))
	// q.Add("page", strconv.Itoa(int(params.Page)))

	// req.URL.RawQuery = q.Encode()

	// response, err := httpClient.Do(req)

	client := resty.New()
	client.DisableWarn = true
	resp, err := client.R().
		SetQueryParam("limit", strconv.Itoa(int(params.Limit))).
		SetQueryParam("page", strconv.Itoa(int(params.Page))).Get(URL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(resp.Body(), &response)
	if err != nil {
		log.Println("error while unmarshalling response", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}
