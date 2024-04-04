package records

import "wordOfWisdom/internal/business/domains"

func (r *Quote) ToDomain() *domains.QuoteDomain {
	return &domains.QuoteDomain{
		ID:   r.ID,
		Text: r.Text,
	}
}
