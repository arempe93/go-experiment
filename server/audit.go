package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/arempe93/experiment/models"
)

type AuditEntity struct {
	ID     uint   `key:"id"`
	Action string `key:"action"`
}

var getAudit = Endpoint{
	Entity: AuditEntity{},
	URI: struct {
		ID string `uri:"id" key:"id" binding:"required"`
	}{},
	Handler: func(c *gin.Context, db *gorm.DB, params map[string]interface{}) (int, interface{}) {
		var audit models.Audit

		if db.First(&audit, params["id"]).RecordNotFound() {
			return http.StatusNotFound, fmt.Sprintf("Audit with id %v not found", params["id"])
		}

		return http.StatusOK, &audit
	},
}

var createAudit = Endpoint{
	Entity: AuditEntity{},
	Params: struct {
		Action string `form:"action" json:"action" key:"action" binding:"required"`
	}{},
	Handler: func(c *gin.Context, db *gorm.DB, params map[string]interface{}) (int, interface{}) {
		audit := models.Audit{
			Action: params["action"].(string),
		}

		db.Create(&audit)

		return http.StatusCreated, &audit
	},
}

func AuditRoutes(r *gin.RouterGroup) {
	r.POST("/audits", CreateEndpoint(createAudit))
	r.GET("/audits/:id", CreateEndpoint(getAudit))
}
