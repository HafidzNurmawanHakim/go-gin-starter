package schema

import (
	"math"

	"gorm.io/gorm"
)

type Paginate struct {
	Limit int
	Page  int
	NextPage int
	PrevPage int
	TotalItems int
}

func NewPaginate(limit int, page int) *Paginate {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if page <= 0 {
		page = 1 // Default page
	}

	return &Paginate{
		Limit:      limit,
		Page:       page,
		NextPage:   0,
		PrevPage:   0,
		TotalItems: 0,
	}
}


func (p *Paginate) PaginatedResult(db *gorm.DB) (meta Paginate, result *gorm.DB) {

	var totalItems int64
	db.Count(&totalItems)

	offset := (p.Page - 1) * p.Limit
	query := db.Offset(offset).Limit(p.Limit)

	totalPages := int(math.Ceil(float64(totalItems)/float64(p.Limit)))
	nextPage := 0
	prevPage := 0

	if p.Page < totalPages {
		nextPage = p.Page + 1
	}

	if p.Page > 1 {
		prevPage = p.Page - 1
	}

	meta = *&Paginate{
		Page: p.Page,
		Limit: p.Limit,
		NextPage: nextPage,
		PrevPage: prevPage,
		TotalItems: int(totalItems),
	}

	return meta, query
}
