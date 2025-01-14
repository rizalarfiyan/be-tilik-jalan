package response

type Home struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
	Website  string `json:"website"`
}
