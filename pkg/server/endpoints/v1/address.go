package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterAddress(g *gin.RouterGroup) {
	v1 := g.Group("addresses")
	v1.GET("/", SearchAllAddress)
	v1.GET("/:id", getAddress)
	v1.PATCH("/:id", patchAddress)
	v1.POST("/", createAddress)
	v1.DELETE("/:id", deleteAddress)
}

func SearchAllAddress(g *gin.Context) {
	addresses, err := wire.Svc.AddressService.Search(g, nil)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, addresses)
}

func getAddress(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))
	address, _ := wire.Svc.AddressService.GetById(g, id)
	g.PureJSON(200, address)
}

func patchAddress(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchAddressRequest
	bind(g, &patch)

	address, _ := wire.Svc.AddressService.Update(g, id, patch)
	g.PureJSON(200, address)

}

func createAddress(g *gin.Context) {
	var createAddress services.CreateAddressRequest
	bind(g, &createAddress)

	address, err := wire.Svc.AddressService.Create(g, createAddress)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, address)
}

func deleteAddress(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))

	address, err := wire.Svc.AddressService.Delete(g, id)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, address)

}
