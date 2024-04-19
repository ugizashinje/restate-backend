package ctx

import (
	"fmt"
	"warrant-api/pkg/messages"
	"warrant-api/pkg/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	Email       string = "email"
	Transaction string = "transaction"
	Session     string = "session"
)

type RequestContext struct {
	gin.Context
}

func (rc *RequestContext) GetTransaction() *gorm.DB {

	if transaction, ok := rc.Value(Transaction).(*gorm.DB); ok {
		return transaction
	}
	return nil
}

func GetMandatoryQueryParam(g *gin.Context, param string) string {
	queryParam := g.Request.URL.Query()[param]
	if len(queryParam) == 0 {
		utils.Handle(messages.Errorf(404, fmt.Sprintf("Query param %s is mandatory", param)))
	}
	return queryParam[0]
}
