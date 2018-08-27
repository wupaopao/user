package wx

import (
	"fmt"
	"os"
	"testing"

	"github.com/mz-eco/mz/settings"
)

func TestMain(m *testing.M) {
	settings.LoadFrom("../../", "")
	os.Exit(m.Run())
}

func TestWxXcx_GetSessionKey(t *testing.T) {
	sessionKey, err := GetWxXcx().GetSessionKey("013Swlol1CpfBl0WQ7pl1Pqtol1Swlor")
	fmt.Println(sessionKey)
	if err != nil {
		t.Error(err)
		return
	}

}
