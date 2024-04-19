package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterShippingInvoice(g *gin.RouterGroup) {
	v1 := g.Group("shippinginvoices")
	v1.GET("/", SearchAllShippingInvoice)
	v1.GET("/:id", getShippingInvoice)
	v1.PATCH("/:id", patchShippingInvoice)
	v1.POST("/", createShippingInvoice)
	v1.DELETE("/:id", deleteShippingInvoice)
}

func SearchAllShippingInvoice(c *gin.Context) {
	shippinginvoicees, err := wire.Svc.ShippingInvoiceService.Search(c, nil)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, shippinginvoicees)
}

func getShippingInvoice(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	shippinginvoice, _ := wire.Svc.ShippingInvoiceService.GetById(c, id)
	c.PureJSON(200, shippinginvoice)
}

func patchShippingInvoice(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchShippingInvoiceRequest
	err := c.BindJSON(&patch)
	utils.Handle(err)

	shippinginvoice, _ := wire.Svc.ShippingInvoiceService.Update(c, id, patch)
	c.PureJSON(200, shippinginvoice)

}

func createShippingInvoice(c *gin.Context) {
	var createShippingInvoice services.CreateShippingInvoiceRequest
	err := c.BindJSON(&createShippingInvoice)
	utils.Handle(err)
	shippinginvoice, err := wire.Svc.ShippingInvoiceService.Create(c, createShippingInvoice)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, shippinginvoice)
}

func deleteShippingInvoice(c *gin.Context) {
	idParam := c.Param("id")
	id := string([]byte(idParam))

	shippinginvoice, err := wire.Svc.ShippingInvoiceService.Delete(c, id)
	if err != nil {
		c.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	c.PureJSON(200, shippinginvoice)

}
