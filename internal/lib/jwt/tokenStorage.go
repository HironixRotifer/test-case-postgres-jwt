package jwt

import (
	"fmt"
	"sync"
)

// Лучше использовать Redis
// Для карты придётся реализовать функционал пересоздания и копировании данных
var (
	mapTokens = map[string]*ts{}
	mu        sync.Mutex
)

type ts struct {
	accessToken  string
	refreshToken string
	Valid        bool
}

// WriteInMap записывает в mapTokens по ключу key информацию о токене ts (token struct)
func WriteInMap(key string, ts ts) error {
	if key == "" {
		return fmt.Errorf("zero value")
	}

	mu.Lock()
	_, ok := mapTokens[key]
	mu.Unlock()
	if ok {
		return fmt.Errorf("token already exist")
	}

	mu.Lock()
	mapTokens[key] = &ts
	mu.Unlock()

	return nil
}

// ReadFromMap читает значение ts из mapTokens по ключу key
func ReadFromMap(key string) (*ts, error) {
	mu.Lock()
	token, ok := mapTokens[key]
	mu.Unlock()

	if !ok {
		fmt.Println(token, ok)
		return nil, fmt.Errorf("token not exist")
	}

	return token, nil
}
