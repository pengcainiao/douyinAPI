package form

type GetUrlRequest struct {
	Url string `json:"url" binding:"required"` // 复制的链接
}
