package main

type image struct {
	data []byte
}

type Article struct {
	AuthorID int     `json:"author_id"`
	Title    string  `json:"title"`
	CoordsN  float64 `json:"coords_n"`
	CoordsW  float64 `json:"coords_w"`
	Brief    string  `json:"brief"`
	Text     string  `json:"text"`
	Images   []image `json:"images"`
}

type Person struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
	Status       string `json:"status"`
	Follows      []int  `json:"follows"`
	FollowersNum int    `json:"followers_num"`
	Avatar       *image `json:"avatar"`
}
