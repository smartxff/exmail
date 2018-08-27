package exmail

type CreateDRequest struct {
	Name                string            `json:"name"`
	ParentId            int64             `json:"parentid"`
	Order               int64             `json:"order"`
}

type Department struct {
	Id                  int64             `json:"id"`
	Name                string            `json:"name"`
	ParentId            int64            `json:"parentid"`
	Order               int64             `json:"order"`
}

type SearchDRequest struct {
	Name                string            `json:"name"`
	Fuzzy               int               `json:"fuzzy"`
}


type CreateURequest struct {
	Userid              string            `json:"userid"`
	Name                string            `json:"name"`
	Department          []int64           `json:"department"`
	Position            string            `json:"position"`
	Mobile              string            `json:"mobile"`
	Tel                 string             `json:"tel"`
	Extid               string             `json:"extid"`
	Gender              string             `json:"gender"`
	Slaves              []string             `json:"slaves"`
	Password            string             `json:"password,omitempty"`
	CpwdLogin           int                `json:"cpwd_login"`
}


type User struct {
	CreateURequest
	Enable              int                `json:"enable,omitempty"`
}

type UserListRequest struct {
	Userlist      []string                 `json:"userlist"`
}

