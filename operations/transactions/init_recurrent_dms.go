package transactions

import "github.com/TarlanPayments/gw-go-client/structures"

// InitRecurrentDMSAssembly is default structure for Init Rcurrent DMS Hold transactions operation
type InitRecurrentDMSAssembly struct {
	// HTTPData contains HTTP request method and operation action value for request in URL path
	opHTTPData structures.OperationRequestHTTPData
	// Command Data, isn't for any request type and in that case it's combined
	CommandData struct {
		structures.CommandData
		structures.CommandDataFormID
		structures.CommandDataTerminalMID
	} `json:"command-data,omitempty"`
	GeneralData   structures.GeneralData       `json:"general-data,omitempty"`
	PaymentMethod structures.PaymentMethodData `json:"payment-method-data"`
	Money         structures.MoneyData         `json:"money-data"`
	// System data contains user(cardholder) IPv4 address and IPv4 address in case of proxy
	System structures.SystemData `json:"system"`
}

// NewInitRecurrentDMSAssembly returns new instance with prepared HTTP request data InitRecurrentDMSAssembly
func NewInitRecurrentDMSAssembly() *InitRecurrentDMSAssembly {
	// Predefine default HTTP request data for sms operations
	var opd structures.OperationRequestHTTPData

	opd.SetHTTPMethod("POST")
	opd.SetOperationType(structures.InitRecurrentDMS)

	return &InitRecurrentDMSAssembly{
		opHTTPData: opd,
	}
}

// Implement methods for InitRecurrentDMSAssembly structure, form pck structures OperationRequestInterface

// GetHTTPMethod return HTTP method which will be used for send request
func (op *InitRecurrentDMSAssembly) GetHTTPMethod() string {
	return op.opHTTPData.GetHTTPMethod()
}

// GetOperationType return part of route path which will be used for send request
func (op *InitRecurrentDMSAssembly) GetOperationType() structures.OperationType {
	return op.opHTTPData.GetOperationType()
}
