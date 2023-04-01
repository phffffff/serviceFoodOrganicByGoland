package common

import "strings"

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"total"`

	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor" form:"next_cursor"`
}

func (p *Paging) FullFill() {
	if p.Limit <= 0 {
		p.Limit = 50
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
