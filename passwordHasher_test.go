package pwdHasher

import "testing"

func TestHash(t *testing.T) {
	_, err := Hash("plainText")
	if err != nil {
		t.Fatalf("Failed to hash the plaintext: %s", err)
	}
}

func TestHashCustom(t *testing.T) {

	plainText := "This is a test"
	expectedHash := "m2jtckW-PFTd?10?8cc4225f1dc9bf97a63829fc4fef407cdf897fad92f524172c4e200b9c6353a8"
	hash, err := HashCustom(plainText, "m2jtckW-PFTd", 10)

	if err != nil {
		t.Fatalf("Failed to hash the plaintext: %s", err)
	}

	if hash != expectedHash {
		t.Fatalf("The hashes did not match is: %s", hash)
	}
}

func TestCompare(t *testing.T) {

	plainText := "There is a donkey in the salt"
	hash := "donkey?14?e06f97499cd7e2b0a5be65b56c82918f212bde103655dca3a4c3bc85b897df75"

	err := Compare(plainText, hash)
	if err != nil {
		t.Fatalf("Comparing hashes failed!")
	}
}

func TestErrorHandling(t *testing.T) {

	var err error

	_, err = HashCustom("text", "salt", 1)
	if err == nil {
		t.Fatalf("A round of 1 should have failed")
	}

	_, err = HashCustom("text", "d465fAS7D6G8FNYBS987DFASD7BVHAISDGKJFHADSLKJGHVNLASDFGN", 10)
	if err == nil {
		t.Fatalf("A salt that is too long should have failed")
	}

	_, err = HashCustom("text", "", 10)
	if err != nil {
		t.Fatalf("A random salt should have been generated")
	}

	err = Compare("test", "?")
	if err == nil {
		t.Fatalf("The hash format was invalid and should have failed")
	}

	err = Compare("This plaintext is not for this hash", "donkey?14?e06f97499cd7e2b0a5be65b56c82918f212bde103655dca3a4c3bc85b897df75")
	if err == nil {
		t.Fatalf("The plaintext should not have matched this hash")
	}

}
