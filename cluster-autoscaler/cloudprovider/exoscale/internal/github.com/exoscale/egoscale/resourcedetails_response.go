/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package egoscale

import "fmt"

// Response returns the struct to unmarshal
func (ListResourceDetails) Response() interface{} {
	return new(ListResourceDetailsResponse)
}

// ListRequest returns itself
func (ls *ListResourceDetails) ListRequest() (ListCommand, error) {
	if ls == nil {
		return nil, fmt.Errorf("%T cannot be nil", ls)
	}
	return ls, nil
}

// SetPage sets the current apge
func (ls *ListResourceDetails) SetPage(page int) {
	ls.Page = page
}

// SetPageSize sets the page size
func (ls *ListResourceDetails) SetPageSize(pageSize int) {
	ls.PageSize = pageSize
}

// Each triggers the callback for each, valid answer or any non 404 issue
func (ListResourceDetails) Each(resp interface{}, callback IterateItemFunc) {
	items, ok := resp.(*ListResourceDetailsResponse)
	if !ok {
		callback(nil, fmt.Errorf("wrong type, ListResourceDetailsResponse was expected, got %T", resp))
		return
	}

	for i := range items.ResourceDetail {
		if !callback(&items.ResourceDetail[i], nil) {
			break
		}
	}
}
