package services

import (
	"douyinAPI/internal/v1/form"
	model "douyinAPI/internal/v1/models"
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
	//var data []map[string]interface{}
	resp := model.Dy{}.GetUrl(ctx, params)
	//_ = json.Unmarshal([]byte(resp), &data)
	//var ss []T
	//for _, v := range data {
	//	l, _ := v["level"].(int)
	//	ti, _ := v["tag_id"].(int)
	//	tn, _ := v["tag_name"].(string)
	//
	//	ss = append(ss, T{
	//		Level:   l,
	//		TagId:   ti,
	//		TagName: tn,
	//	})
	//}
	fmt.Println("data", resp)
}
