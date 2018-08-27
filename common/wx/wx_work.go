package wx

import (
	"github.com/mz-eco/mz/settings"
	"github.com/mz-eco/mz/wx"
)

/**
企业微信开发文档：http://work.weixin.qq.com/api/doc
*/

var (
	wxWork = &wx.WxWork{}
)

func init() {
	settings.RegisterWith(func(viper *settings.Viper) error {
		return viper.Unmarshal(wxWork)
	}, "wx_work.wx_work")

}

func GetWxWork() *wx.WxWork {
	return wxWork
}
