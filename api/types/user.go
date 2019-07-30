package types

type User struct {
	Model
	Name      string     `json:"name"`
	Thumbmail string     `json:"thumbnail"`
	Location  Location   `json:"location"`
	Nick      string     `json:"nick"`
	Password  string     `json:"-"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	Todos     []*Todo    `json:"todos"`
	Articles  []*Article `json:"articles"`
}

type Location struct {
	Model
	Name      string `json:"name"`
	Thumbmail string `json:"thumbnail"`
}

type Todo struct {
	Model
	Task string `json:"task"`
	Done bool   `json:"bone"`
}
