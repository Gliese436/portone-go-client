// Package jsonutil 은 생성된 타입이 사용하는 JSON 보조 함수를 제공합니다.
package jsonutil

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// IsNull 은 data 가 JSON null 인지 확인합니다.
func IsNull(data []byte) bool {
	return bytes.Equal(bytes.TrimSpace(data), []byte("null"))
}

// MarshalWithDiscriminators 는 v 를 JSON 객체로 직렬화한 뒤
// 판별자 필드(key, value 쌍)를 맨 앞에 추가합니다.
//
// v 는 MarshalJSON 재귀를 피하기 위해 메서드가 없는 alias 타입이어야 합니다.
func MarshalWithDiscriminators(v any, keyValues ...string) ([]byte, error) {
	if len(keyValues)%2 != 0 {
		return nil, fmt.Errorf("jsonutil: odd number of discriminator key/values")
	}
	encoded, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	trimmed := bytes.TrimSpace(encoded)
	if len(trimmed) < 2 || trimmed[0] != '{' {
		return nil, fmt.Errorf("jsonutil: expected JSON object, got %s", trimmed)
	}
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i := 0; i < len(keyValues); i += 2 {
		if i > 0 {
			buf.WriteByte(',')
		}
		key, _ := json.Marshal(keyValues[i])
		value, _ := json.Marshal(keyValues[i+1])
		buf.Write(key)
		buf.WriteByte(':')
		buf.Write(value)
	}
	rest := bytes.TrimSpace(trimmed[1:])
	if len(rest) > 0 && rest[0] != '}' {
		buf.WriteByte(',')
	}
	buf.Write(rest)
	return buf.Bytes(), nil
}
