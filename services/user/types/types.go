package types

import (
	"context"

	admin "github.com/vinayaknolastname/our/services/common/admin"
)

type AdminService interface {
	CreatePG(context.Context, admin.CreatePGRequest) error
}
