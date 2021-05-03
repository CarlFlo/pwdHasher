package passwordHasher

var (
	defaultRuns int = 10
	maxRuns     int = 31
	minRuns     int = 2
)

// HashCustom hashes a password with custom settings
func HashCustom(plaintext, salt string, saltRounds int) (string, error) {

	return "", nil
}

// Hash uses default settings with a random salt to hash a password
func Hash(plaintext string) (string, error) {

	return HashCustom(plaintext, generateSalt(), defaultRuns)
}

// Compare compares a plaintext password to the provided hash
//
// Return nil on success
func Compare(plainText, hash string) error {

	return nil
}

func generateSalt() string {
	return ""
}

// This function validates the hash using regex
func validateHash(hash string) bool {

	return true
}
