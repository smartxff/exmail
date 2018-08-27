package exmail

import (
	"net/url"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
)

func (this *Exmail)departmentEndpoint(method string)(*url.URL, error){
	return url.ParseRequestURI(this.endPoint.String()  + "department/" + method)
}

func (this *Exmail)CreateDepartment(req CreateDRequest)(*CreateDResponse,error){
	dEndPoint,err := this.departmentEndpoint("create")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,fmt.Errorf("marshal req error: %v",err.Error())
	}

	r,err := http.NewRequest("POST",dEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	if err !=nil{
		return nil,err
	}
	createDResponse := &CreateDResponse{}
	if err := json.Unmarshal(resp,createDResponse);err !=nil{
		return nil,fmt.Errorf("unmarshal resp error: %v",err.Error())
	}
	return createDResponse,nil
}

func (this *Exmail)UpdateDepartment(req Department)(*Response,error){
	dEndPoint,err := this.departmentEndpoint("update")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("POST",dEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	updateDResponse := &Response{}
	if err := json.Unmarshal(resp,updateDResponse);err !=nil{
		return nil,err
	}
	return updateDResponse,nil
}

func (this *Exmail)DeleteDepartment(id string)(*Response,error){
	dEndPoint,err := this.departmentEndpoint("delete")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",dEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("id",id)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	deleteDResponse := &Response{}
	if err := json.Unmarshal(resp,deleteDResponse);err !=nil{
		return nil,err
	}
	return deleteDResponse,nil
}

func (this *Exmail)GetDepartmentList(id string)(*ListDepartment,error){
	dEndPoint,err := this.departmentEndpoint("list")
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("GET",dEndPoint.String(),nil)
	if err !=nil{
		return nil,err
	}

	q := r.URL.Query()
	q.Add("id",id)
	r.URL.RawQuery = q.Encode()


	resp,err := this.sendRequestWithAuth(r)
	listDepartment := &ListDepartment{}
	if err := json.Unmarshal(resp,listDepartment);err !=nil{
		return nil,err
	}
	return listDepartment,nil
}

func (this *Exmail)SearchDepartment(req SearchDRequest)(*ListDepartment,error){
	dEndPoint,err := this.departmentEndpoint("search")
	if err !=nil{
		return nil,err
	}
	c,err := json.Marshal(req)
	if err !=nil{
		return nil,err
	}

	r,err := http.NewRequest("POST",dEndPoint.String(),bytes.NewReader(c))
	if err !=nil{
		return nil,err
	}
	resp,err := this.sendRequestWithAuth(r)
	listDepartment := &ListDepartment{}
	if err := json.Unmarshal(resp,listDepartment);err !=nil{
		return nil,err
	}
	return listDepartment,nil
}


