package types

type CreatePageReq struct {
	PageName       string    `json:"page_name"`
	PageLabel      string    `json:"page_label"`
	PageCover      string    `json:"page_cover"`
	IsCarousel     *int64    `json:"is_carousel,optional"`
	CarouselCovers *[]string `json:"carousel_covers,optional"`
}

type DeletePageReq struct {
	Ids []int64 `json:"ids"`
}

type PageVO struct {
	Id             *int64   `json:"id,optional"`
	PageName       string   `json:"page_name"`
	PageLabel      string   `json:"page_label"`
	PageCover      string   `json:"page_cover"`
	IsCarousel     int64    `json:"is_carousel"`
	CarouselCovers []string `json:"carousel_covers"`
	CreatedAt      int64    `json:"created_at"`
	UpdatedAt      int64    `json:"updated_at"`
}

type QueryPageListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	PageName *string   `json:"page_name,optional"`
}

type UpdatePageReq struct {
	Id             int64     `json:"id"`
	PageName       string    `json:"page_name"`
	PageLabel      string    `json:"page_label"`
	PageCover      string    `json:"page_cover"`
	IsCarousel     *int64    `json:"is_carousel,optional"`
	CarouselCovers *[]string `json:"carousel_covers,optional"`
}
