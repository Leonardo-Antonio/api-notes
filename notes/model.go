package notes

type model struct {
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Body   string `json:"body" xml:"body"`
	IDUser int    `json:"id_user" xml:"id_user"`
	IDType int    `json:"id_type" xml:"id_type"`
}
