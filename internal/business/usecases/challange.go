package usecases

import (
	"context"
	"fmt"
	"wordOfWisdom/internal/business/domains"
	"wordOfWisdom/pkg/challanger"
)

type challangeUsecase struct {
	challanger challanger.Challanger
}

// Generate implements domains.ChallangeUsecase.
func (r *challangeUsecase) Generate(ctx context.Context) (*domains.ChallangeDomain, error) {
	challange := r.challanger.GenerateChallenge()

	return &domains.ChallangeDomain{
		Challange: challange,
		Prefix:    r.challanger.GetPrefix(),
	}, nil

}

// Validate implements domains.ChallangeUsecase.
func (r *challangeUsecase) Validate(ctx context.Context, challange *domains.ChallangeDomain) error {
	if solved := r.challanger.IsValidNonce(challange.Nonce, challange.Challange, challange.Prefix); !solved {
		return fmt.Errorf("challange not solved %v", challange)
	}

	return nil
}

func NewChallangeUsecase(challanger challanger.Challanger) domains.ChallangeUsecase {
	return &challangeUsecase{
		challanger: challanger,
	}
}
