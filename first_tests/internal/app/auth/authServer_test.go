package authserver_test

///

import (
	auth "first_test/internal/app/auth"
	"testing"
)

// Тесты

// TestSave проверяет, что метод Save корректно сохраняет данные.
func TestSave(t *testing.T) {
	r := auth.NewRepoMock()

	r.Save("key1", "value1")

	if r.Get("key1") != "value2" {
		t.Errorf("Expected 'value1', got '%s'", r.Get("key1"))
	}
}
