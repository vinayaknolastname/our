package types

import (
	"context"

	admin "github.com/vinayakbot/our/services/common/admin"
)

type AdminService interface {
	CreatePG(context.Context, admin.CreatePGRequest) error
}
