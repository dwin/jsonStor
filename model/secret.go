package model

import (
	"log"
	"time"

	"github.com/ammario/isokey"
	"github.com/gin-gonic/gin"
)

var (
	ks = isokey.NewSymKeyService([]byte("super_secure111"))
)

func CreateKey(userID uint32) (secret string, err error) {
	key := &isokey.Key{
		UserID:    userID,
		ExpiresAt: time.Now().AddDate(1, 0, 0), // Expires in One Year
	}
	digest, err := ks.Sign(key)
	if err != nil {
		log.Fatalf("Failed to sign key: %v", err)
	}
	return digest, err
}

func ValidateKey() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Check API Not Invalidated
		c.Next()
	}
}

func InvalidateKey(userID uint32) {

}
