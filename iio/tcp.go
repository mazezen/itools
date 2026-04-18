package iio

type ITcp struct {
	msgHeader string
}

func NewITcp(h string) *ITcp {
	if h == "" {
		h = "123456789"
	}
	return &ITcp{msgHeader: h}
}
