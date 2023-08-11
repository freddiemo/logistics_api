package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/model"
	"github.com/freddiemo/logistics-api/internal/logistics/land_shipment/service"
	"github.com/freddiemo/logistics-api/internal/logistics/validators"
)

type LandShipmentController interface {
	Save(ctx *gin.Context) (model.LandShipment, error)
	FindAll(ctx *gin.Context) ([]model.LandShipment, error)
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
