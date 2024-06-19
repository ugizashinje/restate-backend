package v1

import (
	"warrant-api/pkg/services"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterApartment(g *gin.RouterGroup) {
	v1 := g.Group("apartments")
	v1.GET("/", SearchAllApartment)
	v1.GET("/:id", getApartment)
	v1.PATCH("/:id", patchApartment)
	v1.POST("/", createApartment)
	v1.DELETE("/:id", deleteApartment)
}

func SearchAllApartment(g *gin.Context) {
	apartmentes, err := wire.Svc.ApartmentService.Search(g, nil)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, apartmentes)
}

func getApartment(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))
	apartment, _ := wire.Svc.ApartmentService.GetById(g, id)
	g.PureJSON(200, apartment)
}

func patchApartment(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))
	var patch services.PatchApartmentRequest
	bind(g, &patch)

	apartment, _ := wire.Svc.ApartmentService.Update(g, id, patch)
	g.PureJSON(200, apartment)

}

func createApartment(g *gin.Context) {
	var createApartment services.CreateApartmentRequest
	bind(g, &createApartment)

	apartment, err := wire.Svc.ApartmentService.Create(g, createApartment)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, apartment)
}

func deleteApartment(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))

	apartment, err := wire.Svc.ApartmentService.Delete(g, id)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, apartment)

}
