package models

type Post struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Post   string `json:"post"`
	Author string `json:"author"`
}

type PostMutationModel struct {
	Title  string `json:"title"`
	Post   string `json:"post"`
	Author string `json:"author"`
}

type PostFilter struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Author string `json:"author"`
}

type PostAll struct {
	Count int    `json:"count"`
	Posts []Post `json:"posts"`
}
