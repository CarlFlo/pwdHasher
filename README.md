# Password Hasher

A user friendly module for hashing passwords and then verifying them.

The module hashes a plaintext message + salt and then hashes it multible times in rounds to make it computationally infeasible to bruteforce it to uncover the plaintext message.

Test coverage: **87.2%**

## Features

- [X] Create hashes
- [X] Comparing

## Install

```
go get github.com/CarlFlo/pwdHasher
```

## Usage

> This code uses SHA256 (a fast hash) and is not suitable for anything that requires protection, the code is also not tested to verify that it can withstand attacks such as timing attacks. For encrypting passwords so should a better solution such as [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) be used. This code was made for fun as a short hobby project.

Anyways...

Hashes created by this program will follow this format **\<salt>?\<rounds>?\<sha256>**.

Example: **m2jtckW-PFTd?10?8cc4225f1dc9bf97a63829fc4fef407cdf897fad92f524172c4e200b9c6353a8**

This means that with a hash and a plaintext message can a hash be comapred against a plaintext message.

```go
// Hash a message (10 rounds) with a random salt
hash, err := pwdHasher.Hash("password123")

// Here is how it can be done with custom settings
rounds := 10 // Min: 2, max: 256
salt := "salt" // Maximum 16 characters, random if length is 0 -> ""
hash, err := pwdHasher.HashCustom("password123", salt, rounds)

// For comparing. If err is nil then it was a success
hash := "m2jtckW-PFTd?10?8cc4225f1dc9bf97a63829fc4fef407cdf897fad92f524172c4e200b9c6353a8"
err := pwdHasher.Compare("This is a test", hash)
```