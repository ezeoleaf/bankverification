package bankverification

// AccountType represents the account type
type AccountType struct {
	Type            string
	Algorithm       string
	ClearingLength  int
	SerialLength    int
	WeightedNumbers int
}

// ClearingNumber represents a clearing number entry
type ClearingNumber struct {
	Interval            string
	BankName            string
	AccountType         AccountType
	MinLength           int
	MaxLength           int
	ChecksumForClearing bool
	Zerofill            bool
}

const (
	AccountType1 = "TYPE1"
	AccountType2 = "TYPE2"
	AccountType3 = "TYPE3"
	AccountType4 = "TYPE4"
	AccountType5 = "TYPE5"
)

var defaultClearingNumbers = []ClearingNumber{
	{Interval: "1100..1199", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "1200..1399", BankName: "Danske Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "1400..2099", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "2300..2399", BankName: "Ålandsbanken", AccountType: AccountTypes[AccountType2]},
	{Interval: "2400..2499", BankName: "Danske Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "3000..3299", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "3300..3300", BankName: "Nordea", AccountType: AccountTypes[AccountType3], MinLength: 10}, // Personkonto.
	{Interval: "3301..3399", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "3400..3409", BankName: "Länsförsäkringar Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "3410..3781", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "3782..3782", BankName: "Nordea", AccountType: AccountTypes[AccountType3], MinLength: 10}, // Personkonto.
	{Interval: "3783..3999", BankName: "Nordea", AccountType: AccountTypes[AccountType1]},
	{Interval: "4000..4999", BankName: "Nordea", AccountType: AccountTypes[AccountType2]},
	{Interval: "5000..5999", BankName: "SEB", AccountType: AccountTypes[AccountType1]},
	{Interval: "6000..6999", BankName: "Handelsbanken", AccountType: AccountTypes[AccountType4], MaxLength: 9, MinLength: 8},
	{Interval: "7000..7999", BankName: "Swedbank", AccountType: AccountTypes[AccountType1]},
	{Interval: "8000..8999", BankName: "Swedbank", AccountType: AccountTypes[AccountType5], MinLength: 10, ChecksumForClearing: true, Zerofill: true}, // Can be fewer chars but must be zero-filled, so let's call it 10.
	{Interval: "9020..9029", BankName: "Länsförsäkringar Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9040..9049", BankName: "Citibank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9060..9069", BankName: "Länsförsäkringar Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "9090..9099", BankName: "Royal Bank of Scotland", AccountType: AccountTypes[AccountType2]},
	{Interval: "9100..9109", BankName: "Nordnet Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9120..9124", BankName: "SEB", AccountType: AccountTypes[AccountType1]},
	{Interval: "9130..9149", BankName: "SEB", AccountType: AccountTypes[AccountType1]},
	{Interval: "9150..9169", BankName: "Skandiabanken", AccountType: AccountTypes[AccountType2]},
	{Interval: "9170..9179", BankName: "Ikano Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "9180..9189", BankName: "Danske Bank", AccountType: AccountTypes[AccountType3], MinLength: 10},
	{Interval: "9190..9199", BankName: "DNB Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9230..9239", BankName: "Marginalen Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "9250..9259", BankName: "SBAB", AccountType: AccountTypes[AccountType1]},
	{Interval: "9260..9269", BankName: "DNB Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9270..9279", BankName: "ICA Banken", AccountType: AccountTypes[AccountType1]},
	{Interval: "9280..9289", BankName: "Resurs Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "9300..9349", BankName: "Swedbank (fd. Sparbanken Öresund)", AccountType: AccountTypes[AccountType3], MinLength: 10, Zerofill: true},
	{Interval: "9390..9399", BankName: "Landshypotek AB", AccountType: AccountTypes[AccountType2]},
	{Interval: "9400..9449", BankName: "Forex Bank", AccountType: AccountTypes[AccountType1]},
	{Interval: "9460..9469", BankName: "Santander Consumer Bank AS", AccountType: AccountTypes[AccountType1]},
	{Interval: "9470..9479", BankName: "BNP Paribas Fortis Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9500..9549", BankName: "Nordea/Plusgirot", AccountType: AccountTypes[AccountType5], MaxLength: 10, MinLength: 1},
	{Interval: "9550..9569", BankName: "Avanza Bank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9570..9579", BankName: "Sparbanken Syd", AccountType: AccountTypes[AccountType3], MinLength: 10, Zerofill: true},
	{Interval: "9590..9599", BankName: "Erik Penser AB", AccountType: AccountTypes[AccountType2]},
	{Interval: "9630..9639", BankName: "Lån & Spar Bank Sverige", AccountType: AccountTypes[AccountType1]},
	{Interval: "9640..9649", BankName: "Nordax Bank AB", AccountType: AccountTypes[AccountType2]},
	{Interval: "9660..9669", BankName: "Amfa Bank AB", AccountType: AccountTypes[AccountType2]},
	{Interval: "9670..9679", BankName: "JAK Medlemsbank", AccountType: AccountTypes[AccountType2]},
	{Interval: "9680..9689", BankName: "BlueStep Finans AB", AccountType: AccountTypes[AccountType1]},
	{Interval: "9700..9709", BankName: "Ekobanken", AccountType: AccountTypes[AccountType2]},
	{Interval: "9880..9889", BankName: "Riksgälden", AccountType: AccountTypes[AccountType2]},
	{Interval: "9890..9899", BankName: "Riksgälden", AccountType: AccountTypes[AccountType3], MinLength: 10},
	{Interval: "9960..9969", BankName: "Nordea/Plusgirot", AccountType: AccountTypes[AccountType5], MaxLength: 10, MinLength: 1},
}

var AccountTypes = map[string]AccountType{
	AccountType1: {Type: AccountType1, Algorithm: "mod11", ClearingLength: 4, SerialLength: 7, WeightedNumbers: 10},
	AccountType2: {Type: AccountType2, Algorithm: "mod11", ClearingLength: 4, SerialLength: 7, WeightedNumbers: 11},
	AccountType3: {Type: AccountType3, Algorithm: "mod10", ClearingLength: 4, SerialLength: 10, WeightedNumbers: 10},
	AccountType4: {Type: AccountType4, Algorithm: "mod11", ClearingLength: 4, SerialLength: 9, WeightedNumbers: 9},
	AccountType5: {Type: AccountType5, Algorithm: "mod10", ClearingLength: 5, SerialLength: 10, WeightedNumbers: 10},
}
