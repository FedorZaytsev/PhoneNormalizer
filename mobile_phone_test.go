package msisdn

import (
	"testing"
)

func TestMobilePhone(t *testing.T) {
	ok := []struct {
		Test, Result string
	}{
		{"79264932152", "79264932152"},
		{"89264932152", "79264932152"},
		{"+79264932152", "79264932152"},
		{"+7 926 493 21 52", "79264932152"},
		{"+7 (926) 493 21 52", "79264932152"},
		{"+7(926)493-21-52", "79264932152"},
		{"+7-(926)-493-21-52", "79264932152"},
		{"+7     (92 6)-49   3-2(1)-52", "79264932152"},
	}
	errors := []string{
		"sadfasdf",
		"79264932152asdf",
		"9192913",
		"9264932152",
		"755559",
		"4957521212",
	}

	for _, test := range ok {
		msisdn, err := ParsePhoneNumber(test.Test)
		if err != nil {
			t.Errorf("err == %s", err)
		}
		if msisdn != test.Result {
			t.Errorf("Not equal. Test %s, expected %s, got %s", test.Test, test.Result, msisdn)
		}
		isMsisdn, err := IsPhoneNumber(test.Test)
		if err != nil || !isMsisdn {
			t.Errorf("MSISDN %s must be a phone number. Error %s", test.Test, err)
		}
	}
	for _, test := range errors {
		isMsisdn, _ := IsPhoneNumber(test)
		if isMsisdn {
			t.Errorf("Must be not a msisdn %s", test)
		}
		_, err := ParsePhoneNumber(test)
		if err == nil {
			t.Errorf("Must be error %s", test)
		}
	}
}
