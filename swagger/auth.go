package swagger

// swagger:response signUpResp
type swaggerSignUpResp struct {
	// in:body
	Body struct {
		// Example: true
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Id uint64 `json:"id"`
		} `json:"data"`
	} `json:"body"`
}

// swagger:response errorResp
type swaggerError struct {
	// in:body
	Body struct {
		// Example: false
		Status  bool   `json:"status"`
		Message string `json:"message"`
	} `json:"body"`
}

// swagger:response refreshCookieResp
type swaggerRefreshCookieResp struct {
	// in:body
	Body struct {
		// Example: true
		Status bool `json:"status"`
		// Example: successfully refresh cookie
		Message string `json:"message"`
	} `json:"body"`
}

// swagger:parameters signupLoginReq
type swaggerSignUpLoginReq struct {
	// in: body
	Body struct {
		// Example: email4@gmail.com
		Email string `json:"email"`
		// Example: passworduser4
		Password string `json:"password"`
	} `json:"body"`
}
