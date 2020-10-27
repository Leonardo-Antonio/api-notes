package user

type model struct {
	ID int `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	LastName string `json:"last_name" xml:"last_name"`
	NickName string `json:"nick_name" xml:"nick_name"`
	Profile []byte `json:"profile" xml:"profile"`
	Email string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}
