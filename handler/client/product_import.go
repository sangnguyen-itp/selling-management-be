package client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/tealeg/xlsx/v3"
	"selling-management-be/context"
	"selling-management-be/defined/domain"
	"selling-management-be/helper/excel"
	"selling-management-be/pkg/app"
	"selling-management-be/pkg/logger"
	"selling-management-be/service"
)

// ProductImport docs
// @Summary      ProductImport
// @Description  /api/v1/client/product/get
// @Tags         Client
// @Accept       json
// @Produce      json
// @Success      201
// @Security     ApiKeyAuth
// @Router       /v1/client/product/import [post]
func ProductImport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, err := ctx.FormFile("file")
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "load file", err)
			app.Response(ctx, 400, "file not exist", nil)
			return
		}

		_, ok := excel.ValidateExt(file.Filename)
		if !ok {
			logger.Log().Error(domain.ProductDomain, "ext file", err)
			app.Response(ctx, 400, "file's ext is not support. Please use (.xlsx, .xls)", nil)
			return
		}

		cFile, err := file.Open()
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "open file", err)
			app.Response(ctx, 400, "open file failed", nil)
			return
		}
		defer cFile.Close()

		f, err := xlsx.OpenReaderAt(cFile, file.Size)
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "open file", err)
			app.Response(ctx, 400, "open file failed", nil)
			return
		}
		actor := context.NewBase(ctx)
		var request []*service.ProductCreateRequest
		defaultSheet := f.Sheets[0]
		ignoreFirst := true
		err = defaultSheet.ForEachRow(func(r *xlsx.Row) error {
			if !ignoreFirst {
				tpe := r.GetCell(0).String()
				name := r.GetCell(1).String()
				code := r.GetCell(2).String()
				wholesalePrice := r.GetCell(3).String()
				retailPrice := r.GetCell(4).String()
				wholesaleUnit := r.GetCell(5).String()
				retailUnit := r.GetCell(6).String()
				currency := r.GetCell(7).String()

				if len(tpe) == 0 && len(name) == 0 && len(code) == 0 {
					return nil
				}

				var wholesalePriceDecimal, retailPriceDecimal = decimal.NewFromInt(0), decimal.NewFromInt(0)
				if len(wholesalePrice) > 0 {
					wholesalePriceDecimal, err = decimal.NewFromString(wholesalePrice)
					if err != nil {
						logger.Log().Error(domain.ProductDomain, "file line is invalid", err)
						app.Response(ctx, 400, fmt.Sprintf("err: line"), nil)
						return err
					}
				}

				if len(retailPrice) > 0 {
					retailPriceDecimal, err = decimal.NewFromString(retailPrice)
					if err != nil {
						logger.Log().Error(domain.ProductDomain, "file line is invalid", err)
						app.Response(ctx, 400, fmt.Sprintf("err: line"), nil)
						return err
					}
				}

				request = append(request, &service.ProductCreateRequest{
					Type:           tpe,
					Name:           name,
					Code:           code,
					WholesalePrice: wholesalePriceDecimal,
					RetailPrice:    retailPriceDecimal,
					WholesaleUnit:  wholesaleUnit,
					RetailUnit:     retailUnit,
					Currency:       currency,
					OrganizationID: actor.OrganizationID,
					UpdatedBy:      actor.UserID,
					UpdatedAt:      actor.UpdateTime,
					CreatedBy:      actor.UserID,
					CreatedAt:      actor.UpdateTime,
				})
			}
			ignoreFirst = false
			return nil
		})
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "open file", err)
			app.Response(ctx, 400, "open file failed", nil)
			return
		}

		_, err = service.ProductBatchPut(&service.ProductBatchInsertRequest{
			Products: request,
		})
		if err != nil {
			logger.Log().Error(domain.ProductDomain, "import failed", err)
			app.Response(ctx, 400, "import failed", nil)
			return
		}

		app.Response(ctx, 201, "import successfully", nil)
	}
}
