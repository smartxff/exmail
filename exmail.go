package exmail

import (
	"net/url"
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
)
type Exmail struct {
	endPoint             *url.URL                    `json:"end_point"`
	token                *Token                       `json:"token"`
	client               *http.Client                `json:"client"`
	Config               *config                     `json:"config"`
}

type config struct {
	Corpid               string                       `json:"corpid"`
	Corpsecret           string                       `json:"corpsecret"`
}


func NewExmail(location,corpid,corpsecret string) (*Exmail,error) {
	u, err := url.ParseRequestURI(location)
	if err != nil{
		return nil,err
	}

	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	u.Path += "cgi-bin/"

	exmail := new(Exmail)
	exmail.endPoint = u
	exmail.client = &http.Client{}
	exmail.Config = &config{corpid,corpsecret}

	return exmail,nil
}

func (this *Exmail)sendRequest(req *http.Request)([]byte,error){
	req.Header.Add("Accept", "application/json")
	resp, err := this.client.Do(req)
	if err !=nil{
		return nil,err
	}
	res, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}


	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:

		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("Authentication failed. : %s",res)
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("Service is not available (%s)",res)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("Internal server error: %s", res)
	}

	return nil, fmt.Errorf("Unknown response status : %s", res)
}






















