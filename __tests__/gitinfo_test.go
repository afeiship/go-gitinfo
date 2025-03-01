package gitinfo_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/afeiship/go-gitinfo"
)

func TestParseUrl(f *testing.T) {
	var jsonInfo []byte

	url11 := "git@git.saybot.net:web_app/rhino-h5.git"
	info11, _ := gitinfo.ParseGitUrl(url11)
	// jsonInfo, _ := json.Marshal(info11)
	// fmt.Println("jsonInfo11: ", string(jsonInfo))
	if info11.Hostname != "git.saybot.net" || info11.Owner != "web_app" || info11.Repo != "rhino-h5" {
		f.Error("TestParseUrl failed")
	}

	url13 := "git@git.saybot.net:ACE/courseware/pptWebPlayer.git"
	info13, _ := gitinfo.ParseGitUrl(url13)
	jsonInfo, _ = json.Marshal(info13)
	fmt.Println("jsonInfo13: ", string(jsonInfo))
	if info13.Hostname != "git.saybot.net" || info13.Owner != "ACE" || info13.Repo != "courseware/pptWebPlayer" {
		f.Error("TestParseUrl failed")
	}

	url14 := "git@git.saybot.net:aric.zheng/frontend-ci.git"
	info14, _ := gitinfo.ParseGitUrl(url14)
	jsonInfo, _ = json.Marshal(info14)
	if info14.Hostname != "git.saybot.net" || info14.Owner != "aric.zheng" || info14.Repo != "frontend-ci" {
		f.Error("TestParseUrl failed")
	}

	// gitlab https
	url12 := "https://git.saybot.net/web_app/rhino-h5.git"
	info12, _ := gitinfo.ParseGitUrl(url12)
	jsonInfo, _ = json.Marshal(info12)
	//fmt.Println("jsonInfo12: ", string(jsonInfo))
	if info12.Hostname != "git.saybot.net" || info11.Owner != "web_app" || info11.Repo != "rhino-h5" {
		f.Error("TestParseUrl failed")
	}

	// github
	url21 := "git@github.com:afeiship/nx.git"
	info21, _ := gitinfo.ParseGitUrl(url21)
	// jsonInfo, _ = json.Marshal(info21)
	// fmt.Println("jsonInfo21: ", string(jsonInfo))
	if info21.Hostname != "github.com" || info21.Owner != "afeiship" || info21.Repo != "nx" {
		f.Error("TestParseUrl failed")
	}

	// github
	url22 := "https://github.com/afeiship/nx.git"
	info22, _ := gitinfo.ParseGitUrl(url22)
	// jsonInfo, _ = json.Marshal(info22)
	// fmt.Println("jsonInfo22: ", string(jsonInfo))
	if info22.Hostname != "github.com" || info22.Owner != "afeiship" || info22.Repo != "nx" {
		f.Error("TestParseUrl failed")
	}
}

func TestGetGitInfo(f *testing.T) {
	_, info := gitinfo.GetGitInfo()
	jsonInfo, _ := json.Marshal(info)
	fmt.Println("jsonInfo: ", string(jsonInfo))
}
