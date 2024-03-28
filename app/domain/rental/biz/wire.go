package biz

import (
	"github.com/blackhorseya/sion/app/domain/rental/repo/asset/irent"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewRentalBiz,
	irent.NewAssetRepo,
)
