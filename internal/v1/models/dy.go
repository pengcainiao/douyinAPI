package douyin

import (
	"douyinAPI/internal/v1/form"
	"fmt"
	"github.com/pengcainiao/zero/core/logx"
	"github.com/pengcainiao/zero/rest/httprouter"
	"github.com/tidwall/gjson"
)

type Dy struct {
}

func (d Dy) GetUrl(ctx *httprouter.Context, params form.GetUrlRequest) string {
	var (
		data    string
		dy      = NewDouYin()
		content = params.Url
	)
	urlStr := dy.Pattern.FindString(content)
	videoId, err := dy.ParseShareUrl(urlStr)
	if err != nil {
		logx.NewTraceLogger(ctx).Err(err).Msg("解析URL失败")
		return data
	}
	rawUrlStr, err := dy.GetDetailUrlByVideoId(videoId)
	if err != nil {
		logx.NewTraceLogger(ctx).Err(err).Msg("查询详情失败")
		return data
	}
	fmt.Println("rawUrlStr",rawUrlStr)
	fmt.Println("dy",dy)
	b, err := dy.GetVideoInfo(rawUrlStr)

	fmt.Println(b)
	data = gjson.Get(b, "aweme_detail.video.play_addr.url_list.0").String()
	return data
}
func (d Dy) GetWord(ctx *httprouter.Context, params form.GetUrlRequest) string {
	var (
		data    string
		dy      = NewDouYin()
		content = params.Url
	)
	fmt.Println(content)
	urlStr := dy.Pattern.FindString(content)
	fmt.Println(urlStr)
	videoId, err := dy.ParseShareUrl(urlStr)
	fmt.Println(videoId)

	if err != nil {
		fmt.Println(111)
		logx.NewTraceLogger(ctx).Err(err).Msg("解析URL失败")
		return data
	}
	rawUrlStr, err := dy.GetDetailUrlByVideoId(videoId)
	fmt.Println(rawUrlStr)
	if err != nil {
		fmt.Println(222)
		logx.NewTraceLogger(ctx).Err(err).Msg("查询详情失败")
		return data
	}
	b, err := dy.GetVideoInfo(rawUrlStr)
	if err != nil {
		fmt.Println(333)
		logx.NewTraceLogger(ctx).Err(err).Msg("查询详情失败")
		return data
	}

	fmt.Println("b",b)
	target:=gjson.Get(b,"aweme_detail.video_tag").String()
	fmt.Println("target",target)

	data=target
	fmt.Println("data",data)
	return data
}
