package services

import (
	"douyinAPI/internal/v1/form"
	model "douyinAPI/internal/v1/models"
	"encoding/json"
	"fmt"
	"github.com/pengcainiao/zero/rest/httprouter"
	"testing"
)


func TestDyServices_GetWord(t *testing.T) {
	var params = form.GetUrlRequest{
		Url: "在双重，通在单轻... https://v.douyin.com/y8mj2uS/",
	}
	var (
		ctx *httprouter.Context
	)
	var data []map[string]interface{}
	resp := model.Dy{}.GetWord(ctx, params)
	_ = json.Unmarshal([]byte(resp), &data)
	var ss []T
	for _, v := range data {
		fmt.Println(v["level"])
		l, _ := v["level"].(float64)
		ti, _ := v["tag_id"].(float64)
		tn, _ := v["tag_name"].(string)

		ss = append(ss, T{
			Level:   l,
			TagId:   ti,
			TagName: tn,
		})
	}
	fmt.Println("ss", ss)
}
