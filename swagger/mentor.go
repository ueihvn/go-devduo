package swagger

import "github.com/ueihvn/go-devduo/handler"

// swagger:response mentorResp
type swaggerMentorResp struct {
	// in:body
	Body struct {
		// Example: true
		Status  bool
		Message string
		Data    []handler.Mentor
	}
}
