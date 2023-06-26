package services

import (
	"douyinAPI/internal/v1/form"
	model "douyinAPI/internal/v1/models"
	"encoding/json"
	"errors"
	"github.com/pengcainiao/zero/rest/httprouter"
)

type DyServices struct {
}

func (d DyServices) GetUrl(ctx *httprouter.Context, params form.GetUrlRequest) httprouter.Response {
	if len(params.Url) == 0 {
		return httprouter.GetError(httprouter.ErrInvalidParameterCode, errors.New("无效参数"))
	}
	var dy model.Dy
	resp := dy.GetUrl(ctx, params)

	return httprouter.Success(map[string]interface{}{
		"data": resp,
	})
}

type T struct {
	Level   float64       `json:"level"`
	TagId   float64       `json:"tag_id"`
	TagName interface{} `json:"tag_name"`
}

func (d DyServices) GetWord(ctx *httprouter.Context, params form.GetUrlRequest) httprouter.Response {
	if len(params.Url) == 0 {
		return httprouter.GetError(httprouter.ErrInvalidParameterCode, errors.New("无效参数"))
	}
	var dy model.Dy
	resp := dy.GetWord(ctx, params)

	var data []map[string]interface{}
	var ss []T

	_ = json.Unmarshal([]byte(resp), &data)
	for _, v := range data {
		l, _ := v["level"].(float64)
		ti, _ := v["tag_id"].(float64)
		tn, _ := v["tag_name"].(string)

		ss = append(ss, T{
			Level:   l,
			TagId:   ti,
			TagName: tn,
		})
	}

	return httprouter.Success(map[string]interface{}{
		"data": ss,
	})
}
