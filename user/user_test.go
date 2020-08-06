package user

import (
	"log"
	"testing"
)

func TestNewCredential(t *testing.T) {
	c0, err := NewCredentials("user1", "AbC123")
	if err != nil {
		t.Error("Failed to create a credential c0\n", err.Error())
	}
	c1, err := NewCredentials("user1", "AbC123")
	if err != nil {
		t.Error("Failed to create a credential c1\n", err.Error())
	}

	if string(c0.Salt) == string(c1.Salt) {
		t.Errorf("Failed to create different salts for c0 and c1:\nc0: %s\nc1:%s", c0.Salt, c1.Salt)
	}
	if string(c0.Hash) == string(c1.Hash) {
		t.Errorf("Failed to create different hashs for c0 and c1:\nc0: %s\nc1:%s", c0.Hash, c1.Hash)
	}
}

func TestSalt(t *testing.T) {
	s0 := salt()
	log.Println("s0:", s0)
	s1 := salt()
	log.Println("s1:", s1)

	if string(s0) == string(s1) {
		t.Error("s0 equals s1")
	}
}

func TestHash(t *testing.T) {
	h0, err := hash("ABC123", []byte("DEF456"))
	if err != nil {
		t.Error(err)
	}
	h1, err := hash("ABC123", []byte("DEF456"))
	if err != nil {
		t.Error(err)
	}
	if string(h0) != string(h1) {
		t.Errorf("h0 must be equals h1")
	}
}
