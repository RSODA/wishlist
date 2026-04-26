package media

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Upload(file string) (string, error) {
	base64String := strings.Split(file, ",")

	file = base64String[1]
	format := strings.Split(base64String[0], "/")
	format = strings.Split(format[1], ";")

	decodedFile, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		log.Println("base64 decode error:", err)
		return "", err
	}

	fmt.Println(format)

	filename := time.Now().Format("20060102150405")

	err = os.WriteFile("assets/img/"+filename+"."+format[0], decodedFile, 0666)

	if err != nil {
		log.Println("file write error:", err)
		return "", err
	}

	return filename + "." + format[0], nil
}
