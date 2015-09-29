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
	ikey     = "WueIirKvsQpSc6x3ZSHd5g=="
	iiv      = "1PNr7RSgUy2ITtD/iEJGOg=="
	passPath = "/etc/cryptpass.key"
)

var ErrLengthNotMatch = errors.New("length not match")

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

func readKeyIV() ([]byte, []byte, error) {
	file, err := os.Open(passPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	key, err := getBytes(reader, ikey)
	if err != nil {
		return nil, nil, err
	}

	iv, err := getBytes(reader, iiv)
	if err != nil {
		return nil, nil, err
	}

	return key, iv, nil
}

func EncryptPass(s string) (string, error) {
	key, iv, err := readKeyIV()
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(c, iv)

	buf := make([]byte, len(s))
	stream.XORKeyStream(buf, []byte(s))

	return base64.StdEncoding.EncodeToString(buf), nil
}

func DecryptPass(s string) (string, error) {
	key, iv, err := readKeyIV()
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBDecrypter(c, iv)

	buf := make([]byte, len(s))
	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	stream.XORKeyStream(buf, src)

	return string(buf), nil
}
