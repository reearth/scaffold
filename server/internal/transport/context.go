package transport

import (
	"context"

	"github.com/reearth/reearthx/appx"
)

type AuthInfo = appx.AuthInfo

func GetAuthInfo(ctx context.Context) *AuthInfo {
	return appx.GetAuthInfoFromContext(ctx)
}
