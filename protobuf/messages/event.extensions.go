package messages

import "log"

func (m *Event) GetStringId() string {
	if m != nil {
		oid, er := m.Id.GetObjectID()
		if er != nil {
			log.Printf("error getting objectid %v\n", er)
		}
		return oid.Hex()
	}
	return ""
}
