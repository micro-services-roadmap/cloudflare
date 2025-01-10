package util

import (
	"fmt"
	"github.com/alice52/jasypt-go"
	"os"
)

func GetEnv(key string) string {
	token := os.Getenv(key)
	if len(token) == 0 {
		return token
	}
	if decrypt, err := jasypt.New().DecryptWrapper(token); err == nil {
		token = decrypt
	} else {
		fmt.Printf("decrypt %s[%s] faled: %v\n", key, token, err)
	}

	return token
}

func GetEnvDefault(key, defaultVal string) string {
	token := GetEnv(key)
	if len(token) != 0 {
		return token
	}

	if decrypt, err := jasypt.New().DecryptWrapper(defaultVal); err == nil {
		return decrypt
	} else {
		fmt.Printf("decrypt %s[%s] faled: %v\n", key, defaultVal, err)
		return token
	}
}
