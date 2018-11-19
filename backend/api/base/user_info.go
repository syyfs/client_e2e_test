package base

import (
	"fmt"

	"crypto/aes"
	"crypto/cipher"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"brilliance/client_e2e_test/blockchain/common/util"
)

type CryptoData struct {
	EncryptedData string `json:"encryptedData"`
	SessionKey    string `json:"sessionKey"`
	Iv            string `json:"iv"`
}

func UploadUnionId(c *gin.Context) {
	body, err := util.ProcessBody(c)
	if err != nil {
		fmt.Println("error is : ", err.Error())
	}

	fmt.Println("body is : ", string(body))

	cryptoData := CryptoData{}
	json.Unmarshal(body, &cryptoData)
	fmt.Println("cryptoData is : ", cryptoData)
	eb, _ := json.Marshal(cryptoData.EncryptedData)
	sb, _ := json.Marshal(cryptoData.SessionKey)
	ib, _ := json.Marshal(cryptoData.Iv)
	unionId, err := getOriginalInformation(eb, sb, ib)
	if err != nil {
		fmt.Println("error is : ", err.Error())

	}
	fmt.Println("unionid is : ", unionId)
}

func getOriginalInformation(encryptedData []byte, sessionKey []byte, iv []byte) (string, error) {
	var aesBlockDecrypter cipher.Block
	aesBlockDecrypter, err := aes.NewCipher(sessionKey)
	if err != nil {
		return "", err
	}
	decrypted := make([]byte, len(encryptedData))
	aesDecrypter := cipher.NewCBCDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.CryptBlocks(decrypted, encryptedData)

	return string(decrypted), nil
}
