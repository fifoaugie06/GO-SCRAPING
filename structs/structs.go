package structs

type News struct {
	Title    string `json:"title,omitempty"`
	Image    string `json:"image,omitempty"`
	Link     string `json:"link,omitempty"`
	PostDate string `json:"post_date,omitempty"`
}

type NewsDetail struct {
	Title       string `json:"title,omitempty"`
	Image       string `json:"image,omitempty"`
	Description string `json:"description,omitempty"`
	PostDate    string `json:"post_date,omitempty"`
}
