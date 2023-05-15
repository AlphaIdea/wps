package wps

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
)

// Sign signs the values with the secretKey
// @param values: the values to sign, secretKey: the secret key
// @return the signature
func Sign(values url.Values, secretKey string) string {
	contents := stringToSign(values, secretKey)
	signature := base64.StdEncoding.EncodeToString(GetHmacSHA1(secretKey, string(contents)))
	return signature
}

// stringToSign generates the string to sign
// @param values: the values to sign, secretKey: the secret key
// @return the string to sign
func stringToSign(values url.Values, secretKey string) []byte {
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	buf := &bytes.Buffer{}
	for _, k := range keys {
		fmt.Fprintf(buf, "%s=%s", k, values.Get(k))
	}
	fmt.Fprintf(buf, "_w_secretkey=%s", secretKey)
	return buf.Bytes()
}

// GetHmacSHA1 generates the hmac sha1
// @param key: the key, data: the data
// @return the hmac sha1
func GetHmacSHA1(key, data string) []byte {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return mac.Sum(nil)
}
