package wx

import (
	"github.com/mz-eco/mz/settings"
	"github.com/mz-eco/mz/wx"
)

var (
	wxXcx = &wx.WxXcx{}
)

func init() {
	settings.RegisterWith(func(viper *settings.Viper) error {
		err := viper.Unmarshal(wxXcx)
		if err != nil {
			panic(err)
			return err
		}

		return err
	}, "wx_xcx")
}

func GetWxXcx() *wx.WxXcx {
	return wxXcx
}
