package wx

import (
	"fmt"
	"os"
	"testing"

	"github.com/mz-eco/mz/settings"
)

func TestWxWork_GetAccessToken(t *testing.T) {
	token, err := GetWxWork().GetAccessToken()
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(token)
}

func TestWxWork_GetUser(t *testing.T) {
	user, err := GetWxWork().GetUser("luoh")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(user)
}

func TestMain(m *testing.M) {
	settings.LoadFrom("../../", "")
	os.Exit(m.Run())
}
