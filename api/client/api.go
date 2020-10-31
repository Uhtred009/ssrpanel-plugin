package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	"strconv"
	"time"
	"bytes"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

var restyc *resty.Client

func init() {
	restyc = resty.New().
		SetTransport(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}).
		SetTimeout(5 * time.Second).
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(2))
}

var (
	HOST = "xiaoxue.ga"
)

func SetHost(host string) {
	HOST = host
}


type UserInfo struct {
	ID      uint
	VmessID string
	Email   string 	
	Port    int
}

// implement for vnet api get request
func get(url string, header map[string]string) (result string, err error) {
	logrus.WithFields(logrus.Fields{
		"url": url,
	}).Debug("get")

	r, err := restyc.R().SetHeaders(header).Get(url)
	if err != nil {
		return "", errors.Wrap(err, "get request error")
	}
	if r.StatusCode() != http.StatusOK {
		return "", errors.New(fmt.Sprintf("get request status: %d body: %s", r.StatusCode(), string(r.Body())))
	}
	body := r.Body()
	responseJson := BUnicodeToUtf8(body)

	return responseJson, nil
}

func GetUserList(nodeID int, key string) ([]UserInfo, error) {
	response, err := get(fmt.Sprintf("%s/api/web/v1/userList/%s", HOST, strconv.Itoa(nodeID)), map[string]string{
		"key":       key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil {
		return nil, err
	}
	if gjson.Get(response, "status").String() != "success" {
		return nil, errors.New((UnicodeToUtf8(gjson.Get(response, "message").String())))
	}
	value := gjson.Get(response, "data").String()
	if value == "" {
		return nil, errors.New("get data not found: " + response)
	}

	result := make([]UserInfo, 0)
	
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Convert string like \u4f60\u597d to utf-8 encode
// \u4f60\u597d means 你好(hello)
func UnicodeToUtf8(s string) string {
	slen := len(s)
	i := 0
	stringBuffer := new(bytes.Buffer)
	for i < slen {
		if s[i] == 92 && (s[i+1] == 85 || s[i+1] == 117) {
			temp, err := strconv.ParseInt(s[i+2:i+6], 16, 32)
			if err != nil {
				panic(err)
			}
			stringBuffer.WriteString(fmt.Sprintf("%c", temp))
			i += 6
			continue
		} else {
			stringBuffer.WriteByte(s[i])
			i++
			continue
		}
	}
	return stringBuffer.String()
}

// Convert string like \u4f60\u597d to utf-8 encode
// \u4f60\u597d means 你好(hello)
func BUnicodeToUtf8(s []byte) string {
	slen := len(s)
	i := 0
	stringBuffer := new(bytes.Buffer)
	for i < slen {
		if s[i] == 92 && (s[i+1] == 85 || s[i+1] == 117) {
			temp, err := strconv.ParseInt(string(s[i+2:i+6]), 16, 32)
			if err != nil {
				panic(err)
			}
			stringBuffer.WriteString(fmt.Sprintf("%c", temp))
			i += 6
			continue
		} else {
			stringBuffer.WriteByte(s[i])
			i++
			continue
		}
	}
	return stringBuffer.String()
}

