package cryptogen

import (
	"os"
	"path/filepath"
)

func GenerateCaDir() (cakeystorepath, tlscakeystory string, err error) {
	dir, _ := filepath.Abs("./")
	tempDir := filepath.Join(dir, "temp")
	folders := []string{
		filepath.Join(tempDir, "cakeystory"),
		filepath.Join(tempDir, "tlscakeystory"),
	}
	for _, folder := range folders {

		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return "", "", err
		}
	}
	cakeystorepath, err = filepath.Abs("./temp/cakeystory")
	if err != nil {
		return "", "", err
	}
	tlscakeystory, err = filepath.Abs("./temp/tlscakeystory")
	if err != nil {
		return "", "", err
	}

	return cakeystorepath, tlscakeystory, err
}

func GenerateNodeDir() (string, string, error) {
	dir, _ := filepath.Abs("./")
	tempDir := filepath.Join(dir, "temp")
	folders := []string{

		filepath.Join(tempDir, "nodekeystory"),
		filepath.Join(tempDir, "nodetlskeystory"),
	}
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return "", "", err
		}
	}
	nodekeystory, err := filepath.Abs("./temp/nodekeystory")
	if err != nil {
		return "", "", err
	}
	nodetlskeystory, err := filepath.Abs("./temp/nodetlskeystory")
	if err != nil {
		return "", "", err
	}
	return nodekeystory, nodetlskeystory, err
}

