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
	expectedHash := "m2jtckW-PFTd?10?ccce082519ed081bc2f6926cc924fa9ffdda96bd4a9be9217b6191d5518c58d9"
	hash, err := HashCustom(plainText, "m2jtckW-PFTd", 10)

	if err != nil {
		t.Fatalf("Failed to hash the plaintext: %s", err)
	}

	if hash != expectedHash {
		t.Fatalf("The hashes did not match. Got: %s", hash)
	}
}

func TestCompare(t *testing.T) {

	plainText := "There is a donkey in the salt"
	hash := "donkey?14?a69dcc315a693f8bd643f563b53472965595176a111f49a89e3f1db91322bc3d"

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
