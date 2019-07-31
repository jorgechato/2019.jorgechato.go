package api

type Affiliate struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Public    bool   `json:"public"`
	Thumbmail string `json:"thumbnail"`
	Preview   string `json:"preview"`
	Bucket    string
	Score     int `json:"score"`
}
