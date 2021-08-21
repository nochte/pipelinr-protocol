package accounts

import "log"

func (m *User) GetStringId() string {
	if m != nil {
		oid, er := m.Id.GetObjectID()
		if er != nil {
			log.Printf("error getting objectid %v\n", er)
		}
		return oid.Hex()
	}
	return ""
}

func (m *APIKey) GetStringId() string {
	if m != nil {
		oid, er := m.Id.GetObjectID()
		if er != nil {
			log.Printf("error getting objectid %v\n", er)
		}
		return oid.Hex()
	}
	return ""
}
