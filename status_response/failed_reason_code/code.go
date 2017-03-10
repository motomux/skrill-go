package failedreasoncode

// Code describes failed_reason_code from Skrill API
type Code int

// Code list
const (
	CardIssue                            Code = 1  // Referred by Card Issuer
	InvalidMerchant                      Code = 2  // Invalid Merchant
	PickupCard                           Code = 3  // Pick‐up card
	Declined                             Code = 4  // Declined by Card Issuer
	InsufficientFUnds                    Code = 5  // Insufficient funds
	TransactionFailed                    Code = 6  // Transaction failed
	IncorrectPIN                         Code = 7  // Incorrect PIN
	PINTriesExceed                       Code = 8  // PIN tries exceed ‐ card blocked
	InvalidTransaction                   Code = 9  // Invalid Transaction
	TransactioLimitExceeded              Code = 10 // Transaction frequency limit exceeded
	InvalidAmount                        Code = 11 // Invalid Amount/ Amount too high /Limit Exceeded
	InvalidCreditCardOrBank              Code = 12 // Invalid credit card or bank account
	InvalidCard                          Code = 13 // Invalid card Issuer
	DuplicateTransaction                 Code = 15 // Duplicate transaction
	RetryTransaction                     Code = 19 // Retry transaction
	CardExpired                          Code = 24 // Card expired
	RequestedFuncNotAvailable            Code = 27 // Requested function not available
	LostCard                             Code = 28 // Lost/stolen card
	FormatFailure                        Code = 30 // Format Failure
	WrongSecurityCode                    Code = 32 // Card Security Code (CVV2/CVC2) Check Failed
	IllegalTransaction                   Code = 34 // Illegal Transaction
	CardRestricted                       Code = 37 // Card restricted by Card Issuer
	SecrityViolation                     Code = 38 // Security violation
	CardBlocked                          Code = 42 // Card blocked by Card Issuer
	BankOrNetworkUnavailable             Code = 44 // Card Issuing Bank or Network is not available
	ProcessingError                      Code = 45 // Processing error ‐ card type is not processed by the authorization centre
	SytemError                           Code = 51 // System error
	TransactionNotPermittedByAcquirer    Code = 58 // Transaction not permitted by acquirer
	TransactionNotPermittedForCardholder Code = 63 // Transaction not permitted for cardholder
	Wrong3DSVerification                 Code = 70 // Customer failed 3DS verification
	FraudRulesDeclined                   Code = 80 // Fraud rules declined
	ErrorWithProvider                    Code = 98 // Error in communication with provider
	Other                                Code = 99 // Other
)
