package entities

type SearchOptions int

const (
	Authors = iota
	Articles
)

type image struct {
	data []byte
}

type PersonIDStruct struct {
	ID int `json:"id"`
}

type LikeStruct struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type Article struct {
	ID           int      `json:"id"`
	AuthorID     int      `json:"author_id"`
	AuthorName   string   `json:"author_name"`
	AuthorTag    string   `json:"author_tag"`
	AuthorAvatar string   `json:"author_avatar"`
	Title        string   `json:"title"`
	Date         int64    `json:"date"`
	CoordsN      float64  `json:"coords_n"`
	CoordsW      float64  `json:"coords_w"`
	Brief        string   `json:"brief"`
	Text         string   `json:"text"`
	Images       []string `json:"images"`
}

// Maybe add likes
type Person struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
	Status       string `json:"status"`
	Follows      []int  `json:"follows"`
	FollowersNum int    `json:"followers_num"`
	Avatar       string `json:"avatar"`
	Likes        []int  `json:"-"`
}

type RegisterPerson struct {
	Name     string `json:"name"`
	Tag      string `json:"tag"`
	Password string `json:"password"`
}

type LoginPerson struct {
	Tag      string `json:"tag"`
	Password string `json:"password"`
}

type RefreshStruct struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type NewsStruct struct {
	UserId   int `json:"user_id"`
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

type SearchStruct struct {
	UserId            int           `json:"user_id"`
	SearchFieldString string        `json:"search_field_string"`
	SearchingType     SearchOptions `json:"searching_type"`
}
