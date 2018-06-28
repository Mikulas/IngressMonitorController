package v3

import (
	"net/http"
	"time"
	"net/url"
	"fmt"
	"log"
)

type Uptrends struct {
	username, password string
	netClient *http.Client
}

func MakeUptrends(username, password string) *Uptrends {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return &Uptrends{
		username: username,
		password: password,
		netClient: netClient,
	}
}

func (up *Uptrends) request(req *http.Request) (*http.Response, error) {
	req.URL = &url.URL{
		Scheme: "https",
		Host: "api.uptrends.com",
		Path: fmt.Sprintf("/v3%v", req.URL.Path),
		RawQuery: req.URL.RawQuery,
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(up.username, up.password)
	log.Printf("requesting %v %v\n", req.Method, req.URL)
	return up.netClient.Do(req)
}

//{
//"UserName": "9c4b2571cf3345b29cfb4f1ad2c1fd80",
//"Password": "ifWbe52MLBK+SkINKLICdCkRa1eCJ7SD",
//"AccountId": "322941",
//"OperatorName": "Mikulas Dite",
//"status": "OK"
//}
