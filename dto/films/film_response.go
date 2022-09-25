package filmsdto

import "dumbflix/models"

type FilmResponse struct {
	ID         int                     `json:"id"`
	Title      string                  `json:"title" form:"title" validate:"required"`
	Thumbnail  string                  `json:"image" form:"image" validate:"required"`
	Year       string                  `json:"year" form:"year" validate:"required"`
	Category   models.CategoryResponse `json:"category"`
	CategoryID int                     `json:"category_id"`
	Link_Film  string                  `json:"link_film"`
	Desc       string                  `json:"desc" form:"desc" validate:"required"`
}
