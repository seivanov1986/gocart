package cache

import (
	"encoding/json"
)

type BreadcrumbsHtml struct {
	Url  string `db:"url"`
	Name string `db:"name"`
}

type SchemaOrg interface {
	SchemaOrgBreadCrumbs(items []BreadcrumbsHtml)
	SchemaOrgItemList(items []string, name string)
	Get() string
	Clean()
}

type schemaOrg struct {
	schemes []map[string]interface{}
}

func NewSchemaOrg() *schemaOrg {
	return &schemaOrg{
		schemes: []map[string]interface{}{},
	}
}

type item struct {
	ID   string `json:"@id"`
	Name string `json:"name"`
}

type listItem struct {
	Type     string `json:"@type"`
	Position int64  `json:"position"`
	Item     item   `json:"item"`
}

func (s *schemaOrg) Clean() {
	s.schemes = []map[string]interface{}{}
}

func (s *schemaOrg) SchemaOrgBreadCrumbs(items []BreadcrumbsHtml) {
	list := make([]listItem, 0, len(items))
	for k, v := range items {
		list = append(list, listItem{
			Type:     "ListItem",
			Position: int64(k),
			Item: item{
				ID:   v.Url,
				Name: v.Name,
			},
		})
	}

	jsonObject := map[string]interface{}{
		"@context":        "https://schema.org",
		"@type":           "BreadcrumbList",
		"itemListElement": list,
	}

	s.schemes = append(s.schemes, jsonObject)
}

func (s *schemaOrg) SchemaOrgItemList(items []string, name string) {
	jsonObject := map[string]interface{}{
		"@context":        "https://schema.org",
		"@type":           "ItemList",
		"itemListElement": items,
		"itemListOrder":   "https://schema.org/ItemListOrderDescending",
		"name":            name,
	}

	s.schemes = append(s.schemes, jsonObject)
}

func (s *schemaOrg) Get() string {
	bytes, _ := json.Marshal(s.schemes)
	return string(bytes)
}
