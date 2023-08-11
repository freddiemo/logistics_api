package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/maritime_shipment/service"
	"github.com/freddiemo/logistics-api/internal/logistics/validators"
)

type MaritimeShipmentController interface {
	Save(ctx *gin.Context) (model.MaritimeShipment, error)
	FindAll(ctx *gin.Context) ([]model.MaritimeShipment, error)
	FindById(ctx *gin.Context) (model.MaritimeShipment, error)
	Update(ctx *gin.Context) (model.MaritimeShipment, error)
	Delete(ctx *gin.Context) error
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
	maritimeShipments, err := controller.maritimeShipmentService.FindAll(ctx)
	if err != nil {
		return []model.MaritimeShipment{}, err
	}

	return maritimeShipments, nil
}

func (controller *maritimeShipmentController) FindById(ctx *gin.Context) (model.MaritimeShipment, error) {
	var maritimeShipment model.MaritimeShipment

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	maritimeShipment, err = controller.maritimeShipmentService.FindById(id)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}

func (controller *maritimeShipmentController) Update(ctx *gin.Context) (model.MaritimeShipment, error) {
	var maritimeShipment model.MaritimeShipment
	err := ctx.ShouldBind(&maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	err = validate.Struct(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	maritimeShipment.ID = uint(id)
	maritimeShipment, err = controller.maritimeShipmentService.Update(maritimeShipment)
	if err != nil {
		return model.MaritimeShipment{}, err
	}

	return maritimeShipment, nil
}

func (controller *maritimeShipmentController) Delete(ctx *gin.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	if err = controller.maritimeShipmentService.Delete(id); err != nil {
		return err
	}

	return nil
}
