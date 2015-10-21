package cryptpass

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"
)

const (
	ikey = "WueIirKvsQpSc6x3ZSHd5g=="
	iiv  = "1PNr7RSgUy2ITtD/iEJGOg=="
)

var (
	PassPath          = "/etc/cryptpass.key"
	ErrLengthNotMatch = errors.New("length not match")
)

var (
	masterKey    []byte
	masterIV     []byte
	cachedPasswd map[string]string
)

func init() {
	cachedPasswd = make(map[string]string)
}

func xorBytes(b1 []byte, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, ErrLengthNotMatch
	}
	buf := make([]byte, len(b1))
	for i := 0; i < len(b1); i++ {
		buf[i] = b1[i] ^ b2[i]
	}
	return buf, nil
}

func getBytes(reader *bufio.Reader, internal string) ([]byte, error) {
	ibyte, err := base64.StdEncoding.DecodeString(internal)
	if err != nil {
		return nil, err
	}

	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	fbyte, err := base64.StdEncoding.DecodeString(line)
	if err != nil {
		return nil, err
	}

	tbyte, err := xorBytes(ibyte, fbyte)
	if err != nil {
		return nil, err
	}
	return tbyte, nil
}

func readKeyIV() error {
	file, err := os.Open(PassPath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	masterKey, err = getBytes(reader, ikey)
	if err != nil {
		return err
	}

	masterIV, err = getBytes(reader, iiv)
	if err != nil {
		return err
	}

	return nil
}

func EncryptPass(s string) (string, error) {
	if masterKey == nil || masterIV == nil {
		err := readKeyIV()
		if err != nil {
			return "", err
		}
	}

	c, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(c, masterIV)

	buf := make([]byte, len(s))
	stream.XORKeyStream(buf, []byte(s))

	return base64.StdEncoding.EncodeToString(buf), nil
}

func DecryptPass(s string) (string, error) {
	if masterKey == nil || masterIV == nil {
		err := readKeyIV()
		if err != nil {
			return "", err
		}
	}

	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", err
	}
	stream := cipher.NewCFBDecrypter(c, masterIV)

	buf := make([]byte, len(src))
	stream.XORKeyStream(buf, src)
	return string(buf), nil
}

// AutoPass will try to decrypt s to real password.
// it use cache to speed up decrypt.
// CAUTION: if it can't, it will return original string.
// and original string will not set to cache.
func AutoPass(s string) string {
	r, ok := cachedPasswd[s]
	if ok {
		return r
	}

	r, err := DecryptPass(s)
	if err != nil {
		return s
	}

	cachedPasswd[s] = r
	return r
}

func SafePass(s string) string {
	if s[:3] != ".[~" {
		return s
	}
	k := s[3:]

	r, ok := cachedPasswd[k]
	if ok {
		return r
	}

	r, err := DecryptPass(k)
	if err != nil {
		return s
	}

	cachedPasswd[k] = r
	return r
}
