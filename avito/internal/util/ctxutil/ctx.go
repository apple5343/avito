package ctxutil

import "context"

type ContextKey string

var (
	apiKeyCtxKey ContextKey = "api_key"
	userIDCtxKey ContextKey = "user_id"
)

func ContextWithAPIKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, apiKeyCtxKey, key)
}

func APIKeyFromContext(ctx context.Context) string {
	key, ok := ctx.Value(apiKeyCtxKey).(string)
	if !ok {
		return ""
	}
	return key
}

func ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDCtxKey, userID)
}

func UserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(userIDCtxKey).(int)
	if !ok {
		return 0
	}
	return userID
}
