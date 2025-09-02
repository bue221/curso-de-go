package greatings

import "testing"

func TestSayHi(t *testing.T) {
	message, err := SayHi("Andres")
	if err != nil {
		t.Errorf("Expected 'Hi Andres', got '%s'", message)
	}
	if message != "Hi Andres" {
		t.Errorf("Expected 'Hi Andres', got '%s'", message)
	}
}

func TestSayHiEmptyName(t *testing.T) {
	message, err := SayHi("")
	if err == nil {
		t.Errorf("Expected error, got '%s'", message)
	}
}
