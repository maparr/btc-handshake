package checksum

import (
	"testing"
)

func TestCreateChecksum(t *testing.T) {
	payload := []byte("test payload")
	expected := [4]byte{0x18, 0x1e, 0x61, 0x99}
	result := CreateChecksum(payload)
	if result != expected {
		t.Errorf("Expected %x, got %x", expected, result)
	}
}
