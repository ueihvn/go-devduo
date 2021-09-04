package swagger

import "github.com/ueihvn/go-devduo/model"

// swagger:response userResp
type swaggUserResp struct {
	// in:body
	Body struct {
		// Example: true
		Status bool `json:"status"`
		// Example: successfully get user by user id
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	} `json:"body"`
}
