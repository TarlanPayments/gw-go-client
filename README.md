# Tarlan Payments Gateway v3 Go package

This package provide ability to make requests to Tarlan Payments Gateway API v3.

## Documentation
This `README` provide introduction to the library usage.

### Supported operations
- Transactions
  - SMS
  - DMS HOLD
  - DMS CHARGE
  - CANCEL
  - MOTO SMS
  - MOTO DMS
  - CREDIT
  - P2P
  - B2P
  - INIT RECURRENT DMS
  - RECURRENT DMS
  - INIT RECURRENT SMS
  - RECURRENT SMS
  - REFUND
  - REVERSAL

- Information
  - HISTORY
  - RECURRENTS
  - REFUNDS
  - RESULT
  - STATUS

- Verifications
  - Verify card 3-D Secure enrollment
  - Complete card verification

- Tokenization
  - Create payment data token

#### Basic usage
```go
    // Setup your credentials for authorized requests
    AccGUID := "someAccountGUID" // Your account GUID from Tarlan Payments
    SecKey := "someSecretKey" // Your API secret key

    // Setup new Gateway Client
    gateCli, gateCliErr := tarlanpaymentsgateway.NewGatewayClient(AccGUID, SecKey)
    if gateCliErr != nil {
        log.Fatal(gateCliErr)
    }
	gateCli.API.BaseURI = "https://<Gateway URL>"

    // Prepare operation builder to handle your operations
    specOpsBuilder :=  gateCli.OperationBuilder()

    // Now, define your special operation for processing
    sms := specOpsBuilder.NewSms()

    // Set transaction data
    sms.GeneralData.OrderData.OrderDescription = "Operation Single-Message Transactions"
    sms.GeneralData.CustomerData.Email = "some@email.com"
    sms.PaymentMethod.Pan = "1111111111111111"
    sms.PaymentMethod.ExpMmYy = "10/60"
    sms.PaymentMethod.Cvv = "123"
    sms.Money.Amount = 1500
    sms.Money.Currency = "USD"
    sms.System.UserIP = "199.99.99.1"
    sms.System.XForwardedFor = "199.99.99.1"

    // Now process the operation
    opResp, opErr := gateCli.NewRequest(sms)
    if opErr != nil {
        log.Printf(opErr)
    }
```

### Card verification

```go
// set card verification init mode for a payment
payment.CommandData.CardVerificationMode = structures.CardVerificationModeInit

// complete card verification
request := specOpsBuilder.NewVerifyCard()
request.VerifyCardData.GWTransactionID = initialPaymentGatewayId
gateCli.NewRequest(request)

// set card verification verify mode for subsequent payments
newPayment.CommandData.CardVerificationMode = structures.CardVerificationModeVerify
```

### Payment data tokenization

```go
// option 1: create a payment with flag to save payment data
payment.CommandData.PaymentMethodDataSource = structures.DataSourceSaveToGateway

// option 2: send "create token" request with payment data
operation = specOpsBuilder.NewCreateToken();
operation.PaymentMethod.Pan = "1111111111111111"
operation.PaymentMethod.ExpMmYy = "10/60"
operation.PaymentMethod.CardholderName = "John Doe"
operation.Money.Currency = "EUR"
gateCli.NewRequest(operation)

// send a payment with flag to load payment data by token
newPayment.CommandData.PaymentMethodDataSource = structures.DataSourceUseGatewaySavedCardholderInitiated
newPayment.CommandData.PaymentMethodDataToken = "<initial gateway-transaction-id>"
```

### Submit bugs and feature requests
Bugs and feature request are tracked on [GitHub](https://github.com/TarlanPayments/gw-go-client/issues)


### How to run unit tests by executing command in terminal:
```bash
$: go test ./...
```

### License
This library is licensed under the MIT License - see the `LICENSE` file for details.