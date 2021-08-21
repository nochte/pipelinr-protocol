package messages

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/wolfeidau/unflatten"
)

func decodeDecoration(value string) interface{} {
	var msi interface{}
	er := json.Unmarshal([]byte(value), &msi)
	if er == nil {
		return msi
	}
	return value
}

func (m *MessageEnvelop) ApplyDecorations() {
	var payloadjs map[string]interface{}
	er := json.Unmarshal([]byte(m.GetPayload()), &payloadjs)
	if er != nil {
		log.Printf("error unmarshaling original payload, using empty (%v)\n", er.Error())
	}
	if len(m.GetDecorations()) != 0 {
		for _, dec := range m.GetDecorations() {
			payloadjs[dec.Key] = decodeDecoration(dec.Value)
		}
	}
	tree := unflatten.Unflatten(payloadjs, func(k string) []string { return strings.Split(k, ".") })
	payload, er := json.Marshal(tree)
	if er != nil {
		log.Printf("error marshaling decoration tree %v\n", er.Error())
	} else {
		m.DecoratedPayload = string(payload)
	}
}
