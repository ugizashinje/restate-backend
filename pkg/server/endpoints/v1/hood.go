package v1

import (
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func RegisterLocations(g *gin.RouterGroup) {
	v1 := g.Group("locations")
	v1.GET("/", SearchAllLocation)
	v1.GET("/:id", getLocation)
	// v1.PATCH("/:id", patchLocation)
	// v1.POST("/", createLocation)
	// v1.DELETE("/:id", deleteLocation)
}

func SearchAllLocation(g *gin.Context) {
	locationes, err := wire.Svc.LocationService.Search(g, nil)
	if err != nil {
		g.PureJSON(500, struct{ error string }{error: err.Error()})
	}
	g.PureJSON(200, locationes)
}

func getLocation(g *gin.Context) {
	idParam := g.Param("id")
	id := string([]byte(idParam))
	location, _ := wire.Svc.LocationService.GetById(g, id)
	g.PureJSON(200, location)
}

// func patchLocation(g *gin.Context) {
// 	idParam := g.Param("id")
// 	id := string([]byte(idParam))
// 	var patch services.PatchLocationRequest
// 	bind(g, &patch)

// 	location, _ := wire.Svc.LocationService.Update(g, id, patch)
// 	g.PureJSON(200, location)

// }

// func createLocation(g *gin.Context) {
// 	var createLocation services.CreateLocationRequest
// 	bind(g, &createLocation)

// 	location, err := wire.Svc.LocationService.Create(g, createLocation)
// 	if err != nil {
// 		g.PureJSON(500, struct{ error string }{error: err.Error()})
// 	}
// 	g.PureJSON(200, location)
// }

// func deleteLocation(g *gin.Context) {
// 	idParam := g.Param("id")
// 	id := string([]byte(idParam))

// 	location, err := wire.Svc.LocationService.Delete(g, id)
// 	if err != nil {
// 		g.PureJSON(500, struct{ error string }{error: err.Error()})
// 	}
// 	g.PureJSON(200, location)

// }
