package websites

type Blogs struct {
	Headline string `json:"headline"`
	Url      string `json:"link"`
	Site     string `json:"site"`
}

var Sites = []string{"cnet", "codingdojo", "digitaltrends", "hackernews", "mashable", "medium"}