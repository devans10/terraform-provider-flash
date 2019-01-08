package purestorage

import (
	"github.com/devans10/go-purestorage/flasharray"
)

func flattenHgroupVolume(in []flasharray.HostgroupConnection) []map[string]interface{} {
	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["vol"] = v.Vol
		m["lun"] = v.Lun

		out[i] = m
	}
	return out
}
