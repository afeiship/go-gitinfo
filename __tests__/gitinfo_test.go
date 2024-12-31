package gitinfo_test

import (
	"encoding/json"
	"fmt"
	"github.com/afeiship/go-gitinfo"
	"testing"
)

func TestParseUrl(f *testing.T) {
	url11 := "git@git.saybot.net:web_app/rhino-h5.git"
	info11 := gitinfo.ParseUrl(url11)
	jsonInfo, _ := json.Marshal(info11)
	fmt.Println("jsonInfo: ", string(jsonInfo))
	fmt.Println("info: ", info11)

	// gitlab https
	url12 := "https://git.saybot.net/web_app/rhino-h5.git"
	info12 := gitinfo.ParseUrl(url12)
	jsonInfo, _ = json.Marshal(info12)
	fmt.Println("jsonInfo: ", string(jsonInfo))
	fmt.Println("info: ", info12)

	// github
	url2 := "git@github.com:afeiship/nx.git"
	info2 := gitinfo.ParseUrl(url2)
	jsonInfo, _ = json.Marshal(info2)
	fmt.Println("jsonInfo: ", string(jsonInfo))
	fmt.Println("info: ", info2)
}
