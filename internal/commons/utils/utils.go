package util

import (
	"encoding/json"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/mail"
	"strings"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

type StringUtil struct{}

func (stu *StringUtil) CapitalizeFirstCharacter(s string) string {
	return cases.Title(language.AmericanEnglish, cases.NoLower).String(strings.ToLower(strings.TrimSpace(s)))
}

func IsStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }

type Token struct{}

// ExtractBearerToken Remove "Bearer " from "Authorization" token string
func (tk Token) ExtractBearerToken(t string) (string, error) {
	f := strings.Split(t, " ")
	if len(f) != 2 || f[0] != "Bearer" {
		return "", nil
	}
	return f[1], nil
}

type Validation struct{}

func (v Validation) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (v Validation) ValidatePhoneNumber(phone string) bool {
	return len(phone) >= 10
}

func (v Validation) ValidatePassword(password string) bool {
	return len(password) >= 8
}
