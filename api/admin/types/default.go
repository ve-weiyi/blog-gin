package types

type PingReq struct {
}

type PingResp struct {
	Env         string `json:"env"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Runtime     string `json:"runtime"`
}
