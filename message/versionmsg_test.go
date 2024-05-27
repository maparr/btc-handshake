package message

import (
	"testing"
)

func TestEncodeDecodeVersionMessage(t *testing.T) {
	original := CreateVersionMessage("127.0.0.1")
	encoded, err := EncodeVersionMessage(original)
	if err != nil {
		t.Fatalf("Failed to encode version message: %v", err)
	}

	decoded, err := DecodeVersionMessage(encoded)
	if err != nil {
		t.Fatalf("Failed to decode version message: %v", err)
	}

	if original.Version != decoded.Version ||
		original.Services != decoded.Services ||
		original.Timestamp != decoded.Timestamp ||
		original.Nonce != decoded.Nonce ||
		original.UserAgent != decoded.UserAgent ||
		original.StartHeight != decoded.StartHeight ||
		original.Relay != decoded.Relay {
		t.Errorf("Mismatch between original and decoded version message")
	}
}
