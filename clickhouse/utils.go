package clickhouse

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func diffArrays[T any, V comparable](a []T, b []T, hash func(T) V) ([]T, []T) {
	aSet := map[V]bool{}
	bSet := map[V]bool{}

	add := []T{}
	remove := []T{}

	for _, item := range a {
		aSet[hash(item)] = true
	}

	for _, item := range b {
		bSet[hash(item)] = true
	}

	for _, item := range a {
		_, ok := bSet[hash(item)]
		if !ok {
			remove = append(remove, item)
		}
	}

	for _, item := range b {
		_, ok := aSet[hash(item)]
		if !ok {
			add = append(add, item)
		}
	}

	return add, remove
}

func resetValue(d *schema.ResourceData, key string) {
	oldValue, _ := d.GetChange(key)
	d.Set(key, oldValue)
}
