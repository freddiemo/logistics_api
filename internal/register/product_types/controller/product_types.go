package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/freddiemo/logistics-api/internal/register/product_types/model"
	"github.com/freddiemo/logistics-api/internal/register/product_types/service"
)

type ProductTypeController interface {
	Save(ctx *gin.Context) (model.ProductType, error)
	FindAll(ctx *gin.Context) ([]model.ProductType, error)
	FindById(ctx *gin.Context) (model.ProductType, error)
	Update(ctx *gin.Context) (model.ProductType, error)
}

type productTypeController struct {
	productTypeService service.ProductTypeServiceInterface
}

func NewProductTypeController(productTypeService service.ProductTypeServiceInterface) ProductTypeController {
	return &productTypeController{
		productTypeService: productTypeService,
	}
}

func (controller *productTypeController) Save(ctx *gin.Context) (model.ProductType, error) {
	var productType model.ProductType
	err := ctx.ShouldBind(&productType)
	if err != nil {
		return model.ProductType{}, err
	}

	productType, err = controller.productTypeService.Save(productType)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}

func (controller *productTypeController) FindAll(ctx *gin.Context) ([]model.ProductType, error) {
	productTypes, err := controller.productTypeService.FindAll()
	if err != nil {
		return []model.ProductType{}, err
	}

	return productTypes, nil
}

func (controller *productTypeController) FindById(ctx *gin.Context) (model.ProductType, error) {
	var productType model.ProductType

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.ProductType{}, err
	}

	productType, err = controller.productTypeService.FindById(id)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}

func (controller *productTypeController) Update(ctx *gin.Context) (model.ProductType, error) {
	var productType model.ProductType
	err := ctx.ShouldBind(&productType)
	if err != nil {
		return model.ProductType{}, err
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.ProductType{}, err
	}

	productType.ID = uint(id)
	productType, err = controller.productTypeService.Update(productType)
	if err != nil {
		return model.ProductType{}, err
	}

	return productType, nil
}
