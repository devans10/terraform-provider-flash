package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
)

func flattenVolume(in []flasharray.ConnectedVolume) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["vol"] = v.Vol
		m["lun"] = v.Lun

		out[i] = m
	}
	return out
}
