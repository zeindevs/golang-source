package util

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gliderlabs/ssh"
)

var (
	maxFileSize = 3 << 10
)

func parseFileHeader(r *bufio.Reader) (*FileHeader, error) {
	raw, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	headerParts := strings.Split(raw, " ")
	if len(headerParts) != 3 {
		return nil, fmt.Errorf("invalid header")
	}
	filesizeStr := headerParts[1]
	filesize, _ := strconv.Atoi(filesizeStr)
	if filesize > maxFileSize {
		return nil, fmt.Errorf("max filesize exceeded: %d > %d", filesize, maxFileSize)
	}
	return &FileHeader{
		FileSize: int64(filesize),
	}, nil
}

func GenerateRandomString(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type FileHeader struct {
	FileSize int64
	Filename string
}

func Getenv(env, def string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return def
	}
	return val
}

func IsEnvDev() bool {
	appEnv := Getenv("SENDIT_ENV", "")
	return appEnv == "dev"
}

func FingerprintSSHKey(key string) string {
	hash := sha256.Sum256([]byte(key))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func ParsePublickeyBytes(key string) ([]byte, error) {
	ekey, err := extractKeyData(key)
	if err != nil {
		return nil, err
	}
	decodedKey, err := base64.StdEncoding.DecodeString(ekey)
	if err != nil {
		return nil, err
	}
	parsedKey, err := ssh.ParsePublicKey(decodedKey)
	if err != nil {
		return nil, err
	}
	return parsedKey.Marshal(), nil
}

func IsValidSSHKey(key string) bool {
	data, err := extractKeyData(key)
	if err != nil {
		fmt.Println(err)
		return false
	}
	decodedKey, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = ssh.ParsePublicKey(decodedKey)
	if err != nil {
		fmt.Print(err)
		return false
	}
	return true
}

func extractKeyData(publicKey string) (string, error) {
	parts := strings.Fields(publicKey)
	if len(parts) != 3 || parts[0] != "ssh-rsa" {
		return "", fmt.Errorf("invalid public key format")
	}
	decodedKey, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	block, _ := pem.Decode(decodedKey)
	if block != nil {
		return string(block.Bytes), nil
	}
	return parts[1], nil
}

func IsValidSubdomain(domain string) bool {
	if len(domain) < 2 {
		return false
	}
	list := "sendit,streamit,about,abuse,access,account,accounts,acme,activate,activities,activity,ad,wws,www,wwws,wwww,xfn,xhtml,xhtrnl,xml,xmpp,xpg,xxx,yaml,year,you,yourdomain,yourname,yoursite"
	restricted := strings.Split(list, ",")
	for _, r := range restricted {
		if domain == r {
			return false
		}
	}
	pattern := "^[a-zA-Z0-9-]+$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(domain)
}

func MakeFullLink(link string) string {
	if IsEnvDev() {
		port := Getenv("SENDIT_HTTP_PORT", ":3000")
		return fmt.Sprintf("http://localhost%s/%s", port, link)
	}
	return fmt.Sprintf("https://sendit.io/%s", link)
}

func IsValidLink(link string) bool {
	return true
}

func ValidateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)

}
