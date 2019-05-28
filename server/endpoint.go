package server

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"github.com/arempe93/experiment/database"
)

type HandlerFunc func(c *gin.Context, db *gorm.DB, params map[string]interface{}) (int, interface{})

type Endpoint struct {
	Entity  interface{}
	Handler HandlerFunc
	Params  interface{}
	URI     interface{}
}

func CreateEndpoint(endpoint Endpoint) gin.HandlerFunc {
	db := database.Instance()

	return func(c *gin.Context) {
		var params, uri interface{}
		var paramsMap = make(map[string]interface{})
		var uriMap = make(map[string]interface{})

		if endpoint.Params != nil {
			params = reflect.New(reflect.TypeOf(endpoint.Params)).Interface()

			if err := c.ShouldBind(params); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			paramsMap = structToMap(params)
		}

		if endpoint.URI != nil {
			uri = reflect.New(reflect.TypeOf(endpoint.URI)).Interface()

			if err := c.ShouldBindUri(uri); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
				return
			}

			uriMap = structToMap(uri)
		}

		for k, v := range uriMap {
			paramsMap[k] = v
		}

		fmt.Println(params, uri, paramsMap)

		status, value := endpoint.Handler(c, db, paramsMap)

		if status == 0 {
			return
		}

		if value == nil {
			c.JSON(status, gin.H{})
		} else if endpoint.Entity == nil {
			c.JSON(status, value)
		} else if m, ok := value.(gin.H); ok {
			c.JSON(status, m)
		} else if m, ok := value.(string); ok {
			c.JSON(status, gin.H{"message": m})
		} else {
			entity := reflect.New(reflect.TypeOf(endpoint.Entity)).Interface()
			c.JSON(status, BuildResponse(value, entity))
		}
	}
}

func BuildResponse(val, ref interface{}) map[string]interface{} {
	copier.Copy(ref, val)
	return structToMap(ref)
}

func structToMap(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("Interface is not a struct!")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := t.Field(i)

		if tag := f.Tag.Get("key"); tag != "" {
			if f.Type.Kind() == reflect.Struct {
				out[tag] = structToMap(v.Field(i).Interface())
			} else {
				out[tag] = v.Field(i).Interface()
			}
		}
	}

	return out
}
