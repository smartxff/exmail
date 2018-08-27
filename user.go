package exmail

import (
	"net/url"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

func (this *Exmail)userEndpoint(method string)(*url.URL, error){
	return url.ParseRequestURI(this.endPoint.String()  + "user/" + method)
}

func (this *Exmail)CreateUser(req CreateURequest)(*Response,error){
	uEndPoint,err := this.userEndpoint("create")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,fmt.Errorf("marshal req error: %v",err.Error())
	}

	r,err := http.NewRequest("POST",uEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	if err !=nil{
		return nil,err
	}
	response := &Response{}
	if err := json.Unmarshal(resp,response);err !=nil{
		return nil,fmt.Errorf("unmarshal resp error: %v",err.Error())
	}
	return response,nil
}

func (this *Exmail)UpdateUser(req User)(*Response,error){
	uEndPoint,err := this.userEndpoint("update")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("POST",uEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	updateUResponse := &Response{}
	if err := json.Unmarshal(resp,updateUResponse);err !=nil{
		return nil,err
	}
	return updateUResponse,nil
}

func (this *Exmail)DeleteUser(id string)(*Response,error){
	uEndPoint,err := this.userEndpoint("delete")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",uEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("userid",id)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	deleteUResponse := &Response{}
	if err := json.Unmarshal(resp,deleteUResponse);err !=nil{
		return nil,err
	}
	return deleteUResponse,nil
}

func (this *Exmail)GetUser(id string)(*User,error){
	uEndPoint,err := this.userEndpoint("get")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",uEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("userid",id)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	user := &User{}
	if err := json.Unmarshal(resp,user);err !=nil{
		return nil,err
	}
	return user,nil
}

func (this *Exmail)SimpleList(department_id,fetch_child string)(*SimpleUserList,error){
	uEndPoint,err := this.userEndpoint("simplelist")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",uEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("department_id", department_id)
	q.Add("fetch_child", fetch_child)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	users := &SimpleUserList{}
	if err := json.Unmarshal(resp,users);err !=nil{
		return nil,err
	}
	return users,nil
}

func (this *Exmail)List(department_id,fetch_child string)(*UserList,error){
	uEndPoint,err := this.userEndpoint("list")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",uEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("department_id", department_id)
	q.Add("fetch_child", fetch_child)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	users := &UserList{}
	if err := json.Unmarshal(resp,users);err !=nil{
		return nil,err
	}
	return users,nil
}

func (this *Exmail)BatchCheck(req UserListRequest)(*BatchCheckResponse,error){
	uEndPoint,err := this.userEndpoint("batchcheck")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("POST",uEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	uResponse := &BatchCheckResponse{}
	if err := json.Unmarshal(resp,uResponse);err !=nil{
		return nil,err
	}
	return uResponse,nil
}