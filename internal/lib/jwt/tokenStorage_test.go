//go:build !race

package jwt

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ErrorZeroValue    = errors.New("zero value")
	ErrorAlreadyExist = fmt.Errorf("token already exist")
	ErrorNotExist     = fmt.Errorf("token not exist")
	ErrorNil          error
)

// TestWriteInMap_DataRace проверяет работу одновременной записи 100 горутин в map
func TestWriteInMap_DataRace(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value ts
	}{
		{
			name: "write",
			key:  "key1",
			value: ts{
				accessToken:  "",
				refreshToken: "token1",
				Valid:        true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var countGorutines = 100
			var wg sync.WaitGroup
			wg.Add(countGorutines)

			for i := 0; i < countGorutines; i++ {
				go func() {
					defer wg.Done()
					s := strconv.Itoa(i)

					err := WriteInMap(tt.key+s, tt.value)
					assert.NoError(t, err)
				}()
			}

			wg.Wait()

			t.Log(len(mapTokens))
		})
	}
}

// TestWriteInMap проверяет получение исключений
func TestWriteInMap_GetErrors(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		value     ts
		errString string
	}{
		{
			name: "zero value",
			key:  "",
			value: ts{
				accessToken:  "accessToken",
				refreshToken: "refreshToken",
				Valid:        true,
			},
			errString: "zero value",
		},
		{
			name: "token already exist",
			key:  "accessToken",
			value: ts{
				accessToken:  "accessToken",
				refreshToken: "refreshToken",
				Valid:        true,
			},
			errString: "token already exist",
		},
	}

	WriteInMap("accessToken", ts{
		accessToken:  "accessToken",
		refreshToken: "refreshToken",
		Valid:        true,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteInMap(tt.key, tt.value)
			assert.EqualError(t, err, tt.errString)
		})
	}
}

// TestReadFromMap_ReadToken проверка чтения из map несуществующего токена
func TestReadFromMap_TokenNotExist(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		errString string
	}{
		{
			name:      "token not exist",
			key:       "key",
			errString: "token not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadFromMap(tt.key)
			assert.EqualError(t, err, tt.errString)
		})
	}
}

// TestReadFromMap_ReadToken проверка чтения из map
func TestReadFromMap_ReadToken(t *testing.T) {
	tests := []struct {
		name           string
		key            string
		expectedResult ts
	}{
		{
			name: "token not exist",
			key:  "key",
			expectedResult: ts{
				accessToken:  "accessToken",
				refreshToken: "refreshToken",
				Valid:        true,
			},
		},
	}

	WriteInMap("key", ts{
		accessToken:  "accessToken",
		refreshToken: "refreshToken",
		Valid:        true,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, err := ReadFromMap(tt.key)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedResult.accessToken, ts.accessToken)
			assert.Equal(t, tt.expectedResult.refreshToken, ts.refreshToken)
			assert.Equal(t, tt.expectedResult.Valid, ts.Valid)
		})
	}
}
