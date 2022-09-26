package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"maxblog-me-main/internal/core"
	"maxblog-me-main/src/service"
	"maxblog-me-main/src/utils"
	"net/http"
)

var DemoSet = wire.NewSet(wire.Struct(new(HDemo), "*"))

type HDemo struct {
	BDemo *service.BDemo
	IRes  core.IResponse
}

func (hDemo *HDemo) GetDemos(c *gin.Context) {
	count := c.Query("count")
	numCount, err := utils.Str2Uint32(count)
	if err != nil {
		hDemo.IRes.ResFailure(c, core.GetFuncName(), http.StatusBadRequest, core.FormatError(201, err))
		return
	}
	dataRes, err := hDemo.BDemo.GetDemos(c, numCount)
	if err != nil {
		hDemo.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(399, err))
		return
	}
	hDemo.IRes.ResSuccess(c, core.GetFuncName(), dataRes)
}

func (hDemo *HDemo) GetDemoById(c *gin.Context) {
	idRaw := c.Param("id")
	id, err := utils.Str2Uint32(idRaw)
	if err != nil {
		hDemo.IRes.ResFailure(c, core.GetFuncName(), http.StatusBadRequest, core.FormatError(201, err))
		return
	}
	dataRes, err := hDemo.BDemo.GetDemoById(c, id)
	if err != nil {
		hDemo.IRes.ResFailure(c, core.GetFuncName(), http.StatusInternalServerError, core.FormatError(399, err))
		return
	}
	hDemo.IRes.ResSuccess(c, core.GetFuncName(), dataRes)
}
