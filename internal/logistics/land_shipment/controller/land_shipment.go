package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/service"
	"github.com/freddiemo/logistics-api/internal/logistics/validators"
)

type LandShipmentController interface {
	Save(ctx *gin.Context) (model.LandShipment, error)
	FindAll(ctx *gin.Context) ([]model.LandShipment, error)
	FindById(ctx *gin.Context) (model.LandShipment, error)
	Update(ctx *gin.Context) (model.LandShipment, error)
	Delete(ctx *gin.Context) error
}

type landShipmentController struct {
	landShipmentService service.LandShipmentServiceInterface
}

var validate *validator.Validate

func NewLandShipmentController(landShipmentService service.LandShipmentServiceInterface) LandShipmentController {
	validate = validator.New()
	validate.RegisterValidation("is-vehicle-plate", validators.IsVehiclePlate)
	validate.RegisterValidation("is-guide-number", validators.IsGuideNumber)
	return &landShipmentController{
		landShipmentService: landShipmentService,
	}
}

func (controller *landShipmentController) Save(ctx *gin.Context) (model.LandShipment, error) {
	var landShipment model.LandShipment
	err := ctx.ShouldBind(&landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	err = validate.Struct(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	landShipment, err = controller.landShipmentService.Save(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (controller *landShipmentController) FindAll(ctx *gin.Context) ([]model.LandShipment, error) {
	landShipments, err := controller.landShipmentService.FindAll()
	if err != nil {
		return []model.LandShipment{}, err
	}

	return landShipments, nil
}

func (controller *landShipmentController) FindById(ctx *gin.Context) (model.LandShipment, error) {
	var landShipment model.LandShipment

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.LandShipment{}, err
	}

	landShipment, err = controller.landShipmentService.FindById(id)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (controller *landShipmentController) Update(ctx *gin.Context) (model.LandShipment, error) {
	var landShipment model.LandShipment
	err := ctx.ShouldBind(&landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.LandShipment{}, err
	}

	err = validate.Struct(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	landShipment.ID = uint(id)
	landShipment, err = controller.landShipmentService.Update(landShipment)
	if err != nil {
		return model.LandShipment{}, err
	}

	return landShipment, nil
}

func (controller *landShipmentController) Delete(ctx *gin.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	if err = controller.landShipmentService.Delete(id); err != nil {
		return err
	}

	return nil
}
