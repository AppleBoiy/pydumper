package util

import (
	"fmt"
	"strings"
)

func SqlGenGroup(grp_n int, mem_n int, start_n int) string {
	grp_list := []string{}
	start := start_n
	for i := 0; i < grp_n; i++ {
		mem_list := []string{}
		for j := 0; j < mem_n; j++ {
			mem_list = append(mem_list, fmt.Sprintf("$%d", start))
			start++
		}
		grp_list = append(grp_list, "("+strings.Join(mem_list, ",")+")")
	}
	return strings.Join(grp_list, ",")
}
