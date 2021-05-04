package pwdHasher

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var (
	defaultRuns       int = 10
	maxRuns           int = 256
	minRuns           int = 2
	defaultSaltLength int = 12
	maxSaltLength     int = 16
)

// HashCustom hashes a password with custom settings
func HashCustom(plaintext, salt string, rounds int) (string, error) {

	if rounds < minRuns || rounds > maxRuns {
		return "", fmt.Errorf("number of rounds is greater than %d or smaller than %d", maxRuns, minRuns)
	}

	if len(salt) > maxSaltLength {
		return "", fmt.Errorf("the salt is invalid. It must be less than 16 in length")
	}

	if len(salt) == 0 {
		var err error
		salt, err = generateSalt(defaultSaltLength)
		if err != nil {
			return "", err
		}
	}

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", plaintext, salt)))

	for i := 1; i < rounds; i++ {
		h.Write([]byte(fmt.Sprintf("%x", h.Sum(nil))))
	}

	return fmt.Sprintf("%s?%d?%x", salt, rounds, h.Sum(nil)), nil
}

// Hash uses default settings with a random salt to hash a password
func Hash(plaintext string) (string, error) {

	salt, err := generateSalt(defaultSaltLength)
	if err != nil {
		return "", err
	}
	return HashCustom(plaintext, salt, defaultRuns)
}

// Compare compares a plaintext password to the provided hash
//
// Return nil on success
func Compare(plainText, hash string) error {

	split := strings.Split(hash, "?")
	if len(split) < 3 {
		return fmt.Errorf("invalid hash format. Expecting: <salt>?<rounds>?<sha256>")
	}

	salt := split[0]
	rounds, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return err
	}

	toCompare, err := HashCustom(plainText, salt, int(rounds))
	if err != nil {
		return err
	}

	if toCompare != hash {
		return fmt.Errorf("hashes do not match")
	}

	return nil
}

func generateSalt(length int) (string, error) {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}
		ret[i] = alphabet[num.Int64()]
	}

	return string(ret), nil
}
