package types

type CreateCommentReq struct {
	TopicId        *int64  `json:"topic_id,optional"`
	ParentId       *int64  `json:"parent_id,optional"`
	ReplyId        *int64  `json:"reply_id,optional"`
	ReplyUserId    *string `json:"reply_user_id,optional"`
	CommentContent string  `json:"comment_content"`
	Status         *int64  `json:"status,optional"`
	Type           int64   `json:"type"`
}

type LikeCommentReq struct {
	CommentId int64 `json:"comment_id"`
}

type QueryCommentListReq struct {
	PageQuery
	Page     *int64    `json:"page,optional"`      // 当前页码
	PageSize *int64    `json:"page_size,optional"` // 每页数量
	Sorts    *[]string `json:"sorts,optional"`     // 排序
	TopicId  *int64    `json:"topic_id,optional"`
	ParentId *int64    `json:"parent_id,optional"`
	Type     *int64    `json:"type,optional"`
}

type QueryCommentReplyListReq struct {
	PageQuery
	CommentId int64     `json:"comment_id"`
	Page      *int64    `json:"page,optional"`      // 当前页码
	PageSize  *int64    `json:"page_size,optional"` // 每页数量
	Sorts     *[]string `json:"sorts,optional"`     // 排序
	TopicId   *int64    `json:"topic_id,optional"`
	Type      *int64    `json:"type,optional"`
}

type UpdateCommentReq struct {
	CommentId      int64   `json:"comment_id"`
	ReplyUserId    *string `json:"reply_user_id,optional"`
	CommentContent string  `json:"comment_content"`
	Status         *int64  `json:"status,optional"`
}
