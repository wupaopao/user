package db

import (
	"fmt"
	"os"
	"testing"

	"business/user/cidl"

	"github.com/mz-eco/mz/settings"
)

func TestMain(m *testing.M) {
	settings.LoadFrom("../../", "")
	os.Exit(m.Run())
}

func TestMallUser_AddUser(t *testing.T) {
	_, err := NewMallUser().AddUser(&cidl.User{
		UserID:              "1",
		WxUnionId:           "123456",
		Nickname:            "张三的歌",
		Avatar:              "http://www.qqzhi.com/uploadpic/2014-09-24/110416121.jpg",
		Name:                "张三",
		Mobile:              "18676726608",
		IdCardNumber:        "440981199001011234",
		IdCardFront:         "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1515655661&di=7e3282b38f397402859c7ac632113b49&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.qqzhi.com%2Fwenwen%2Fuploads%2Fservice.qq.com%2Fuserfiles%2FImage%2Fzakardchen%2Fmohu.jpg",
		IdCardBack:          "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2445163217,3569169920&fm=27&gp=0.jpg",
		IsOrgManager:        true,
		IsOrgStaff:          false,
		IsCmtManager:        false,
		IsDisableOrgManager: false,
		IsDisableOrgStaff:   false,
		IsDisableCmtManger:  false,
	})
	if err != nil {
		t.Error(err)
		return
	}

	_, err = NewMallUser().AddUser(&cidl.User{
		UserID:              "2",
		WxUnionId:           "1234567",
		Nickname:            "团长-小花",
		Avatar:              "http://www.qqzhi.com/uploadpic/2014-09-24/110416121.jpg",
		Name:                "花儿",
		Mobile:              "18676726609",
		IdCardNumber:        "440981199001011234",
		IdCardFront:         "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1515655661&di=7e3282b38f397402859c7ac632113b49&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.qqzhi.com%2Fwenwen%2Fuploads%2Fservice.qq.com%2Fuserfiles%2FImage%2Fzakardchen%2Fmohu.jpg",
		IdCardBack:          "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2445163217,3569169920&fm=27&gp=0.jpg",
		IsOrgManager:        false,
		IsOrgStaff:          false,
		IsCmtManager:        true,
		IsDisableOrgManager: false,
		IsDisableOrgStaff:   false,
		IsDisableCmtManger:  false,
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestMallUser_AddUser_OrgManager(t *testing.T) {
	malUser := NewMallUser()
	for i := uint32(1); i < 100; i++ {
		id := i
		user := &cidl.User{
			WxUnionId:    "1234567" + fmt.Sprintf("%d", i),
			Nickname:     fmt.Sprintf("张三的歌-%d", id),
			Avatar:       "http://www.qqzhi.com/uploadpic/2014-09-24/110416121.jpg",
			Name:         fmt.Sprintf("张三-%d", id),
			Mobile:       "186767266" + fmt.Sprintf("%d", i),
			IdCardNumber: "440981199001011234",
			IdCardFront:  "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1515655661&di=7e3282b38f397402859c7ac632113b49&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.qqzhi.com%2Fwenwen%2Fuploads%2Fservice.qq.com%2Fuserfiles%2FImage%2Fzakardchen%2Fmohu.jpg",
			IdCardBack:   "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2445163217,3569169920&fm=27&gp=0.jpg",
			IsOrgManager: true,
			IsOrgStaff:   false,
			IsCmtManager: false,
		}
		_, err := malUser.AddUser(user)
		if err != nil {
			t.Error(err)
			return
		}
	}

}

func TestMallUser_AddUser_CmtManager(t *testing.T) {
	malUser := NewMallUser()
	for i := uint32(1); i <= 1; i++ {
		id := i
		user := &cidl.User{
			WxUnionId:    "103456748" + fmt.Sprintf("%d", i),
			Nickname:     fmt.Sprintf("李四-%d", id),
			Avatar:       "http://www.qqzhi.com/uploadpic/2014-09-24/110416121.jpg",
			Name:         fmt.Sprintf("李四-%d", id),
			Mobile:       "106767266" + fmt.Sprintf("%d", i),
			IdCardNumber: "440981199001011234",
			IdCardFront:  "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1515655661&di=7e3282b38f397402859c7ac632113b49&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.qqzhi.com%2Fwenwen%2Fuploads%2Fservice.qq.com%2Fuserfiles%2FImage%2Fzakardchen%2Fmohu.jpg",
			IdCardBack:   "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2445163217,3569169920&fm=27&gp=0.jpg",
			IsOrgManager: false,
			IsOrgStaff:   false,
			IsCmtManager: true,
		}
		_, err := malUser.AddUser(user)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestMallUser_AddUser_OrgStaff(t *testing.T) {
	malUser := NewMallUser()
	for i := uint32(1); i <= 100; i++ {
		id := i
		user := &cidl.User{
			WxUnionId:    "143456748" + fmt.Sprintf("%d", i),
			Nickname:     fmt.Sprintf("王五-%d", id),
			Avatar:       "http://www.qqzhi.com/uploadpic/2014-09-24/110416121.jpg",
			Name:         fmt.Sprintf("王五-%d", id),
			Mobile:       "166767266" + fmt.Sprintf("%d", i),
			IdCardNumber: "440981199001011234",
			IdCardFront:  "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1515655661&di=7e3282b38f397402859c7ac632113b49&imgtype=jpg&er=1&src=http%3A%2F%2Fwww.qqzhi.com%2Fwenwen%2Fuploads%2Fservice.qq.com%2Fuserfiles%2FImage%2Fzakardchen%2Fmohu.jpg",
			IdCardBack:   "https://ss3.bdstatic.com/70cFv8Sh_Q1YnxGkpoWK1HF6hhy/it/u=2445163217,3569169920&fm=27&gp=0.jpg",
			IsOrgManager: false,
			IsOrgStaff:   true,
			IsCmtManager: false,
		}
		_, err := malUser.AddUser(user)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestMallUser_OrgManagerCount(t *testing.T) {
	malUser := NewMallUser()
	count, err := malUser.OrgManagerCount()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(count)
}

func TestMallUser_OrgManagerList(t *testing.T) {
	malUser := NewMallUser()
	users, err := malUser.OrgManagerList(1, 1, true)
	if err != nil {
		t.Error(err)
		return
	}
	for _, user := range users {
		fmt.Println(user)
	}
}

func TestMallUser_CmtManagerCount(t *testing.T) {
	malUser := NewMallUser()
	count, err := malUser.CmtManagerCount()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(count)
}
