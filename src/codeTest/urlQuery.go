package main

import (
	"strings"
	"fmt"
)

type value map[string][]string

func test(m value, query string)(err error)  {
	for query != "" {
		key := query
		if i := strings.IndexAny(key, "&;"); i >= 0 {
			fmt.Println(i)
			key, query = key[:i], key[i+1:]
			fmt.Println(key, query)
		} else {
			query = ""
		}
		if key == "" {
			continue
		}
		value := ""
		if i := strings.Index(key, "="); i >= 0 {
			key, value = key[:i], key[i+1:]
		}

		m[key] = append(m[key], value)
	}
	return err
}

func main() {
	url := "ck?info=CmUKKDdic3JkZmd5aWZvYjJfNjAzMDYzNDMyNDU1NTA0NV8xMjQzNjQ4NTQSDWNvbS5zaG91LmRlbmcYtNoGILWVGSjI_gIw8BhAgAVIwAdQAlihCGCAzp7_B2iAzp7_B3CAzp7_BxIDZ2R0GNaYv9cFIkEIibcDENUUGNUUKAAwATjVFEkBAAAAAAAUQFgAYQAAAAAAAPA_aQAAAAAAAAAAecvMzMzMDFlAgQEAAAAAAFB_QCpOCAMSIGM3ZGZmODkzZDU0YWM3NDQ1NDFiYmMwNzU1Y2I4ZjQ1GAgiCk9TX0FORFJPSUQqAzYuMDIFNTA0MDQ6CENBTS1BTDAwQAFQAVgCMMcaOhBiakJfNjBiNWVkMmVlZTc3QABQAGAAaQAAAAAAAPA_cLueLIoBATCiAQc3MDJiZWRmwgEAyAEC0gERCAEVZmbGPx0AAAAAJQAAgD_aAQDiAQoxMTU2NjEwOTAw6AEA-gEAlQIAAIA_&isPb=1"
	v := value{}
	test(v,url)
	fmt.Println(v)
}
