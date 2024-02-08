package bankverification

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// bank account length consts.
const (
	DEFAULT_BANK_ACCOUNT_MIN_LENGTH = 7
	DEFAULT_BANK_ACCOUNT_MAX_LENGTH = 7
)

// BankAccount represents a bank account for validation.
type BankAccount struct {
	value           string
	clearingNumbers []ClearingNumber
}

// NewBankAccount initialises a new bank account.
func NewBankAccount(number string) BankAccount {
	return BankAccount{
		value:           number,
		clearingNumbers: defaultClearingNumbers,
	}
}

// SetClearingNumbers overrides the default clearing numbers with the provided ones.
func (b *BankAccount) SetClearingNumbers(clearingNumbers []ClearingNumber) {
	b.clearingNumbers = clearingNumbers
}

// IsValid checks if a bank account is valid.
func (b BankAccount) IsValid() bool {
	return len(b.ValidationErrors()) == 0
}

// ValidationErrors runs validations and return validation errors.
func (b BankAccount) ValidationErrors() []error {
	validationErrors := []error{}

	minSNLength, maxSNLength := b.getSerialNumberLength()
	serialNumber := b.GetSerialNumber()

	if len(serialNumber) < minSNLength {
		validationErrors = append(validationErrors, errors.New(ERR_SERIAL_TOO_SHORT))
	}

	if len(serialNumber) > maxSNLength {
		validationErrors = append(validationErrors, errors.New(ERR_SERIAL_TOO_LONG))
	}

	// Use a regular expression to check for invalid characters.
	invalidChars := regexp.MustCompile(`[^\d -.]`)
	if invalidChars.MatchString(b.GetAccountNumber()) {
		validationErrors = append(validationErrors, errors.New(ERR_INVALID_CHARACTER_ACCOUNT))
	}

	if !b.isValidSerial() {
		validationErrors = append(validationErrors, errors.New(ERR_INVALID_CHECKSUM))
	}

	if bankName, err := b.GetBank(); err != nil || bankName == "" {
		validationErrors = append(validationErrors, errors.New(ERR_UNKNOWN_BANK_DATA))
	}

	return validationErrors
}

// GetBankData finds a clearing number for the bank account.
func (b BankAccount) GetBankData() (ClearingNumber, error) {
	clearingNumber, _ := strconv.Atoi(b.value[0:4])

	for i := range b.clearingNumbers {
		interval := strings.Split(b.clearingNumbers[i].Interval, "..")

		if len(interval) < 2 {
			return ClearingNumber{}, errors.New("could not get lower and upper value from interval")
		}

		lower, _ := strconv.Atoi(interval[0])
		upper, _ := strconv.Atoi(interval[1])

		if lower <= clearingNumber && clearingNumber <= upper {
			return b.clearingNumbers[i], nil
		}
	}

	return ClearingNumber{}, errors.New("could not find clearing number")
}

// GetAccountNumber returns the account number.
func (b BankAccount) GetAccountNumber() string {
	return b.value
}

// GetBank returns the bank associated to the bank number.
func (b BankAccount) GetBank() (string, error) {
	bankData, err := b.GetBankData()
	if err != nil {
		return "", err
	}

	return bankData.BankName, nil
}

// GetClearingNumber calculates and returns the clearing number for a bank account.
func (b BankAccount) GetClearingNumber() (string, error) {
	digits := b.digits()
	parts := []string{digits[:4]}

	if b.getChecksumForClearing() {
		parts = append(parts, digits[4:5])
	}

	filteredParts := filter(parts, func(e string) bool {
		return e != ""
	})

	return strings.Join(filteredParts, "-"), nil
}

// GetSerialNumber calculates and returns the serial number for a bank account.
func (b BankAccount) GetSerialNumber() string {
	number := b.digits()[b.getClearingNumberLength():]

	if len(number) == 0 {
		return number
	}

	if b.isZeroFill() {
		minLength, _ := b.getSerialNumberLength()

		return padLeft(number, minLength, '0')
	}

	return number
}

// isValidSerial checks if a serial is valid.
func (b BankAccount) isValidSerial() bool {
	bankData, err := b.GetBankData()
	if err != nil {
		return false
	}

	if bankData.AccountType.Algorithm == "" {
		return true
	}

	if bankData.AccountType.Algorithm == AlgorithmMod10 {
		return mod10(b.getValidationNumbers(bankData))
	}

	if bankData.AccountType.Algorithm == AlgorithmMod11 {
		return mod11(b.getValidationNumbers(bankData))
	}

	return true
}

// getValidationNumbers returns numbers used for validation.
func (b BankAccount) getValidationNumbers(bankData ClearingNumber) string {
	length := bankData.AccountType.WeightedNumbers
	accountType := bankData.AccountType.Type
	var numbers string

	if accountType == AccountType1 || accountType == AccountType2 {
		d := b.digits()
		if len(d) > 0 {
			numbers = d[len(d)-length:]
		}
	}

	if accountType == AccountType3 || accountType == AccountType4 || accountType == AccountType5 {
		sn := b.GetSerialNumber()
		if len(sn) > 0 {
			numbers = sn[len(sn)-length:]
		}
	}

	return numbers
}

// getSerialNumberLength returns the min and max length for a serial number
// depending on the account number.
func (b BankAccount) getSerialNumberLength() (int, int) {
	bankData, err := b.GetBankData()
	if err != nil {
		return DEFAULT_BANK_ACCOUNT_MIN_LENGTH, DEFAULT_BANK_ACCOUNT_MAX_LENGTH
	}

	if bankData.MinLength == 0 {
		bankData.MinLength = DEFAULT_BANK_ACCOUNT_MIN_LENGTH
	}

	if bankData.MaxLength == 0 {
		bankData.MaxLength = bankData.MinLength
	}

	return bankData.MinLength, bankData.MaxLength
}

// isZeroFill returns if a serial number should be zero filled.
func (b BankAccount) isZeroFill() bool {
	bankData, err := b.GetBankData()
	if err != nil {
		return false
	}

	return bankData.Zerofill
}

// getChecksumForClearing validates checksum for clearing number.
func (b BankAccount) getChecksumForClearing() bool {
	bankData, err := b.GetBankData()
	if err != nil {
		return false
	}

	if bankData.ChecksumForClearing {
		if bankData.BankName != "Swedbank" {
			return true
		}

		if bankData.AccountType.Algorithm == AlgorithmMod10 {
			return mod10(b.digits()[5:])
		} else if bankData.AccountType.Algorithm == AlgorithmMod11 {
			return mod11(b.digits()[5:])
		}
	}

	return false
}

// digits normalize bank account.
func (b BankAccount) digits() string {
	re := regexp.MustCompile(`\D`)

	return re.ReplaceAllString(b.value, "")
}

// getClearingNumberLength calculates clearing number length based on the checksum.
func (b BankAccount) getClearingNumberLength() int {
	if b.getChecksumForClearing() {
		return 5
	}

	return 4
}
