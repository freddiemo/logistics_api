package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/service"
	"github.com/freddiemo/logistics-api/internal/logistics/validators"
)

type MaritimeShipmentController interface {
	Save(ctx *gin.Context) (model.MaritimeShipment, error)
	FindAll(ctx *gin.Context) ([]model.MaritimeShipment, error)
}

type maritimeShipmentController struct {
	maritimeShipmentService service.MaritimeShipmentServiceInterface
}

var validate *validator.Validate

func NewMaritimeShipmentController(maritimeShipmentService service.MaritimeShipmentServiceInterface) MaritimeShipmentController {
	validate = validator.New()
	validate.RegisterValidation("is-fleet-number", validators.IsFleetNumber)
	validate.RegisterValidation("is-guide-number", validators.IsGuideNumber)
	return &maritimeShipmentController{
		maritimeShipmentService: maritimeShipmentService,
	}
}

func (controller *maritimeShipmentController) Save(ctx *gin.Context) (model.MaritimeShipment, error) {
	var maritimeShipment model.MaritimeShipment
	err := ctx.ShouldBind(&maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	err = validate.Struct(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	maritimeShipment, err = controller.maritimeShipmentService.Save(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}

func (controller *maritimeShipmentController) FindAll(ctx *gin.Context) ([]model.MaritimeShipment, error) {
	maritimeShipments, err := controller.maritimeShipmentService.FindAll()
	if err != nil {
		return []model.MaritimeShipment{}, err
	}

	return maritimeShipments, nil
}
