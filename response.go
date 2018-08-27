package exmail

type CreateDResponse struct {
	Response
	Id               int64                `json:"id"`
}

type Response struct {
	Errcode          int64                `json:"errcode"`
	Errmsg           string               `json:"errmsg"`
}

type ListDepartment struct {
	Response
	Department       []Department          `json:"department"`
}

type GetTokenResponse struct {
	Errcode          int64                `json:"errcode"`
	Errmsg           string               `json:"errmsg"`
	AccessToken      string               `json:"access_token"`
	ExpiresIn        int64                `json:"expires_in"`
}

type SimpleUser struct {
	Userid              string            `json:"userid"`
	Name                string             `json:"name"`
	T                   []int64            `json:"t"`
}

type SimpleUserList struct {
	Response
	UserList            []SimpleUser      `json:"userlist"`
}

type UserList struct {
	Response
	UserList            []CreateURequest    `json:"userlist"`
}


type BatchCheckResponse struct {
	Response
	List                []struct{
		User            string              `json:"user"`
		Utype            string              `json:"type"`
	}                  `json:"list"`
}