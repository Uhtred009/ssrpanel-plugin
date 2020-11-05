package client

import (
	"fmt"
	"testing"

	
	"github.com/sirupsen/logrus"
)
/*
func TestGetGet(t *testing.T) {
	HOST = "https://xiaoxue.ga"
	logrus.SetLevel(logrus.DebugLevel)
	result, _ := GetGet ("https://xiaoxue.ga/api/webapi/UserList/2", map[string]string{
		"key":       "123awew",
		"timestamp": "1604555981",
	})
	fmt.Printf("2222 \n")
	fmt.Printf("value: %+v \n", result)
	//Output:
}
*/

func TestGetUserList(t *testing.T) {
	HOST = "https://xiaoxue.ga"
	logrus.SetLevel(logrus.DebugLevel)
	result, _ := GetUserList(2, "txsnvhghmrmg4pjm")
	fmt.Printf("2222 \n")
	fmt.Printf("value00: %+v \n", result)
	for _, s :=range result {

		fmt.Printf("value00: %+v \n", *s)

	} 
	//Output:
}




