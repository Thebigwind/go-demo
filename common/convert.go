package common

import (
	"bytes"
	"encoding/binary"
	encodingJson "encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func IfFloat(v interface{}) (float64, error) {
	switch vv := v.(type) {
	case encodingJson.Number:
		return vv.Float64()
	case float32:
		return float64(vv), nil
	case float64:
		return vv, nil
	case string:
		i, err := strconv.ParseFloat(vv, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case []byte:
		i, err := strconv.ParseFloat(string(vv), 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		return make([]byte, 0)
	}
	return buffer.Bytes()
}

func IfInt(v interface{}) (int, error) {
	switch vv := v.(type) {
	case int:
		return vv, nil
	case int8:
		return int(vv), nil
	case int16:
		return int(vv), nil
	case int32:
		return int(vv), nil
	case int64:
		return int(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 32)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}
func IfInt16(v interface{}) (int16, error) {
	switch vv := v.(type) {
	case int:
		return int16(vv), nil
	case int8:
		return int16(vv), nil
	case int16:
		return int16(vv), nil
	case int32:
		return int16(vv), nil
	case int64:
		return int16(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 16)
		if err != nil {
			return 0, err
		}
		return int16(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 16)
		if err != nil {
			return 0, err
		}
		return int16(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 16)
		if err != nil {
			return 0, err
		}
		return int16(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfInt32(v interface{}) (int32, error) {
	switch vv := v.(type) {
	case int:
		return int32(vv), nil
	case int8:
		return int32(vv), nil
	case int16:
		return int32(vv), nil
	case int32:
		return int32(vv), nil
	case int64:
		return int32(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfInt8(v interface{}) (int8, error) {
	switch vv := v.(type) {
	case int:
		return int8(vv), nil
	case int8:
		return vv, nil
	case int16:
		return int8(vv), nil
	case int32:
		return int8(vv), nil
	case int64:
		return int8(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 8)
		if err != nil {
			return 0, err
		}
		return int8(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 8)
		if err != nil {
			return 0, err
		}
		return int8(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 8)
		if err != nil {
			return 0, err
		}
		return int8(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfInt64(v interface{}) (int64, error) {
	switch vv := v.(type) {
	case int:
		return int64(vv), nil
	case int8:
		return int64(vv), nil
	case int16:
		return int64(vv), nil
	case int32:
		return int64(vv), nil
	case int64:
		return vv, nil
	case float64:
		return int64(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfUint(v interface{}) (uint, error) {
	switch vv := v.(type) {
	case int:
		return uint(vv), nil
	case int8:
		return uint(vv), nil
	case int16:
		return uint(vv), nil
	case int32:
		return uint(vv), nil
	case int64:
		return uint(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}
func IfUint16(v interface{}) (uint16, error) {
	switch vv := v.(type) {
	case int:
		return uint16(vv), nil
	case int8:
		return uint16(vv), nil
	case int16:
		return uint16(vv), nil
	case int32:
		return uint16(vv), nil
	case int64:
		return uint16(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 16)
		if err != nil {
			return 0, err
		}
		return uint16(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 16)
		if err != nil {
			return 0, err
		}
		return uint16(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint16(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfUint32(v interface{}) (uint32, error) {
	switch vv := v.(type) {
	case int:
		return uint32(vv), nil
	case int8:
		return uint32(vv), nil
	case int16:
		return uint32(vv), nil
	case int32:
		return uint32(vv), nil
	case int64:
		return uint32(vv), nil
	case []byte:
		i, err := strconv.ParseUint(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseUint(string(vv), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfUint8(v interface{}) (uint8, error) {
	switch vv := v.(type) {
	case int:
		return uint8(vv), nil
	case int8:
		return uint8(vv), nil
	case int16:
		return uint8(vv), nil
	case int32:
		return uint8(vv), nil
	case int64:
		return uint8(vv), nil
	case []byte:
		i, err := strconv.ParseInt(string(vv), 10, 8)
		if err != nil {
			return 0, err
		}
		return uint8(i), nil
	case encodingJson.Number:
		i, err := strconv.ParseInt(string(vv), 10, 8)
		if err != nil {
			return 0, err
		}
		return uint8(i), nil
	case string:
		i, err := strconv.ParseInt(vv, 10, 8)
		if err != nil {
			return 0, err
		}
		return uint8(i), nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}

func IfUint64(v interface{}) (uint64, error) {
	switch vv := v.(type) {
	case int:
		return uint64(vv), nil
	case int8:
		return uint64(vv), nil
	case int16:
		return uint64(vv), nil
	case int32:
		return uint64(vv), nil
	case int64:
		return uint64(vv), nil
	case []byte:
		i, err := strconv.ParseUint(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case encodingJson.Number:
		i, err := strconv.ParseUint(string(vv), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case string:
		i, err := strconv.ParseUint(vv, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, fmt.Errorf("unsupported type: %v", v)
}
func IfString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}

	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}

func IfBool(v interface{}) bool {
	s, ok := v.(bool)
	if !ok {
		return false
	}

	return s
}

// copy from redigo
type Error string

func (err Error) Error() string { return string(err) }

func Values(reply interface{}, err error) ([]interface{}, error) {
	if err != nil {
		return nil, err
	}

	switch reply := reply.(type) {
	case []interface{}:
		return reply, nil
	case nil:
		return nil, errors.New("nil returned")
	case Error:
		return nil, reply
	}
	return nil, fmt.Errorf("unexpected type for Values, got type %T", reply)
}

func IfMap(result interface{}, err error) (map[string]interface{}, error) {
	values, err := Values(result, err)
	if err != nil {
		return nil, err
	}
	if len(values)%2 != 0 {
		return nil, errors.New("IfMap expects even number of values result")
	}
	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, okKey := values[i].([]byte)
		value, okValue := values[i+1].([]byte)
		if !okKey || !okValue {
			return nil, errors.New("IfMap key not a bulk interface value")
		}
		m[string(key)] = value
	}
	return m, nil
}
