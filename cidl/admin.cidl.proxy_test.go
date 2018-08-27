package cidl

import (
	"fmt"
	"os"
	"testing"

	"business/user/proxy/user"

	"github.com/mz-eco/mz/settings"
)

func TestMain(m *testing.M) {
	settings.LoadFrom("../", "")
	os.Exit(m.Run())
}

func TestMakeApiUserInfoByUserId(t *testing.T) {
	u, err := user.NewProxy("user-service").InnerUserInfoByUserID("1")
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(u)
}
