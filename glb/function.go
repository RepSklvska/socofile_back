package glb
// Global functions

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Method(jsonBody []byte) string {
	//return strings.Contains(
	//	string(json),
	//	fmt.Sprintf(`"method":"%s"`, method),
	//)
	jsonData := struct {
		Method string `json:"method"`
	}{}
	json.Unmarshal(jsonBody, &jsonData)
	return jsonData.Method
}

func GetField(jsonBody []byte, key string) string {
	jsonData := make(map[string]string)
	json.Unmarshal(jsonBody, &jsonData)
	return jsonData[key]
}

func MD5Sum(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func ReplaceRept(str *string, target string) { //Replace specified Repeated string to Single
	*str = strings.ReplaceAll(*str, target+target, target)
	if strings.Contains(*str, target+target) {
		ReplaceRept(str, target)
	}
}

func DirExist(dirname string) bool {
	file, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	if file.IsDir() {
		return true
	} else {
		return false
	}
}

func FileExist(filename string) bool {
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if file.IsDir() {
		return false
	} else {
		return true
	}
}

func ComputeSize(size int64) string {
	var s int = int(size)
	if s == 0 {
		return "0B"
	}
	if s < 1024 {
		
		return strconv.Itoa(s) + "B"
	}
	if s < 1048576 {
		return strconv.Itoa(s/1024) + "KB"
	}
	if s < 1073741824 {
		num := fmt.Sprintf("%.2f", float64(s)/1048576)
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
		if num == "1024" {
			return "1GB"
		}
		return num + "MB"
	}
	if s < 1099511627776 {
		num := fmt.Sprintf("%.2f", float64(s)/1073741824)
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
		if num == "1024" {
			return "1TB"
		}
		return num + "GB"
	}
	if s < 1125899906842624 {
		num := fmt.Sprintf("%.3f", float64(s)/1099511627776)
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
		if num == "1024" {
			return "1PB"
		}
		return num + "TB"
	}
	if s < 1152921504606846976 {
		num := fmt.Sprintf("%.3f", float64(s)/1125899906842624)
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
		if num == "1024" {
			return "1EB"
		}
		return num + "PB"
	}
	{
		num := fmt.Sprintf("%.3f", float64(s)/1125899906842624)
		num = strings.TrimRight(num, "0")
		num = strings.TrimRight(num, ".")
		return num + "EB"
	}
}
