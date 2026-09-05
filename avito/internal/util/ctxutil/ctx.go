package ctxutil

import "context"

type ContextKey string

var (
	apiKeyCtxKey ContextKey = "api_key"
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
