package generate_id

import (
	"fmt"
	"github.com/google/uuid"
)

func NewID(domain string) string {
	code := uuid.New()
	return fmt.Sprintf("%s-%s", domain, code.String())
}
