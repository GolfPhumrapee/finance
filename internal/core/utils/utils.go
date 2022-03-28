package utils

import (
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"gitee.com/beshy/php2go"
	"github.com/GolfPhumrapee/finance/internal/core/context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

const (
	numberOfShortPostText = 100
)

// GenerateUniqueCode gengerate unique code
func GenerateUniqueCode() string {
	return time.Now().In(LoadLocation()).Format("20060102150405")
}

// GenerateSeqNo generate seq no
func GenerateSeqNo(n, length int) string {
	digits := CountDigits(n)
	if digits > length {
		return fmt.Sprintf("%d", n)
	}

	return fmt.Sprintf("%0*d", length, n)
}

// CountDigits count digits
func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

// GenerateNumber generate number
func GenerateNumber(number, length int) string {
	if number <= 0 {
		return ""
	}

	return fmt.Sprintf(fmt.Sprintf("%%0%dd", length), number)
}

// UniqueSliceString for remove duplicate from slice string
func UniqueSliceString(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			if entry != "" {
				list = append(list, entry)
			}
		}
	}

	return list
}

// UniqueSliceInt for remove duplicate from slice int
func UniqueSliceInt(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			if entry != 0 {
				list = append(list, entry)
			}
		}
	}

	return list
}

// UniqueSliceUInt for remove duplicate from slice uint
func UniqueSliceUInt(slice []uint) []uint {
	keys := make(map[uint]bool)
	list := []uint{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			if entry != 0 {
				list = append(list, entry)
			}
		}
	}

	return list
}

// IsSliceStringChanged is slice string change
func IsSliceStringChanged(original, compare []string) bool {
	if len(original) != len(compare) {
		return true
	}

	var count int
	for _, r := range compare {
		var i int
		for i <= len(original)-1 {
			if r == original[i] {
				count++
				break
			}
			i++
		}
	}

	return count != len(original)
}

// IsEmptySlice is empty slice
func IsEmptySlice(values interface{}) bool {
	if values, ok := values.([]int); ok {
		for _, value := range values {
			if value == 0 {
				return true
			}
		}
		return false
	}

	if values, ok := values.([]string); ok {
		for _, value := range values {
			if value == "" {
				return true
			}
		}
		return false
	}

	return true
}

// RoundUp is round up decimal position at 1 if value more than 0
func RoundUp(v float64) float64 {
	return float64(int((v*100)+9)/10) / 10.0
}

// WrapPassword for wrap password Example: ********
func WrapPassword(password string) string {
	var wrapPassword string
	lengthPassword := utf8.RuneCountInString(password)
	for i := 0; i < lengthPassword; i++ {
		wrapPassword = wrapPassword + "*"
	}

	return wrapPassword
}

// GetInitialName for get first character.
func GetInitialName(name string) string {
	runeName := []rune(name)
	if IsValidEnglishAlphabet(string(runeName[0])) {
		name = strings.ToUpper(string(runeName[0]))
		return name
	}

	return string(runeName[0])
}

// IsValidEmails is valid emails
func IsValidEmails(emails []string) bool {
	for _, email := range emails {
		if email == "" || !IsValidEmail(email) {
			return false
		}
	}

	return true
}

// TrimSpaces trim slice
func TrimSpaces(slice []string) []string {
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}

	return slice
}

// ToLower to lower
func ToLower(slice []string) []string {
	for i := range slice {
		slice[i] = strings.ToLower(slice[i])
	}

	return slice
}

// RoundFloat64 round float64 (100 = .2f)
func RoundFloat64(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}

// SubString sub string
func SubString(sourceText string, showDot bool) string {
	newString := sourceText
	if len([]rune(sourceText)) > numberOfShortPostText {
		dotString := ""
		if showDot {
			dotString = "..."
		}
		newString = fmt.Sprintf("%s%s", string([]rune(sourceText)[:numberOfShortPostText]), dotString)
	}

	return newString
}

// ConvertToArrayInterface convert to array interface...
func ConvertToArrayInterface(x interface{}) []interface{} {
	result := []interface{}{}
	switch reflect.TypeOf(x).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(x)
		for i := 0; i < s.Len(); i++ {
			result = append(result, s.Index(i).Interface())
		}
	}

	return result
}

// isValidCitizenID is valid citizenID
func isValidCitizenID(citizenID string) bool {
	if !regexIDCard.MatchString(citizenID) {
		return false
	}

	sum, row := 0, 13
	citizenIDRune := []rune(citizenID)
	for _, n := range string(citizenIDRune) {
		number, _ := strconv.Atoi(string(n))
		sum += number * row
		row--

		if row == 1 {
			break
		}
	}

	citizenIDInt, _ := strconv.Atoi(citizenID)
	result := (11 - (int(sum) % 11)) % 10

	return (citizenIDInt % 10) == result
}

// MD5Hash md5 hash
func MD5Hash(text string) string {
	hasher := md5.New()
	_, _ = hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// FindNumberFromText find number from string
func FindNumberFromText(text string) int {
	slice := regexp.MustCompile("[0-9]+").FindAllString(text, -1)
	if len(slice) == 0 {
		return 0
	}

	number, _ := strconv.Atoi(strings.Join(slice, ""))
	return number
}

// RemoveLastCharacter remove last character from string
func RemoveLastCharacter(text, char string) string {
	data := []rune(text)
	length := len([]rune(char))
	if len(data) > length && string(data[len(data)-length:]) == char {
		return string(data[:len(data)-length])
	}

	return text
}

var (
	// SignKey is private key
	SignKey = &ecdsa.PrivateKey{}
	// VerifyKey is public key
	VerifyKey = &ecdsa.PublicKey{}
)

// ReadECDSAKey read private key && public key
func ReadECDSAKey(privateKey, publicKey string) error {
	privateKeyByte, err := os.ReadFile(privateKey)
	if err != nil {
		return err
	}

	publicKeyByte, err := os.ReadFile(publicKey)
	if err != nil {
		return err
	}

	SignKey, err = jwt.ParseECPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return err
	}

	VerifyKey, err = jwt.ParseECPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		return err
	}

	return nil
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Get preferred outbound ip of this machine
func GetIpAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func InArrayInt(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func InArrayString(value string, list []string) bool {
	for _, data := range list {
		if data == value {
			return true
		}
	}
	return false
}

func ConvertDateTime(dateTime, position string) string {
	if dateTime != "" {
		t := php2go.Explode(" ", dateTime)
		if len(t) == 0 {
			return dateTime
		}
		if position == "0" {
			return t[0]
		} else {
			return t[0] + " " + t[1]
		}
	} else {
		return ""
	}
}

func CheckValiDateValue(c *fiber.Ctx, request interface{}) error {
	ctx := context.New(c)
	err := ctx.BindValue(request, true)
	if err != nil {
		logrus.Errorf("bind value error: %s", err)
		return err
	}
	return nil
}
