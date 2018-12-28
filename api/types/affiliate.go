package types

type Affiliate struct {
	Model
	Name      string `json:"name"`
	Url       string `json:"url"`
	Public    bool   `json:"public"`
	Thumbmail string `json:"thumbnail"`
	Preview   string `json:"preview"`
	Bucket    string
	Score     int    `json:"score"`
	Tags      []*Tag `gorm:"many2many:affiliate_tags" json:"tags"`
}
