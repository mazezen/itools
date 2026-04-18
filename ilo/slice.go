package ilo

import "github.com/samber/lo"

type Result[T any] struct {
  	Items      []T `json:"items"`
  	Total      int `json:"total"`
  	Page       int `json:"page"`
  	Size       int `json:"size"`
  	TotalPages int `json:"totalPages"`
  	Offset     int `json:"offset"`
  	Limit      int `json:"limit"`
  	HasPrev    bool `json:"hasPrev"`
  	HasNext    bool `json:"hasNext"`
  }

  
func Paginate[T any](items []T, page, size int) Result[T] {
  	total := len(items)

  	if size < 1 {
  		size = 10
  	}

  	totalPages := 0
  	if total > 0 {
  		totalPages = (total + size - 1) / size
  	}

  	if totalPages == 0 {
  		page = 1
  	} else {
  		if page < 1 {
  			page = 1
  		}
  		if page > totalPages {
  			page = totalPages
  		}
  	}

  	offset := (page - 1) * size
  	limit := size
  	paged := lo.Subset(items, offset, uint(limit))

  	return Result[T]{
  		Items:      paged,
  		Total:      total,
  		Page:       page,
  		Size:       size,
  		TotalPages: totalPages,
  		Offset:     offset,
  		Limit:      limit,
  		HasPrev:    page > 1,
  		HasNext:    page < totalPages,
  	}
  }