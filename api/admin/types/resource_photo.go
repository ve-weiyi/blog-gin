package types

type CreatePhotoReq struct {
	AlbumId   int64  `json:"album_id"`
	PhotoName string `json:"photo_name"`
	PhotoDesc string `json:"photo_desc"`
	PhotoSrc  string `json:"photo_src"`
}

type DeletePhotoReq struct {
	Ids []int64 `json:"ids"`
}

type PhotoVO struct {
	Id        *int64 `json:"id,optional"`
	AlbumId   int64  `json:"album_id"`
	PhotoName string `json:"photo_name"`
	PhotoDesc string `json:"photo_desc"`
	PhotoSrc  string `json:"photo_src"`
	IsDelete  int64  `json:"is_delete"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type QueryPhotoListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	AlbumId  *int64    `json:"album_id,optional"`
	IsDelete *int64    `json:"is_delete,optional"`
}

type UpdatePhotoDeleteReq struct {
	Ids      []int64 `json:"ids"`
	IsDelete int64   `json:"is_delete"`
}

type UpdatePhotoReq struct {
	Id        int64  `json:"id"`
	AlbumId   int64  `json:"album_id"`
	PhotoName string `json:"photo_name"`
	PhotoDesc string `json:"photo_desc"`
	PhotoSrc  string `json:"photo_src"`
	IsDelete  int64  `json:"is_delete"`
}
