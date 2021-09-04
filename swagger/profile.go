package swagger

import "github.com/ueihvn/go-devduo/model"

// swagger:parameters profileCreateUpdateReq
type swaggerProfileCreateUpdateReq struct {
	// in: body
	Body struct {
		model.ProfileJSON
	} `json:"body"`
}

// swagger:response profileCreateUpdateResp
type swaggerProfileCreateUpdateResp struct {
	// in: body
	Body struct {
		// Example: true
		Status  bool
		Message string
		Data    model.ProfileJSON
	} `json:"body"`
}
