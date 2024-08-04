package zeewImg

import (
	"bytes"
	"net/url"

	ZeewUtils "github.com/Oscar-Dev0/zeew-api-go/utils"
)

type Filter struct {
	Token string
	URI   string
}

func NewFilter(token string) *Filter {
	return &Filter{
		Token: token,
	}
}

func (f *Filter) Triggered(img string) (*ZeewUtils.ZeewError, *bytes.Buffer) {
	if img == "" {
		return ZeewUtils.NewZeewError(ZeewUtils.InvalidImage), nil
	};

	query := url.Values{}
	query.Add("avatar", img);
	f.URI = f.URI + "/triggered?" + query.Encode();

	res, err := ZeewUtils.Request(f.Token, f.URI);
	if err != nil {
		return ZeewUtils.NewZeewError(ZeewUtils.RequestError), nil
	}
	return nil, res
}
