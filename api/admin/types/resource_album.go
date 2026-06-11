package types

type AlbumVO struct {
	Id         *int64 `json:"id,optional"`
	AlbumName  string `json:"album_name"`
	AlbumDesc  string `json:"album_desc"`
	AlbumCover string `json:"album_cover"`
	IsDelete   int64  `json:"is_delete"`
	Status     int64  `json:"status"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
	PhotoCount int64  `json:"photo_count"`
}

type CreateAlbumReq struct {
	AlbumName  string `json:"album_name"`
	AlbumDesc  string `json:"album_desc"`
	AlbumCover string `json:"album_cover"`
	Status     int64  `json:"status"`
}

type DeleteAlbumReq struct {
	Ids []int64 `json:"ids"`
}

type GetAlbumReq struct {
	Id int64 `form:"id"`
}

type QueryAlbumListReq struct {
	PageQuery
	Page      *int64    `json:"page,optional"`      // 当前页码
	PageSize  *int64    `json:"page_size,optional"` // 每页数量
	Sorts     *[]string `json:"sorts,optional"`     // 排序
	AlbumName *string   `json:"album_name,optional"`
	IsDelete  *int64    `json:"is_delete,optional"`
}

type UpdateAlbumDeleteReq struct {
	Ids      []int64 `json:"ids"`
	IsDelete int64   `json:"is_delete"`
}

type UpdateAlbumReq struct {
	Id         int64  `json:"id"`
	AlbumName  string `json:"album_name"`
	AlbumDesc  string `json:"album_desc"`
	AlbumCover string `json:"album_cover"`
	IsDelete   int64  `json:"is_delete"`
	Status     int64  `json:"status"`
}
