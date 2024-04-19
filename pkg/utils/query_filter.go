package utils

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"unicode"
	"warrant-api/pkg/db/model"
	"warrant-api/pkg/messages"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type sqlType string

var skip = map[string]bool{
	"page":      true,
	"pageSize":  true,
	"orderBy":   true,
	"BaseModel": true,
}
var modelCache map[string](map[string]string) = make(map[string]map[string]string)
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

var AllowedFilteringStructs = map[string]bool{
	"StringArray": true,
}

func getStructSelectableFields[T model.BaseInterface](t *[]T) map[string]string {
	result := map[string]string{
		"id":        "id",
		"createdAt": "created_at",
		"updatedAt": "updated_at",
		"deletedAt": "deleted_at",
		"isDeleted": "is_deleted",
		"isActive":  "is_active",
	}
	e := reflect.TypeOf(t).Elem()
	elem := e.Elem()
	if entry := modelCache[elem.Name()]; entry != nil {
		return entry
	}

	for i := 0; i < elem.NumField(); i++ {
		key := elem.Field(i).Name

		input := []rune(key)
		firstChar := input[0]
		keyCopy := []rune{unicode.ToLower(rune(firstChar))}

		// Transform struct fields to json camel case
		for i := 1; i < len(input); i++ {
			if unicode.IsLower(input[i-1]) && unicode.IsUpper(input[i]) {
				keyCopy = append(keyCopy, input[i])
			} else {
				keyCopy = append(keyCopy, unicode.ToLower(input[i]))
			}
		}

		if skip[string(keyCopy)] {
			continue
		}

		// Create mapping for json camel case to db snake_case
		fieldType := elem.Field(i).Type.Name()
		kind := elem.Field(i).Type.Kind()
		r := []rune(fieldType)
		if len(r) == 0 || (kind == reflect.Struct && !AllowedFilteringStructs[fieldType]) {
			continue
		}
		snake := matchFirstCap.ReplaceAllString(key, "${1}_${2}")
		snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
		snake = strings.ToLower(snake)

		if elem.Field(i).Type.Kind() == reflect.Slice {
			v := reflect.SliceOf(elem.Field(i).Type).String()
			if v == "[]pq.StringArray" {

				result[string(keyCopy)] = fmt.Sprintf("%s @> string_to_array(? , ',') ", snake)

			}
		} else {

			// string and integer types
			result[string(keyCopy)] = snake
		}

	}
	modelCache[elem.Name()] = result
	return result

}
func QueryFilter[T model.BaseInterface](g *gin.Context, db *gorm.DB, t *[]T) map[string]string {
	result := make(map[string]string)
	sliceType := reflect.TypeOf(t)
	ptrType := sliceType.Elem().Elem()

	zeroType := reflect.New(ptrType)
	zeroValue := reflect.New(zeroType.Type().Elem())
	b := zeroValue.Interface().(model.BaseInterface)
	resource := b.ResouceName()
	path := strings.Split(g.Request.URL.Path, "/")
	if len(path) < 2 {
		Handle(messages.Errorf(http.StatusBadRequest, "Invalid resrouce"))
	}

	if resource != path[2] {
		return result
	}
	query := g.Request.URL.Query()
	modelFields := getStructSelectableFields(t)

	for k, arr := range query {
		if _, ok := modelFields[k]; !ok {
			continue
		}
		dbColumn := modelFields[k]
		if !skip[dbColumn] {
			result[dbColumn] = arr[len(arr)-1]
		}
	}

	return result
}
