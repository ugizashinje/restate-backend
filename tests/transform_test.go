package tests

import (
	"testing"
	"time"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/services"
	"warrant-api/pkg/utils/transformer"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/guregu/null.v4"
)

func TestUser(t *testing.T) {

	Convey("transform", t, func() {
		apartment := model.Apartment{
			BaseModel: model.BaseModel{
				ID:        "adssadas",
				CreatedAt: time.Now(),
			},
			Name:       "orig",
			Street:     null.StringFrom("Nehruova"),
			Floor:      3,
			Age:        null.IntFrom(1922),
			Registered: false,
		}
		patch := services.PatchApartmentRequest{
			Name:       null.StringFrom("aply"),
			Street:     null.StringFrom("Borcanska"),
			Registered: null.BoolFrom(true),
			Floor:      null.IntFrom(4),
			Age:        null.IntFrom(2002),
		}
		transformer.Patch(&apartment, patch)

		So(apartment.Name, ShouldEqual, "aply")
		So(apartment.Street.String, ShouldEqual, "Borcanska")
		So(apartment.Floor, ShouldEqual, 4)
		So(apartment.Age.Int64, ShouldEqual, 2002)
		So(apartment.Registered, ShouldEqual, true)
	})

}
