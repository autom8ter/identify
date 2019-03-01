package util

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"regexp"
	"strings"
)

func LoadCertificate(crtFile string) (string, error) {
	crtByte, err := ioutil.ReadFile(crtFile)
	if err != nil {
		return "", err
	}
	crtString := string(crtByte)

	re := regexp.MustCompile("---(.*)CERTIFICATE(.*)---")
	crtString = re.ReplaceAllString(crtString, "")
	crtString = strings.Trim(crtString, " \n")
	crtString = strings.Replace(crtString, "\n", "", -1)

	return crtString, err
}

func ReadAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}
	stringReader := strings.NewReader(val)
	csvReader := csv.NewReader(stringReader)
	return csvReader.Read()
}

func RandomToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
func RandomTokenBytes() ([]byte, error) {
	tok, _ := RandomToken()
	return base64.StdEncoding.DecodeString(tok)
}

func UserPassword(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func Uuid() string {
	return fmt.Sprintf("%s", uuid.New())
}

func Replace(old, new, src string) string {
	return strings.Replace(src, old, new, -1)
}
