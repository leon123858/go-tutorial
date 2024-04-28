package key

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateKeys(num int, seed int64) []string {
	generator := rand.New(rand.NewSource(seed))

	keys := make(map[string]bool)
	result := make([]string, 0, len(charset))

	for len(result) < num {
		key := generateKey(generator)
		if _, exists := keys[key]; !exists {
			keys[key] = true
			result = append(result, key)
		}
	}

	return result
}

func generateKey(randGen *rand.Rand) string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[randGen.Intn(len(charset))]
	}
	return string(b)
}
