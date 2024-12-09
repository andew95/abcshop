package abcToken_test

import (
	"abcShop/pkg/abcToken"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	service := abcToken.NewToken("secret", "http:localhost:8080")
	t.Run("Generate Token Success", func(t *testing.T) {
		token, err := service.GenerateToken("test", "user@example.com", "member")
		if err != nil {
			t.Error(err)
		}
		fmt.Println(token)
	})
}
