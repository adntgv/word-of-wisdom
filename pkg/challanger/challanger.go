package challanger

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Challanger interface {
	GenerateChallenge() string
	SolveChallenge(challange, prefix string) string
	IsValidNonce(nonce, challange, prefix string) bool
	GetPrefix() string
}

type defaultChallanger struct {
	difficulty int
}

// GetPrefix implements Challanger.
func (d *defaultChallanger) GetPrefix() string {
	prefix := ""

	for i := 0; i < d.difficulty; i++ {
		prefix += "0"
	}

	return prefix
}

// GenerateChallenge implements Challanger.
func (d *defaultChallanger) GenerateChallenge() string {
	return strconv.Itoa(rand.Intn(100000))
}

// IsValidNonce implements Challanger.
func (d *defaultChallanger) IsValidNonce(nonce string, challange string, prefix string) bool {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%s%s", challange, nonce)))
	return strings.HasPrefix(fmt.Sprintf("%x", hash), prefix)
}

// SolveChallenge implements Challanger.
func (d *defaultChallanger) SolveChallenge(challange string, prefix string) string {
	for i := 0; ; i++ {
		hash := sha256.Sum256([]byte(fmt.Sprintf("%s%d", challange, i)))
		if strings.HasPrefix(fmt.Sprintf("%x", hash), prefix) {
			return fmt.Sprintf("%d", i)
		}
	}
}

func NewChallanger(difficulty int) Challanger {
	return &defaultChallanger{
		difficulty: difficulty,
	}
}
