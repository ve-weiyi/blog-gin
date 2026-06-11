package types

type PingReq struct {
}

type PingResp struct {
	Env         string `json:"env"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Runtime     string `json:"runtime"`
	Description string `json:"description"`
}
