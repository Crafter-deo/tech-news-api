package websites

type Blogs struct {
	Headline string `json:"headline"`
	Url      string `json:"link"`
	Site     string `json:"site"`
}

var Sites = []string{"cnet", "codingdojo", "digitaltrends", "hackernews", "mashable", "medium"}

/* WEBSITES TO ADD
* https://www.engadget.com/
* https://digg.com/technology
* https://www.techradar.com/best
* https://www.theverge.com/ ???????
 */
