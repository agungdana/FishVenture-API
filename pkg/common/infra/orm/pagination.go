package orm

import (
	"math"

	"gorm.io/gorm"
)

func (p *Paginantion) getPage() int {
	if p.Page == 0 {
		return 1
	}
	return p.Page
}

func (p *Paginantion) getLimit() int {
	if p.Limit == 0 {
		return 10
	}
	return p.Limit
}

func (p *Paginantion) getOffset() int {
	return (p.getPage() - 1) * p.getLimit()
}

func (p *Paginantion) getSort() string {
	if p.Sort == "" {
		p.Sort = "created_at desc"
	}
	return p.Sort
}

func (p *Paginantion) setSort() {

	if p.Sort == "" {
		p.Sort = "created_at"
	}

	switch p.Direction {
	case "", "desc":
		p.Sort = p.Sort + " DESC"
	case "asc":
		p.Sort = p.Sort + " ASC"
	}

}

func (p *Paginantion) setTotalPage(page int) {
	p.TotalPage = page
}

func (p *Paginantion) setTotalRows(totalRows int64) {
	p.TotalRows = totalRows
}

func Paginate(db *gorm.DB, pagination *Paginantion) func(db *gorm.DB) *gorm.DB {
	var totalRows int64

	db = db.Model(&pagination.ObjectTable)

	if pagination.Keyword != "" {
		db = db.Where(pagination.FindBy+" ?", pagination.Keyword)
	}

	db.Count(&totalRows)

	pagination.setTotalRows(totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.getLimit())))
	pagination.setTotalPage(totalPages)

	pagination.setSort()

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.getOffset()).Limit(pagination.getLimit()).Order(pagination.getSort())
	}
}
