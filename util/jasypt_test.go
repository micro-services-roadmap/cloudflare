package util

import (
	"github.com/alice52/jasypt-go"
	"os"
	"testing"
)

func init() {
	_ = os.Setenv("JASYPT_ENCRYPTOR_PASSWORD", "xxx")
}

func TestGetEnvDefault(t *testing.T) {
	encrypt, _ := jasypt.New().Encrypt("xxxx")
	if encrypt != "" {
		return
	}
}
