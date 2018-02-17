package server

import (
	"net/http"

	"github.com/fnproject/fn/api"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleAppGetByName(c *gin.Context) {
	ctx := c.Request.Context()

	err := s.FireBeforeAppGet(ctx, c.MustGet(api.AppID).(string))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	app, err := s.datastore.GetAppByID(ctx, c.MustGet(api.AppID).(string))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	err = s.FireAfterAppGet(ctx, app)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, appResponse{"Successfully loaded app", app})
}

func (s *Server) handleAppGetByID(c *gin.Context) {
	ctx := c.Request.Context()

	app, err := s.datastore.GetAppByID(ctx, c.Param(api.CApp))
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, appResponse{"Successfully loaded app", app})
}
