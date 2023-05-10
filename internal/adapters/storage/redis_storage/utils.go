package redis_storage

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
)

const existFlag = "exist"

func compositeKey(keyPrefix string, keyParts ...string) string { // TODO
	builder := strings.Builder{}
	builder.WriteString(keyPrefix)
	for _, part := range keyParts {
		builder.WriteString(":" + part)
	}

	return builder.String()
}

func getKeysByPattern(client *redis.Client, pattern string) ([]string, error) {
	keys := make([]string, 0)
	iter := client.Scan(context.Background(), 0, pattern, 0).Iterator()
	for iter.Next(context.Background()) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}

func delKeysByID(client *redis.Client, keyPrefix string, idArray []int) error {
	for _, id := range idArray {
		idStr := strconv.FormatInt(int64(id), 10)
		key := compositeKey(keyPrefix, idStr)
		err := client.Del(context.Background(), key).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func makeSet[S ~[]E, E comparable](slice S) map[E]bool {
	res := make(map[E]bool, len(slice))

	for _, elem := range slice {
		res[elem] = true
	}

	return res
}

func makeSlice[M map[K]V, K comparable, V any](set M) []K {
	res := make([]K, 0, len(set))

	for elem := range set {
		res = append(res, elem)
	}

	return res
}

func removeDuplicates[S ~[]E, E comparable](slice S) S {
	return makeSlice(makeSet(slice))
}

func containElem[S ~[]E, E comparable](elem E, slice S) bool {
	for _, e := range slice {
		if e == elem {
			return true
		}
	}

	return false
}

func containSlice[S ~[]E, E comparable](subSlice, slice S) bool {
	set := makeSet[S](slice)
	for _, elem := range subSlice {
		contain := set[elem]
		if !contain {
			return false
		}
	}

	return true
}
