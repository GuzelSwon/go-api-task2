package src

type Entity struct {
	Id          int    `json:"id"`
	Slug        string `json:"slug"`
	Url         string `json:"url"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Image       string `json:"image"`
	Thumbnail   string `json:"thumbnail"`
	Status      string `json:"status"`
	Category    string `json:"category"`
	PublishedAt string `json:"publishedAt"`
	UpdatedAt   string `json:"updatedAt"`
	UserId      int    `json:"userId"`
}
