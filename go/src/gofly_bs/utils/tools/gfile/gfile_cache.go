package gfile

import (
	"context"
	"time"

	"gofly/utils/tools/command"
	"gofly/utils/tools/gcache"
	"gofly/utils/tools/gcode"
	"gofly/utils/tools/gerror"
	"gofly/utils/tools/gfsnotify"
	"gofly/utils/tools/intlog"
)

const (
	defaultCacheDuration  = "1m"             // defaultCacheExpire is the expire time for file content caching in seconds.
	commandEnvKeyForCache = "gf.gfile.cache" // commandEnvKeyForCache is the configuration key for command argument or environment configuring cache expire duration.
)

var (
	// Default expire time for file content caching.
	cacheDuration = getCacheDuration()

	// internalCache is the memory cache for internal usage.
	internalCache = gcache.New()
)

func getCacheDuration() time.Duration {
	cacheDurationConfigured := command.GetOptWithEnv(commandEnvKeyForCache, defaultCacheDuration)
	d, err := time.ParseDuration(cacheDurationConfigured)
	if err != nil {
		panic(gerror.WrapCodef(
			gcode.CodeInvalidConfiguration,
			err,
			`error parsing string "%s" to time duration`,
			cacheDurationConfigured,
		))
	}
	return d
}

// GetContentsWithCache returns string content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
func GetContentsWithCache(path string, duration ...time.Duration) string {
	return string(GetBytesWithCache(path, duration...))
}

// GetBytesWithCache returns []byte content of given file by `path` from cache.
// If there's no content in the cache, it will read it from disk file specified by `path`.
// The parameter `expire` specifies the caching time for this file content in seconds.
func GetBytesWithCache(path string, duration ...time.Duration) []byte {
	var (
		ctx      = context.Background()
		expire   = cacheDuration
		cacheKey = commandEnvKeyForCache + path
	)

	if len(duration) > 0 {
		expire = duration[0]
	}
	r, _ := internalCache.GetOrSetFuncLock(ctx, cacheKey, func(ctx context.Context) (interface{}, error) {
		b := GetBytes(path)
		if b != nil {
			// Adding this `path` to gfsnotify,
			// it will clear its cache if there's any changes of the file.
			_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
				_, err := internalCache.Remove(ctx, cacheKey)
				if err != nil {
					intlog.Errorf(ctx, `%+v`, err)
				}
				gfsnotify.Exit()
			})
		}
		return b, nil
	}, expire)
	if r != nil {
		return r.Bytes()
	}
	return nil
}
