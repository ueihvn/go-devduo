package swagger

// swagger:model PlanService
type swaggerPlanService struct {
	// id for planservice. required for Update, Get
	// required: true
	// min: 1
	ID uint64 `json:"id,omitempty"`

	// id for userID. required for Create
	// required: true
	// min: 1
	UserID uint64 `json:"user_id,omitempty"`

	// title for planservice. required for Create
	// required
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`

	// price for planservice. required for Create
	// required
	// min: 0
	Price float64 `json:"price,omitempty"`
}

// swagger:parameters planServiceCreateUpdateReq
type swaggerPlanServiceCreateUpdateReq struct {
	// in: body
	Body struct {
		swaggerPlanService
	} `json:"body"`
}

// swagger:response planServiceCreateUpdateResp
type swaggerPlanServiceCreateUpdateResp struct {
	// in: body
	Body struct {
		// Example: true
		Status  bool
		Message string
		Data    swaggerPlanService
	} `json:"body"`
}

// swagger:response planServiceGetListResp
type swaggerPlanServiceGetListResp struct {
	// in: body
	Body struct {
		// Example: true
		Status  bool
		Message string
		Data    []swaggerPlanService
	} `json:"body"`
}
