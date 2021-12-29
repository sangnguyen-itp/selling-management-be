package excel

import "path/filepath"

func ValidateExt(fileName string) (string, bool) {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".xlsx":
		return "xlsx", true
	case ".xls":
		return "xls", true
	default:
		return "", false
	}
}
