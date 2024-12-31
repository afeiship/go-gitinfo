package gitinfo_test

import (
	"encoding/json"
	"fmt"
	"github.com/afeiship/go-gitinfo"
	"testing"
)

func TestParseUrl(f *testing.T) {
	url11 := "git@git.saybot.net:web_app/rhino-h5.git"
	info11, _ := gitinfo.ParseGitUrl(url11)
	jsonInfo, _ := json.Marshal(info11)
	fmt.Println("jsonInfo11: ", string(jsonInfo))

	// gitlab https
	url12 := "https://git.saybot.net/web_app/rhino-h5.git"
	info12, _ := gitinfo.ParseGitUrl(url12)
	jsonInfo, _ = json.Marshal(info12)
	fmt.Println("jsonInfo12: ", string(jsonInfo))

	// github
	//url21 := "git@github.com:afeiship/nx.git"
	//info21, _ := gitinfo.ParseGitUrl(url21)
	//jsonInfo, _ = json.Marshal(info21)
	//fmt.Println("jsonInfo21: ", string(jsonInfo))
}
