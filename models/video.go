package models

type Video struct {
	Id            uint64 `gorm:"primayKey" json:"id,omitempty"`
	AuthorId      uint64 `gorm:"author_id" json:"author_id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `gorm:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `gorm:"cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64  `gorm:"favorite_count" json:"favorite_count,omitempty"`
	CommentCount  int64  `gorm:"comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool   `gorm:"is_favorite" json:"is_favorite,omitempty"`
	CreateAt      string `gorm:"create_at"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}
