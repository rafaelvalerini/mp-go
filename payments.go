package mp

type MPRequestPayments struct {
	TransactionAmount float64 `json:"transaction_amount"`
	Token             string  `json:"token"`
	Description       string  `json:"description"`
	Installments      int     `json:"installments"`
	PaymentMethodID   string  `json:"payment_method_id"`
	Payer             Payer   `json:"payer"`
}

type Payer struct {
	Email          string      `json:"email,omitempty"`
}

type MPResponsePayments struct {
	AuthorizationCode  interface{} `json:"authorization_code"`
	BinaryMode         bool        `json:"binary_mode"`
	CallForAuthorizeID interface{} `json:"call_for_authorize_id"`
	Captured           bool        `json:"captured"`
	Card               struct {
		Cardholder struct {
			Identification struct {
				Number string `json:"number"`
				Type   string `json:"type"`
			} `json:"identification"`
			Name string `json:"name"`
		} `json:"cardholder"`
		DateCreated     string      `json:"date_created"`
		DateLastUpdated string      `json:"date_last_updated"`
		ExpirationMonth int         `json:"expiration_month"`
		ExpirationYear  int         `json:"expiration_year"`
		FirstSixDigits  string      `json:"first_six_digits"`
		ID              interface{} `json:"id"`
		LastFourDigits  string      `json:"last_four_digits"`
	} `json:"card"`
	CollectorID               int           `json:"collector_id"`
	CouponAmount              int           `json:"coupon_amount"`
	CurrencyID                string        `json:"currency_id"`
	DateApproved              interface{}   `json:"date_approved"`
	DateCreated               string        `json:"date_created"`
	DateLastUpdated           string        `json:"date_last_updated"`
	DeductionSchema           interface{}   `json:"deduction_schema"`
	Description               string        `json:"description"`
	DifferentialPricingID     interface{}   `json:"differential_pricing_id"`
	ExternalReference         interface{}   `json:"external_reference"`
	FeeDetails                []interface{} `json:"fee_details"`
	ID                        int           `json:"id"`
	Installments              int           `json:"installments"`
	IssuerID                  string        `json:"issuer_id"`
	LiveMode                  bool          `json:"live_mode"`
	Metadata                  struct{}      `json:"metadata"`
	MoneyReleaseDate          interface{}   `json:"money_release_date"`
	NotificationURL           interface{}   `json:"notification_url"`
	OperationType             string        `json:"operation_type"`
	Order                     struct{}      `json:"order"`
	Payer                     Payer         `json:"payer"`
	PaymentMethodID           string        `json:"payment_method_id"`
	PaymentTypeID             string        `json:"payment_type_id"`
	Refunds                   []interface{} `json:"refunds"`
	SponsorID                 interface{}   `json:"sponsor_id"`
	StatementDescriptor       string        `json:"statement_descriptor"`
	Status                    string        `json:"status"`
	StatusDetail              string        `json:"status_detail"`
	TransactionAmount         float64       `json:"transaction_amount"`
	TransactionAmountRefunded int           `json:"transaction_amount_refunded"`
	TransactionDetails        struct {
		ExternalResourceURL      interface{} `json:"external_resource_url"`
		FinancialInstitution     interface{} `json:"financial_institution"`
		InstallmentAmount        float64     `json:"installment_amount"`
		NetReceivedAmount        float64     `json:"net_received_amount"`
		OverpaidAmount           float64     `json:"overpaid_amount"`
		PaymentMethodReferenceID interface{} `json:"payment_method_reference_id"`
		TotalPaidAmount          float64     `json:"total_paid_amount"`
	} `json:"transaction_details"`
}
