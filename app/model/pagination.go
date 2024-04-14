package model

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pagination struct {
	TotalRecords int64 `json:"total_records"`
	TotalPage    int64 `json:"total_page"`
	CurrentPage  int64 `json:"current_page"`
	PageSize     int64 `json:"page_size"`
}

func (p *Pagination) AddTotalPage(totalRecords interface{}) {
	var addPage int64 = 0

	total, ok := totalRecords.(int64)
	if !ok {
		return
	}

	if total%p.PageSize != 0 {
		addPage++
	}

	p.TotalPage = total/p.PageSize + addPage
	p.TotalRecords = total
}

func (p *Pagination) Option() *options.FindOptions {
	skip := p.CurrentPage*p.PageSize - p.PageSize
	return &options.FindOptions{
		Limit: &p.PageSize,
		Skip:  &skip,
	}
}
