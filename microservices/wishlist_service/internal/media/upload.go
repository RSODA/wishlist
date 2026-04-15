package media

import (
	"encoding/base64"
	"log"
	"os"
	"time"
)

func Upload(file string) (string, error) {
	decodedFile, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		log.Println("base64 decode error:", err)
		return "", err
	}

	filename := time.Now().String()

	err = os.WriteFile("assets/img/"+filename+".jpg", decodedFile, 0666)

	if err != nil {
		log.Println("file write error:", err)
		return "", err
	}

	return filename, nil
}
