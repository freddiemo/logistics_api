package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"strconv"

	"github.com/freddiemo/logistics-api/internal/clients/model"
	"github.com/freddiemo/logistics-api/internal/clients/service"
	"github.com/freddiemo/logistics-api/validators"
)

type ClientController interface {
	Save(ctx *gin.Context) (model.Client, error)
	FindAll(ctx *gin.Context) ([]model.Client, error)
	FindById(ctx *gin.Context) (model.Client, error)
	Update(ctx *gin.Context) (model.Client, error)
}

type clientController struct {
	clientService service.ClientServiceInterface
}

var validate *validator.Validate

func NewClientController(clientService service.ClientServiceInterface) ClientController {
	validate = validator.New()
	validate.RegisterValidation("is-email", validators.IsEmail)
	return &clientController{
		clientService: clientService,
	}
}

func (controller *clientController) Save(ctx *gin.Context) (model.Client, error) {
	var client model.Client
	err := ctx.ShouldBind(&client)
	if err != nil {
		return model.Client{}, err
	}

	err = validate.Struct(client)
	if err != nil {
		return model.Client{}, err
	}

	client, err = controller.clientService.Save(client)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (controller *clientController) FindAll(ctx *gin.Context) ([]model.Client, error) {
	clients, err := controller.clientService.FindAll()
	if err != nil {
		return []model.Client{}, err
	}

	return clients, nil
}

func (controller *clientController) FindById(ctx *gin.Context) (model.Client, error) {
	var client model.Client

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.Client{}, err
	}

	client, err = controller.clientService.FindById(id)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}

func (controller *clientController) Update(ctx *gin.Context) (model.Client, error) {
	var client model.Client
	err := ctx.ShouldBind(&client)
	if err != nil {
		return model.Client{}, err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.Client{}, err
	}

	err = validate.Struct(client)
	if err != nil {
		return model.Client{}, err
	}

	client.ID = uint(id)
	client, err = controller.clientService.Update(client)
	if err != nil {
		return model.Client{}, err
	}

	return client, nil
}
