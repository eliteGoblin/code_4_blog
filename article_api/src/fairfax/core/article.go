package core



type Article struct {
	Id string 	 	`json:"id"`
	Title string 	`json:"title"`
	Date string 	`json:"date"`
	Body string 	`json:"body"`
	Tags []string 	`json:"tags"`
	CreatedTime string `json:"created_time"`
}



func (selfPtr *Article) IsValid() bool {
	return selfPtr.Id != ""
}
