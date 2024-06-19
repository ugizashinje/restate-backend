package services

import (
	"time"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/repo"
	"warrant-api/pkg/session"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
)

type AdServiceImpl struct {
	AdRepo       repo.Repo[model.Ad]
	LocationRepo repo.Repo[model.Location]
}

type CreateAdRequest struct {
	CompanyID      null.String `json:"-"`
	UserID         null.String `json:"-"` // agent may be working for other company
	Published      time.Time   `json:"published"`
	Status         string      `json:"status"`
	ThumbnailImage string      `json:"tumbnailImage"`
	ThumbnailText  string      `json:"tumbnailText"`
	Description    string      `json:"description"`
}

type PatchAdRequest struct {
	CompanyID      null.String `json:"companyId"`
	UserID         null.String `json:"userId"` // agent may be working for other company
	Status         string      `json:"status"`
	Expiry         time.Time   `json:"expiry"`
	ThumbnailImage string      `json:"tumbnailImage"`
	ThumbnailText  string      `json:"tumbnailText"`
	Description    string      `json:"description"`
}

func (svc *AdServiceImpl) Create(g *gin.Context, createAd CreateAdRequest) (*model.Ad, error) {
	userSession := session.GetSession(g)

	ad := &model.Ad{
		CompanyID:      userSession.User.CompanyID,
		UserID:         userSession.User.ID,
		Published:      createAd.Published,
		ThumbnailImage: createAd.ThumbnailImage,
		ThumbnailText:  createAd.ThumbnailText,
		Description:    createAd.Description,
	}

	result := svc.AdRepo.Create(g, ad)
	utils.Handle(result.Error)
	return ad, nil
}

func (svc *AdServiceImpl) Delete(g *gin.Context, adID string) (ad *model.Ad, err error) {
	ad = &model.Ad{}
	result := svc.AdRepo.Delete(g, ad, adID)
	utils.Handle(result.Error)

	return ad, nil
}

func (svc *AdServiceImpl) GetById(c *gin.Context, adID string) (ad *model.Ad, err error) {
	ad = &model.Ad{}
	result := svc.AdRepo.GetById(c, ad, adID)
	utils.Handle(result.Error)
	return ad, nil
}

func (svc *AdServiceImpl) Update(c *gin.Context, adID string, patch any) (ad *model.Ad, err error) {
	ad = &model.Ad{}
	result := svc.AdRepo.Patch(c, ad, adID, patch)
	utils.Handle(result.Error)

	return ad, nil
}

func (svc *AdServiceImpl) Search(c *gin.Context, query map[string]any) (ads []model.Ad, err error) {
	ads = []model.Ad{}

	result := svc.AdRepo.Search(c, &ads, query)
	utils.Handle(result.Error)
	return ads, nil
}
