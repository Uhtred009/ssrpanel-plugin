package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/rc452860/vnet/common/log"
	"github.com/rc452860/vnet/core"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/rc452860/vnet/model"
	"github.com/rc452860/vnet/utils/langx"
	"github.com/rc452860/vnet/utils/stringx"
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
	HOST = ""
)

func SetHost(host string) {
	HOST = host
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
	responseJson := stringx.BUnicodeToUtf8(body)

	return responseJson, nil
}


// GetNodeInfo
func GetNodeInfo(nodeID int, key string) (*model.NodeInfo, error) {
	response, err := get(fmt.Sprintf("%s/api/web/v1/node/%s", HOST, strconv.Itoa(nodeID)), map[string]string{
		"key":       key,
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	})
	if err != nil {
		return nil, err
	}

	if gjson.Get(response, "status").String() != "success" {
		return nil, errors.New(gjson.Get(response, "message").String())
	}
	value := gjson.Get(response, "data").String()
	if value == "" {
		return nil, errors.New("get data not found: " + response)
	}
	result := &model.NodeInfo{}
	err = json.Unmarshal([]byte(value), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}




}