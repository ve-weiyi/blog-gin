package types

type CleanMenuReq struct {
}

type CleanMenuResp struct {
	SuccessCount int64 `json:"success_count"` // 清空成功数量
}

type CreateMenuReq struct {
	MenuMeta
	ParentId  int64             `json:"parent_id"`         // 父id
	Path      string            `json:"path"`              // 路由地址
	Name      string            `json:"name"`              // 路由名字
	Component string            `json:"component"`         // Layout组件
	Redirect  *string           `json:"redirect,optional"` // 路由重定向
	Children  *[]*CreateMenuReq `json:"children,optional"`
}

type DeleteMenuReq struct {
	Ids []int64 `json:"ids"` // 主键id
}

type GetMenuReq struct {
	Id int64 `form:"id"` // 主键
}

type MenuVO struct {
	MenuMeta
	Id         *int64             `json:"id,optional"`          // 主键
	ParentId   int64              `json:"parent_id"`            // 父id
	Path       string             `json:"path"`                 // 路由地址
	Name       string             `json:"name"`                 // 路由名字
	Component  string             `json:"component"`            // Layout组件
	Redirect   *string            `json:"redirect,optional"`    // 路由重定向
	Type       string             `json:"type"`                 // 菜单类型（0代表目录、1代表菜单、2代表按钮、3代表外链）
	Title      *string            `json:"title,optional"`       // 菜单标题
	Icon       *string            `json:"icon,optional"`        // 菜单图标
	Rank       *int64             `json:"rank,optional"`        // 排序
	Perm       *string            `json:"perm,optional"`        // 权限标识
	Params     *[]*MenuMetaParams `json:"params,optional"`      // 参数
	KeepAlive  *int64             `json:"keep_alive,optional"`  // 是否缓存
	AlwaysShow *int64             `json:"always_show,optional"` // 是否一直显示菜单
	Visible    *int64             `json:"visible,optional"`     // 菜单是否可见
	Status     int64              `json:"status"`               // 状态 0正常 1禁用
	CreatedAt  int64              `json:"created_at"`           // 创建时间
	UpdatedAt  int64              `json:"updated_at"`           // 更新时间
	Children   *[]*MenuVO         `json:"children,optional"`
}

type QueryMenuListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	Name     *string   `json:"name,optional"`      // 路由名字
	Title    *string   `json:"title,optional"`     // 菜单标题
	Status   *int64    `json:"status,optional"`    // 状态
}

type SyncMenuReq struct {
	Menus []*CreateMenuReq `json:"menus"`
}

type SyncMenuResp struct {
	SuccessCount int64 `json:"success_count"` // 同步成功数量
}

type UpdateMenuReq struct {
	MenuMeta
	Id         int64              `json:"id"`                   // 主键
	ParentId   int64              `json:"parent_id"`            // 父id
	Path       string             `json:"path"`                 // 路由地址
	Name       string             `json:"name"`                 // 路由名字
	Component  string             `json:"component"`            // Layout组件
	Redirect   *string            `json:"redirect,optional"`    // 路由重定向
	Type       string             `json:"type"`                 // 菜单类型（0代表目录、1代表菜单、2代表按钮、3代表外链）
	Title      *string            `json:"title,optional"`       // 菜单标题
	Icon       *string            `json:"icon,optional"`        // 菜单图标
	Rank       *int64             `json:"rank,optional"`        // 排序
	Perm       *string            `json:"perm,optional"`        // 权限标识
	Params     *[]*MenuMetaParams `json:"params,optional"`      // 参数
	KeepAlive  *int64             `json:"keep_alive,optional"`  // 是否缓存
	AlwaysShow *int64             `json:"always_show,optional"` // 是否一直显示菜单
	Visible    *int64             `json:"visible,optional"`     // 菜单是否可见
	Status     int64              `json:"status"`               // 状态 0正常 1禁用
}
