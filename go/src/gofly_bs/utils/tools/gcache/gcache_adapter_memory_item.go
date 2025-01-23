package gcache

import "gofly/utils/tools/gtime"

// IsExpired checks whether `item` is expired.
func (item *adapterMemoryItem) IsExpired() bool {
	return item.e < gtime.TimestampMilli()
}
