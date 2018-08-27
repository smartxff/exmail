package exmail

import (
	"net/url"
	"net/http"
	"encoding/json"
	"errors"
	"time"
	"fmt"
)

type Token struct {
	AccessToken          string                       `json:"access_token"`
	ExpiresIn            int64                        `json:"expires_in"`
	updataTime           time.Time                    `json:"-"`
}

func (this *Token)IsTimeOut()bool{
	if time.Now().Add(time.Second).UnixNano() >= this.updataTime.Add(time.Second * time.Duration(this.ExpiresIn)).UnixNano(){
		return true
	}
	return false
}

func (this *Exmail)baseEndpoint(method   string)(*url.URL, error){
	return url.ParseRequestURI(this.endPoint.String()  + method)
}

func (this *Exmail)initToken()(error){
	authEndPoint, err := this.baseEndpoint("gettoken")
	if err !=nil{
		return err
	}
	q := authEndPoint.Query()
	q.Add("corpid", this.Config.Corpid)
	q.Add("corpsecret", this.Config.Corpsecret)
	authEndPoint.RawQuery = q.Encode()
	req,err := http.NewRequest("GET",authEndPoint.String(),nil)
	if err !=nil{
		return err
	}
	resp,err := this.sendRequest(req)
	if err !=nil{
		return err
	}

	gToken := &GetTokenResponse{}
	if err := json.Unmarshal(resp, gToken);err !=nil{
		return fmt.Errorf("unmarsharl token resp error: ",err.Error())
	}
	if gToken.Errcode != 0{
		return errors.New(gToken.Errmsg)
	}
	this.token = &Token{
		updataTime: time.Now(),
		AccessToken: gToken.AccessToken,
		ExpiresIn: gToken.ExpiresIn,
	}
	return nil
}


func (this *Exmail)sendRequestWithAuth(req *http.Request)([]byte,error){
	//检查token是否存在，或是否超时
	if this.token == nil || this.token.IsTimeOut(){
		if err := this.initToken();err !=nil{
			return nil,fmt.Errorf("init token error: %v",err.Error())
		}
	}
	q := req.URL.Query()
	q.Add("access_token",this.token.AccessToken)
	req.URL.RawQuery = q.Encode()
	return this.sendRequest(req)
}
