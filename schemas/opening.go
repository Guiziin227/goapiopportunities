package schemas

import ("gorm.io/gorm"
"time")

type Opening struct {
	gorm.Model

	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpeningResponse struct {
	ID uint `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	DeleteAt time.Time `json:"delete_at,omitempty"`
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool `json:"remote"`
	Link     string `json:"link"`
	Salary   int64 `json:"salary"`
}
