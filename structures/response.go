package structures

// Tarlan Payments Gateway's response structures
type (
	// UnauthorizedResponse will be return if Merchant authorization was incorrect, HTTP status code will be 401
	// Example of json: { "msg": "Unauthorized", "status": 401 }
	UnauthorizedResponse struct {
		Msg    string `json:"msg,omitempty"`
		Status int    `json:"status,omitempty"`
	}

	// TransactionResponse is structure of all transaction operations
	TransactionResponse struct {
		GateWay         gateWay         `json:"gw"`
		Error           gateWayError    `json:"error"`
		AcquirerDetails acquirerDetails `json:"acquirer-details"`
	}

	// Data parts for responses

	// gateWay Tarlan Payments system response
	gateWay struct {
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		StatusCode           int    `json:"status-code,omitempty"`
		StatusText           string `json:"status-text,omitempty"`
	}
	// gateWayError error structure part
	gateWayError struct {
		// Gateway error code
		Code int `json:"code,omitempty"`
		// Gateway error description
		Msg string `json:"msg,omitempty"`
	}

	// acquirerDetails response translated via Tarlan Payments system
	acquirerDetails struct {
		EciSLi            int    `json:"eci-sli,omitempty"`
		TerminalID        string `json:"terminal-mid,omitempty"`
		TransactionID     string `json:"transaction-id,omitempty"`
		ResultCode        string `json:"result-code,omitempty"`
		StatusText        string `json:"status-text,omitempty"`
		StatusDescription string `json:"status-description,omitempty"`
	}

	// ExploringStatusResponse the structure of all exploring operations
	// contained asked Tarlan Payments transaction id and it's status data
	// note: the response will be Json object to parse response you must assign to []ExploringStatusResponse
	ExploringStatusResponse struct {
		// GatewayTransactionID the past Tarlan Payments gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// ExploringStatus contained informational data of transaction
		Status []ExploreStatus `json:"status"`
	}

	// ExploringResultResponse the structure of all exploring operations
	// contained asked Tarlan Payments transaction id and it's result data
	// note: the response will be Json object to parse response you must assign to []ExploringResultResponse
	ExploringResultResponse struct {
		// GatewayTransactionID the past Tarlan Payments gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// ExploringStatus contained informational data of transaction
		DateCreated  string              `json:"date-created"`
		DateFinished string              `json:"date-finished"`
		ResultData   TransactionResponse `json:"result-data"`
	}

	// ExploringHistoryResponse the structure of all exploring operations
	// contained asked Tarlan Payments transaction id and it's result data
	// note: the response will be Json object to parse response you must assign to []ExploringHistoryResponse
	ExploringHistoryResponse struct {
		// GatewayTransactionID the past Tarlan Payments gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// History contained informational data of transaction
		History []ExploreStatus `json:"history"`
	}

	// ExploringRecurrentsResponse the structure of all exploring operations
	// contained asked Tarlan Payments transaction id and it's result data
	// note: the response will be Json object to parse response you must assign to []ExploringRecurrentsResponse
	ExploringRecurrentsResponse struct {
		// GatewayTransactionID the past Tarlan Payments gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// History contained informational data of transaction
		//@TODO add structure fields
	}

	// ExploringRefundResponse the structure of all exploring operations
	// contained asked Tarlan Payments transaction id and it's result data
	// note: the response will be Json object to parse response you must assign to []ExploringRefundResponse
	ExploringRefundResponse struct {
		// GatewayTransactionID the past Tarlan Payments gateway transaction id
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// History contained informational data of transaction
		//@TODO add structure fields
	}

	// ExploreStatus the structure of Tarlan Payments statuses for past transaction
	ExploreStatus struct {
		// GatewayTransactionID past transaction ID in Tarlan Payments system
		GatewayTransactionID string `json:"gateway-transaction-id,omitempty"`
		// StatusCode transaction status code
		StatusCode int `json:"status-code,omitempty"`
		// StatusText transaction status string representation
		StatusText string `json:"status-text,omitempty"`
		// StatusCodeGeneral transaction status code
		StatusCodeGeneral int `json:"status-code-general,omitempty"`
		// StatusTextGeneral transaction status code
		StatusTextGeneral string `json:"status-text-general,omitempty"`
	}

	// ExploreResult structure of past transaction result
	ExploreResult struct {
		GateWay         gateWay         `json:"gw"`
		Error           gateWayError    `json:"error"`
		AcquirerDetails acquirerDetails `json:"acquirer-details"`
	}
)
