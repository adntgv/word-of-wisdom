package usecases

import (
	"context"
	"math/rand"
	"strconv"
	"wordOfWisdom/internal/business/domains"
)

type challangeUsecase struct {
}

// Generate implements domains.ChallangeUsecase.
func (r *challangeUsecase) Generate(ctx context.Context) (*domains.ChallangeDomain, error) {
	panic("unimplemented")
}

// Validate implements domains.ChallangeUsecase.
func (r *challangeUsecase) Validate(ctx context.Context, challange *domains.ChallangeDomain) error {
	panic("unimplemented")
}

func NewChallangeUsecase() domains.ChallangeUsecase {
	return &challangeUsecase{}
}

func generateChallenge() string {
	return strconv.Itoa(rand.Intn(100000))
}
func (h *connectionHandler) isValidPoW(challenge, nonce string) bool {
	return powlib.IsValidPoW(challenge, nonce)
}
