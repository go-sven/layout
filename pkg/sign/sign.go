package sign

import (
	"crypto/sha256"
	"github.com/go-sven/layout/pkg/hash"
	"sort"
	"strings"
)

func Generate(params map[string]string, secretKey string) string {
	//参数按照key排序并拼接
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//打开注释为降序排列
	//sort.Strings(keys)
	//sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	var paramStr string
	for _, key := range keys {
		paramStr += key + "=" + params[key] + "&"
	}
	paramStr = strings.TrimSuffix(paramStr, "&")
	//hmac-sha256 加密
	return hash.HMACHash([]byte(paramStr), []byte(secretKey), sha256.New)
}
