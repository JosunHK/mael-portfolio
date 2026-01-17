package secure

import (
    "context"
	log "github.com/sirupsen/logrus"
	"crypto/rand"
	"encoding/hex"
)

const NonceKey string = "NONCE_KEY"

func GenerateNonce(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}


func GetNonce(ctx context.Context) string {
	ctxNonce := ctx.Value(NonceKey)
	if ctxNonce == nil {
		log.Fatal("error getting nonce - is nil")
	}

    nonce, ok := ctxNonce.(string)
    if !ok || nonce == ""{
		log.Fatal("error getting nonce - not string")
    }


    log.Info("nonce " + nonce)
	return nonce
}
