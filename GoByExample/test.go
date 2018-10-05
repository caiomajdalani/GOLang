// package main

// import (
// 	"encoding/base64"
// 	"fmt"
// )

// func main() {
// 	msg := "258369"
// 	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
// 	fmt.Println(encoded)
// 	decoded, err := base64.StdEncoding.DecodeString(encoded)
// 	if err != nil {
// 		fmt.Println("decode error:", err)
// 		return
// 	}
// 	fmt.Println(string(decoded))
// }
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://uat-api.dotzlabs.com/accounts/v1/connect/token"

	payload := strings.NewReader("grant_type=password&scope=pos.api&username=14900836850&password=MjU4MzY5")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic VEJOSjBCWjJLRkpSVU5OOlg0UEFIVFdMUkJOVDNVSg==")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Postman-Token", "97f69cda-2fea-6455-34a4-fb5d71f8927b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
