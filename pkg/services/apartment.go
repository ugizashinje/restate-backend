package services

import (
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

type ApartmentServiceImpl struct {
	ApartmentRepo repo.Repo[model.Apartment]
}

type CreateApartmentRequest struct {
	Name      string      `json:"name"`
	Street    null.String `json:"street"`
	StreetNo  null.String `json:"streetNo"`
	CompanyID null.String `json:"companyId"`
	UserID    string      `json:"userId"`

	// Location
	City     string      `json:"city"`
	Muni     null.String `json:"muni"`
	Location null.String `json:"location"`

	// Building
	Age             null.Int `json:"age"`
	TotalFloors     int      `json:"totalFloors"`
	Floor           int      `json:"floors"`
	Basement        bool     `json:"basement"`
	GroundFloor     bool     `json:"groundFloor"`
	HighGroundFloor bool     `json:"highGroundFloor"`
	LastFloor       bool     `json:"lastFloor"`
	MonthlyFees     int      `json:"monthlyFees"`

	// Structure
	Sallon    bool   `json:"sallon"`
	Duplex    bool   `json:"duplex"`
	Penthouse bool   `json:"penthouse"`
	EnergyEff bool   `json:"energyEff"`
	Condition string `json:"condition"`
	Heating   string `json:"heating"`
	Rooms     string `json:"rooms"`

	// Legal
	Registered    bool   `json:"registered"`
	Occupiable    bool   `json:"occupiable"`
	WithinHouse   bool   `json:"withinHouse"`
	TaxRefundable bool   `json:"taxRefundable"`
	Swappable     bool   `json:"swappable"`
	Morgage       bool   `json:"morgage"`
	Urgent        bool   `json:"urgent"`
	Seller        string `json:"seller"` // agency , owner, investor
	// MISC
	Tags         pq.StringArray     `json:"tags" gorm:"type:text[]"`
	Thumbnail    string             `json:"thumbnail"`
	FloorPlan    model.FloorPlan    `json:"floorPlan"`
	AllRoundTour model.AllRoundTour `json:"allRoundTour"`
	VirtualTour  string             `json:"virutalTour"`
}

type PatchApartmentRequest struct {
	Name      null.String `json:"name"`
	IsActive  null.Bool   `json:"isActive"`
	Street    null.String `json:"street"`
	StreetNo  null.String `json:"streetNo"`
	CompanyID null.String `json:"companyId"`
	UserID    null.String `json:"userId"`

	// Location
	City     null.String `json:"city"`
	Muni     null.String `json:"muni"`
	Location null.String `json:"location"`

	// Building
	Age             null.Int  `json:"age"`
	TotalFloors     null.Int  `json:"totalFloors"`
	Floor           null.Int  `json:"floors"`
	Basement        null.Bool `json:"basement"`
	GroundFloor     null.Bool `json:"groundFloor"`
	HighGroundFloor null.Bool `json:"highGroundFloor"`
	LastFloor       null.Bool `json:"lastFloor"`
	MonthlyFees     null.Int  `json:"monthlyFees"`

	// Structure
	Sallon    null.Bool   `json:"sallon"`
	Duplex    null.Bool   `json:"duplex"`
	Penthouse null.Bool   `json:"penthouse"`
	EnergyEff null.Bool   `json:"energyEff"`
	Condition null.String `json:"condition"`
	Heating   null.String `json:"heating"`
	Rooms     null.String `json:"rooms"`

	// Legal
	Registered    null.Bool   `json:"registered"`
	Occupiable    null.Bool   `json:"occupiable"`
	WithinHouse   null.Bool   `json:"withinHouse"`
	TaxRefundable null.Bool   `json:"taxRefundable"`
	Swappable     null.Bool   `json:"swappable"`
	Morgage       null.Bool   `json:"morgage"`
	Urgent        null.Bool   `json:"urgent"`
	Seller        null.String `json:"seller"` // agency , owner, investor
	// MISC
	AddTags    pq.StringArray `json:"addTags"`
	RemoveTags pq.StringArray `json:"removeTags"`

	Thumbnail    null.String         `json:"thumbnail"`
	FloorPlan    *model.FloorPlan    `json:"floorPlan"`
	AllRoundTour *model.AllRoundTour `json:"allRoundTour"`
	VirtualTour  null.String         `json:"virutalTour"`
}

func (svc *ApartmentServiceImpl) Create(g *gin.Context, createApartment CreateApartmentRequest) (*model.Apartment, error) {
	apartment := &model.Apartment{
		Ad: model.Ad{
			CompanyID: createApartment.CompanyID,
			UserID:    createApartment.UserID,
		},
		Name:         createApartment.Name,
		Condition:    createApartment.Condition,
		Heating:      createApartment.Heating,
		Street:       createApartment.Street,
		StreetNo:     createApartment.StreetNo,
		City:         createApartment.City,
		Muni:         createApartment.Muni,
		Location:     createApartment.Location,
		Age:          createApartment.Age,
		Rooms:        createApartment.Rooms,
		TotalFloors:  createApartment.TotalFloors,
		Basement:     createApartment.Basement,
		GroundFloor:  createApartment.GroundFloor,
		Seller:       createApartment.Seller,
		Thumbnail:    createApartment.Thumbnail,
		Tags:         createApartment.Tags,
		FloorPlan:    createApartment.FloorPlan,
		AllRoundTour: createApartment.AllRoundTour,
	}

	result := svc.ApartmentRepo.Create(g, apartment)
	utils.Handle(result.Error)
	return apartment, nil
}

func (svc *ApartmentServiceImpl) Delete(g *gin.Context, apartmentID string) (apartment *model.Apartment, err error) {
	apartment = &model.Apartment{}
	result := svc.ApartmentRepo.Delete(g, apartment, apartmentID)
	utils.Handle(result.Error)

	return apartment, nil
}

func (svc *ApartmentServiceImpl) GetById(c *gin.Context, apartmentID string) (apartment *model.Apartment, err error) {
	apartment = &model.Apartment{}
	result := svc.ApartmentRepo.GetById(c, apartment, apartmentID)
	utils.Handle(result.Error)
	return apartment, nil
}

func (svc *ApartmentServiceImpl) Update(g *gin.Context, apartmentID string, patch PatchApartmentRequest) (apartment *model.Apartment, err error) {
	apartment = &model.Apartment{}
	dbRes := svc.ApartmentRepo.GetById(g, apartment, apartmentID)
	utils.Handle(dbRes.Error)
	if patch.IsActive.Valid {
		apartment.BaseModel.IsActive = patch.IsActive.Bool
	}
	if patch.Name.Valid {
		apartment.Name = patch.Name.String
	}
	if patch.Street.Valid {
		apartment.Street = patch.Street
	}
	if patch.StreetNo.Valid {
		apartment.StreetNo = patch.StreetNo
	}
	if patch.City.Valid {
		apartment.City = patch.City.String
	}
	if patch.Muni.Valid {
		apartment.Muni = patch.Muni
	}
	if patch.Location.Valid {
		apartment.Location = patch.Location
	}
	if patch.Age.Valid {
		apartment.Age = patch.Age
	}
	if patch.TotalFloors.Valid {
		apartment.TotalFloors = int(patch.TotalFloors.Int64)
	}
	if patch.Floor.Valid {
		apartment.Floor = int(patch.Floor.Int64)
	}
	if patch.GroundFloor.Valid {
		apartment.Basement = patch.Basement.Bool
	}
	if patch.HighGroundFloor.Valid {
		apartment.HighGroundFloor = patch.HighGroundFloor.Bool
	}
	if patch.LastFloor.Valid {
		apartment.LastFloor = patch.LastFloor.Bool
	}
	if patch.MonthlyFees.Valid {
		apartment.MonthlyFees = int(patch.MonthlyFees.Int64)
	}
	if patch.Registered.Valid {
		apartment.Registered = patch.Registered.Bool
	}

	if patch.Occupiable.Valid {
		apartment.Occupiable = patch.Occupiable.Bool
	}
	if patch.WithinHouse.Valid {
		apartment.WithinHouse = patch.WithinHouse.Bool
	}
	if patch.TaxRefundable.Valid {
		apartment.TaxRefundable = patch.TaxRefundable.Bool
	}

	if patch.Swappable.Valid {
		apartment.Swappable = patch.Swappable.Bool
	}

	if patch.Morgage.Valid {
		apartment.Morgage = patch.Morgage.Bool
	}

	if patch.Urgent.Valid {
		apartment.Urgent = patch.Urgent.Bool
	}
	if patch.Seller.Valid {
		apartment.Seller = patch.Seller.String
	}

	apartment.Tags = append(apartment.Tags, patch.AddTags...)
	apartment.Tags = utils.StringArrayDiff(apartment.Tags, patch.RemoveTags)

	if patch.Thumbnail.Valid {
		apartment.Thumbnail = patch.Thumbnail.String
	}
	if patch.FloorPlan != nil {
		apartment.FloorPlan = *patch.FloorPlan
	}
	if patch.AllRoundTour != nil {
		apartment.AllRoundTour = *patch.AllRoundTour
	}
	if patch.VirtualTour.Valid {
		apartment.VirtualTour = patch.VirtualTour
	}
	result := svc.ApartmentRepo.Save(g, apartment)
	utils.Handle(result.Error)
	return apartment, nil
}

func (svc *ApartmentServiceImpl) Search(c *gin.Context, query map[string]any) (apartments []model.Apartment, err error) {
	apartments = []model.Apartment{}
	result := svc.ApartmentRepo.Search(c, &apartments, query)
	utils.Handle(result.Error)
	return apartments, nil
}
