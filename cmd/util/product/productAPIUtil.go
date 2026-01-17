package productUtil

import (
	"mael/cmd/struct/product"
	"time"
)

func IsResponseObjectCacheValid(object *productStruct.ProductResponseObjectCache) bool {
	if object.ResponseObject == nil || object.CreateTime.IsZero() {
		return false
	}

	expiryTime := object.CreateTime.Add(time.Hour * 24)

	if expiryTime.Before(time.Now()) {
		return false
	}

	return true
}

func CacheResponseObject(object *productStruct.ProductResponseObject, objectCache *productStruct.ProductResponseObjectCache) {
	objectCache.ResponseObject = object
	objectCache.CreateTime = time.Now()
}
