package generator

import "time"

// GenIntKeyUUID генерирует уникальный числовой id
func GenIntKeyUUID() int {
	now := time.Now()
	uuid := now.Unix()*1e3 + int64(now.Nanosecond())/1e6
	return int(uuid)
}
