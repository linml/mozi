package common

import (
	crand "crypto/rand"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	lowercase    = "abcdefghijklmnopqrstuvwxyz"
	lowercaseLen = len(lowercase)

	uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	uppercaseLen = len(uppercase)

	digits    = "0123456789"
	digitsLen = len(digits)
)

func RandString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func RandDigitString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		idx, _ := crand.Int(crand.Reader, big.NewInt(10))
		b[i] = digits[int(idx.Int64())]
	}
	return string(b)
}

func RandLowerString(n int) string {
	return strings.ToLower(RandString(n))
}

func Str2Time(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	t, err := time.ParseInLocation("20060102150405", s, loc)
	return t, err
}

func Time2Str(t time.Time) string {
	return t.Format("20060102150405")
}

func GetTimeNowString() string {
	return time.Now().Format("20060102150405")
}

func GetDateNowString() string {
	return time.Now().Format("20060102")
}

func Time2Str1(t time.Time, format string) string {
	// 2006-01-02 15:04:05
	return t.Format(format)
}

func InetNtoa(ip int64) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAton(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()

}

func Float64Equal(f1, f2 float64) bool {
	_f1 := decimal.NewFromFloat(f1)
	_f2 := decimal.NewFromFloat(f2)
	return _f1.Equal(_f2)
}

func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func GetInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func GetContent(m map[string]string) string {
	s := ""
	for i, _ := range m {
		s += fmt.Sprintf("%s:%s;", i, m[i])
	}
	return s
}

func GetContentWithMosaics(m map[string]string, keys []string) string {
	s := ""
	for i, _ := range m {
		if CheckInStr(keys, i) {
			s += fmt.Sprintf("%s:%s;", i, "***")
		} else {
			s += fmt.Sprintf("%s:%s;", i, m[i])
		}
	}
	return s
}

func CheckInStr(dis []string, k string) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func CheckInInt(dis []int, k int) bool {
	for i, _ := range dis {
		if dis[i] == k {
			return true
		}
	}
	return false
}

func GetFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func RandNewLowercaseStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(lowercaseLen)
		data[i] = byte(lowercase[idx])
	}

	return string(data)
}

func RandNewUppercaseStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(uppercaseLen)
		data[i] = byte(uppercase[idx])
	}

	return string(data)
}

func RandNewDigitsStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(digitsLen)
		data[i] = byte(digits[idx])
	}

	return string(data)
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func HttpPost(url string, contentType string, params map[string]string) (string, error) {
	var b string
	for k, v := range params {
		b += fmt.Sprintf("&%s=%s", k, v)
	}
	if len(b) > 0 {
		b = b[1:]
	}
	resp, err := http.Post(url, contentType, strings.NewReader(b))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil

}

func Mosaics(s string, start int, end int) string {
	nameRune := []rune(s)
	sLen := len(nameRune)
	if sLen <= 1 {
		return s
	}
	if sLen <= start+end {
		return "***"
	}
	return string(nameRune[0:start]) + "***" + string(nameRune[sLen-end:])
}

func S2IList(l []string) ([]int, error) {
	_l := make([]int, len(l))

	for i, v := range l {
		b, err := strconv.Atoi(v)
		if err != nil {
			return _l, err
		} else {
			_l[i] = b
		}
	}
	return _l, nil
}
