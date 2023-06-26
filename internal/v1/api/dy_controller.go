package api

import (
	"douyinAPI/internal/v1/form"
	"douyinAPI/internal/v1/services"
	"github.com/gin-gonic/gin"
	"github.com/pengcainiao/zero/rest/httprouter"
	"github.com/pengcainiao/zero/rest/validator"
)

type DyController struct {
	service services.DyServices
}

// GetUrl 换取url
// @Summary 换取url
// @Tags 换取url（任务）
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param   body      body      form.GetUrlRequest        true    "json数据"
// @Success 200 {object} httprouter.Response
// @Failure 400 {object} httprouter.Response
// @Router /v1/url [post]
func (o DyController) GetUrl(c *gin.Context) {
	var params = form.GetUrlRequest{}
	if err := c.ShouldBind(&params); err != nil {
		httprouter.ResponseJSONContent(c, httprouter.ErrorSame(
			httprouter.ErrInvalidParameterCode,
			validator.NewValidator().Translate(err),
		))
		return
	}
	res := o.service.GetUrl(httprouter.NewContext(c), params)
	httprouter.ResponseJSONContent(c, res)
}

// GetWord 查询关键词
// @Summary 查询关键词
// @Tags 查询关键词（任务）
// @Security ApiKeyAuth
// @accept json
// @Produce json
// @Param   body      body      form.GetUrlRequest        true    "json数据"
// @Success 200 {object} httprouter.Response
// @Failure 400 {object} httprouter.Response
// @Router /v1/word [post]
func (o DyController) GetWord(c *gin.Context) {
	var params = form.GetUrlRequest{}
	if err := c.ShouldBind(&params); err != nil {
		httprouter.ResponseJSONContent(c, httprouter.ErrorSame(
			httprouter.ErrInvalidParameterCode,
			validator.NewValidator().Translate(err),
		))
		return
	}
	res := o.service.GetWord(httprouter.NewContext(c), params)
	httprouter.ResponseJSONContent(c, res)
}
