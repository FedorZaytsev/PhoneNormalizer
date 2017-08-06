package msisdn

import (
	"bytes"
	"fmt"
	"regexp"
	"unicode"
)

const msisdnRegexp = "^((8|\\+7|7)[\\- ]?)(\\(?\\d{3}\\)?[\\- ]?)?[\\d\\- ]{7,10}$"

var NotMsisdn = fmt.Errorf("Not a msisdn")

type PhoneChecker struct {
	reg *regexp.Regexp
}

func (pc *PhoneChecker) IsPhoneNumber(msisdn string) bool {
	return pc.reg.MatchString(msisdn)
}

func (pc *PhoneChecker) ParsePhoneNumber(msisdn string) (string, error) {
	var buffer bytes.Buffer

	for _, ch := range msisdn {
		if unicode.IsSpace(ch) || ch == '+' || ch == '-' || ch == '(' || ch == ')' {
			continue
		}
		buffer.Write([]byte(string(ch)))
	}
	msisdn = buffer.String()
	if msisdn[0] == '8' {
		msisdn = "7" + msisdn[1:]
	}

	if !pc.IsPhoneNumber(msisdn) {
		return "", NotMsisdn
	}

	return msisdn, nil
}

func IsPhoneNumber(msisdn string) (bool, error) {
	ph, err := NewPhoneNumberParser()
	if err != nil {
		return false, fmt.Errorf("Cannot create phone number object. Reason %s", err)
	}

	_, err = ph.ParsePhoneNumber(msisdn)
	if err == NotMsisdn {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}

func ParsePhoneNumber(msisdn string) (string, error) {
	ph, err := NewPhoneNumberParser()
	if err != nil {
		return "", fmt.Errorf("Cannot create phone number object. Reason %s", err)
	}
	msisdn, err = ph.ParsePhoneNumber(msisdn)
	return msisdn, err
}

func NewPhoneNumberParser() (*PhoneChecker, error) {
	reg, err := regexp.Compile(msisdnRegexp)
	if err != nil {
		return nil, fmt.Errorf("Cannot compile regular expression. Reason %s", reg)
	}

	return &PhoneChecker{
		reg: reg,
	}, nil
}
