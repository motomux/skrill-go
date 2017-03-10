package currency

// Currency is the list of supported currencies.
type Currency string

// Currencies which can be used for a payment through Skrill
const (
	EUR Currency = "EUR" // Euro TWD Taiwan Dollar
	USD Currency = "USD" // U.S. Dollar THB Thailand Baht
	GBP Currency = "GBP" // British Pound CZK Czech Koruna
	HKD Currency = "HKD" // Hong Kong Dollar HUF Hungarian Forint
	SGD Currency = "SGD" // Singapore Dollar BGN Bulgarian Leva
	JPY Currency = "JPY" // Japanese Yen PLN Polish Zloty
	CAD Currency = "CAD" // Canadian Dollar ISK Iceland Krona
	AUD Currency = "AUD" // Australian Dollar INR Indian Rupee
	CHF Currency = "CHF" // Swiss Franc KRW South‐Korean Won
	DKK Currency = "DKK" // Danish Krone ZAR South‐African Rand
	SEK Currency = "SEK" // Swedish Krona RON Romanian Leu New
	NOK Currency = "NOK" // Norwegian Krone HRK Croatian Kuna
	ILS Currency = "ILS" // Israeli Shekel JOD Jordanian Dinar
	MYR Currency = "MYR" // Malaysian Ringgit OMR Omani Rial
	NZD Currency = "NZD" // New Zealand Dollar RSD Serbian Dinar
	TRY Currency = "TRY" // New Turkish Lira TND Tunisian Dinar
	AED Currency = "AED" // Utd. Arab Emir. Dirham BHD Bahraini Dinar
	MAD Currency = "MAD" // Moroccan Dirham KWD Kuwaiti Dinar
	QAR Currency = "QAR" // Qatari Rial
	SAR Currency = "SAR" // Saudi Riy
)
