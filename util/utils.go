package util;
import (
	"os"
)

func GetWorkDirectory() (path string) {
	path, _ = os.Getwd()
	return path
}