package auth

import (
	"context"
	"server/internal/domain/user/model"
)

type Key string

const ContextKeyUser Key = "user"
const ContextKeyAccessToken Key = "accessToken"

func GetCtxUser(ctx context.Context) *model.User {
	user := ctx.Value(ContextKeyUser).(*model.User)
	return user
}

func SetCtxUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, ContextKeyUser, user)
}

func GetCtxAccessToken(ctx context.Context) *UserToken {
	token := ctx.Value(ContextKeyAccessToken).(*UserToken)
	return token
}

func SetCtxAccessToken(ctx context.Context, token *UserToken) context.Context {
	return context.WithValue(ctx, ContextKeyAccessToken, token)
}
