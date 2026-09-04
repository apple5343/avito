package ctxutil

import "context"

var (
	ApiKeyCtxKey = "api_key"
)

func ContextWithApiKey(ctx context.Context, key string) context.Context {
	return context.WithValue(ctx, ApiKeyCtxKey, key)
}

func ApiKeyFromContext(ctx context.Context) string {
	key, ok := ctx.Value(ApiKeyCtxKey).(string)
	if !ok {
		return ""
	}
	return key
}
