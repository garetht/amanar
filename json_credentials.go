// To parse and unparse this JSON data, add this code to your project and do:
//
//    jSONCredentials, err := UnmarshalJSONCredentials(bytes)
//    bytes, err = jSONCredentials.Marshal()

package amanar

import "encoding/json"

type JSONCredentials []JSONCredential

func UnmarshalJSONCredentials(data []byte) (JSONCredentials, error) {
	var r JSONCredentials
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JSONCredentials) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type JSONCredential struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
