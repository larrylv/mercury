package mercury

import (
	"encoding/json"
)

func marshalJSON(v interface{}) []byte {
	b, _ := json.Marshal(v)
	return b
}
