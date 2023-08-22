package utils

import (
	"fmt"
	"os"
)

func ConnectionURLBuilder(n string) (string, error) {
	var url string
	switch n {
	case "minio":
		// URL for MinIO connection.
		url = fmt.Sprintf(
			os.Getenv("MINIO_ACCESSKEY"),
			os.Getenv("MINIO_SECRETKEY"),
			os.Getenv("MINIO_HOST"),
			os.Getenv("MINIO_PORT"),
		)

	case "fiber":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}
	fmt.Println(url)
	return url, nil
}
