package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"net/mail"
	"strings"
	"time"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func MonthYearSort(month time.Month, year int) (time.Time, time.Time) {
	currentLocation := time.Now().Location()

	first := time.Date(year, month, 1, 0, 0, 0, 0, currentLocation)
	last := first.AddDate(0, 1, -1)
	//lastOfMonth := first.AddDate(0, 1, -1).Format("2006-01-02")
	return first, last
}

func isNewMonth(month time.Month, year int) bool {
	//thisMonth := time.Now().Month()
	//thisYear := time.Now().Year()
	return true
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

func GetUUID() string {
	return uuid.NewString()
}
