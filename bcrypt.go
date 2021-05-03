package bcrypt

// HashCustom hashes a password with custom settings
func HashCustom(plaintext, salt string, saltRounds int) (string, error) {

	return "", nil
}

// Hash uses default settings with a random salt to hash a password
func Hash(plaintext string) (string, error) {

	return "", nil
}

// Compare compares a plaintext password to the provided hash
func Compare(plainText, hash string) (bool, error) {

	return true, nil
}

// This function validates the hash using regex
func validateHash(hash string) bool {

	return true
}
