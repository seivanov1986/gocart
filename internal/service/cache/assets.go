package cache

import (
	"bytes"
	"html/template"
	"os"
	"sort"
	"strings"
	"time"
)

type AssetItemDependency struct {
	Path       string
	Type       string
	Preload    bool
	Dependency []string
}

type ResultList struct {
	Path    string
	Time    int64
	Type    string
	Preload bool
}

type AssetOption struct {
	Sort    int64
	Type    string
	Preload bool
}

type assets struct {
	jsList      map[string]AssetOption
	cssList     map[string]AssetOption
	preloadList map[string]AssetOption
}

type Assets interface {
	AddJsList(pathList []AssetItemDependency)
	AddCssList(pathList []AssetItemDependency)
	AddPreloadList(pathList []AssetItemDependency)
	GetJsList() []ResultList
	GetJsTemplate() string
	GetCssList() []ResultList
	GetCssTemplate() string
	GetPreloadList() []ResultList
	GetPreloadTemplate() string
}

func NewAsset() *assets {
	return &assets{
		jsList:      map[string]AssetOption{},
		cssList:     map[string]AssetOption{},
		preloadList: map[string]AssetOption{},
	}
}

func (a assets) AddJsList(pathList []AssetItemDependency) {
	for _, v := range pathList {
		pathIndex, ok := a.jsList[v.Path]
		if !ok {
			pathIndex = AssetOption{Sort: 0}

			for index, p := range a.jsList {
				value := a.jsList[index]
				value.Sort = p.Sort + 1
				a.jsList[index] = value
			}

			a.jsList[v.Path] = pathIndex
		}

		for _, dep := range v.Dependency {
			depIndex, ok := a.jsList[dep]
			if !ok {
				value := a.jsList[dep]
				value.Sort = pathIndex.Sort + 1
				a.jsList[dep] = value
			} else if depIndex.Sort <= pathIndex.Sort {
				value := a.jsList[dep]
				value.Sort = pathIndex.Sort + 1
				a.jsList[dep] = value
			}
		}
	}
}

func (a assets) AddCssList(pathList []AssetItemDependency) {
	for _, v := range pathList {
		pathIndex, ok := a.cssList[v.Path]
		if !ok {
			pathIndex = AssetOption{Sort: 0, Type: v.Type, Preload: v.Preload}

			for index, p := range a.cssList {
				value := a.cssList[index]
				value.Sort = p.Sort + 1
				a.cssList[index] = value
			}

			a.cssList[v.Path] = pathIndex
		}

		for _, dep := range v.Dependency {
			depIndex, ok := a.cssList[dep]
			if !ok {
				value := a.cssList[dep]
				value.Sort = pathIndex.Sort + 1
				a.cssList[dep] = value
			} else if depIndex.Sort <= pathIndex.Sort {
				value := a.cssList[dep]
				value.Sort = pathIndex.Sort + 1
				a.cssList[dep] = value
			}
		}
	}
}

func (a assets) AddPreloadList(pathList []AssetItemDependency) {
	for _, v := range pathList {
		pathIndex, ok := a.preloadList[v.Path]
		if !ok {
			pathIndex = AssetOption{Sort: 0, Type: v.Type, Preload: v.Preload}

			for index, p := range a.preloadList {
				value := a.preloadList[index]
				value.Sort = p.Sort + 1
				a.preloadList[index] = value
			}

			a.preloadList[v.Path] = pathIndex
		}

		for _, dep := range v.Dependency {
			depIndex, ok := a.preloadList[dep]
			if !ok {
				value := a.preloadList[dep]
				value.Sort = pathIndex.Sort + 1
				a.preloadList[dep] = value
			} else if depIndex.Sort <= pathIndex.Sort {
				value := a.preloadList[dep]
				value.Sort = pathIndex.Sort + 1
				a.preloadList[dep] = value
			}
		}
	}
}

func (a assets) GetJsList() []ResultList {
	serviceBasePath := os.Getenv("SERVICE_BASE_PATH")

	type key_value struct {
		Key   string
		Value AssetOption
	}

	var sorted_struct []key_value

	for key, value := range a.jsList {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value.Sort > sorted_struct[j].Value.Sort
	})

	result := make([]ResultList, 0, len(a.jsList))

	for _, v := range sorted_struct {
		modifiedtime := time.Now()

		if strings.HasPrefix(v.Key, "/static") {
			file, err := os.Stat(serviceBasePath + "/service/schemes/public" + v.Key)
			if err != nil {
				continue
			}
			modifiedtime = file.ModTime()
			v.Key = "/static" + v.Key
		}

		result = append(result, ResultList{
			Path:    v.Key,
			Time:    modifiedtime.Unix(),
			Type:    v.Value.Type,
			Preload: v.Value.Preload,
		})
	}

	return result
}

func (a assets) GetCssList() []ResultList {
	serviceBasePath := os.Getenv("SERVICE_BASE_PATH")

	type key_value struct {
		Key   string
		Value AssetOption
	}

	var sorted_struct []key_value

	for key, value := range a.cssList {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value.Sort > sorted_struct[j].Value.Sort
	})

	result := make([]ResultList, 0, len(a.cssList))

	for _, v := range sorted_struct {
		modifiedtime := time.Now()

		if strings.HasPrefix(v.Key, "/static") {
			file, err := os.Stat(serviceBasePath + "/service/schemes/public" + v.Key)
			if err != nil {
				continue
			}
			modifiedtime = file.ModTime()

			v.Key = "/static" + v.Key
		}

		result = append(result, ResultList{
			Path:    v.Key,
			Time:    modifiedtime.Unix(),
			Type:    v.Value.Type,
			Preload: v.Value.Preload,
		})
	}

	return result
}

func (a assets) GetPreloadList() []ResultList {
	serviceBasePath := os.Getenv("SERVICE_BASE_PATH")

	type key_value struct {
		Key   string
		Value AssetOption
	}

	var sorted_struct []key_value

	for key, value := range a.preloadList {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value.Sort > sorted_struct[j].Value.Sort
	})

	result := make([]ResultList, 0, len(a.preloadList))

	for _, v := range sorted_struct {
		modifiedtime := time.Now()

		if strings.HasPrefix(v.Key, "/static") {
			file, err := os.Stat(serviceBasePath + "/service/schemes/public" + v.Key)
			if err != nil {
				continue
			}
			modifiedtime = file.ModTime()

			v.Key = "/static" + v.Key
		}

		result = append(result, ResultList{
			Path:    v.Key,
			Time:    modifiedtime.Unix(),
			Type:    v.Value.Type,
			Preload: v.Value.Preload,
		})
	}

	return result
}

func (a assets) GetCssTemplate() string {
	t, err := template.New("scripts").Parse(`
		{{range $val := .scripts}}<link rel="stylesheet" href="{{$val.Path}}?{{$val.Time}}">
		{{end}}
	`)
	if err != nil {
		return ""
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "scripts", map[string]interface{}{
		"scripts": a.GetCssList(),
	})
	if err != nil {
		return ""
	}

	return buf.String()
}

func (a assets) GetJsTemplate() string {
	t, err := template.New("scripts").Parse(`
		{{range $val := .scripts}}<script src="{{$val.Path}}?{{$val.Time}}"></script>
		{{end}}
	`)
	if err != nil {
		return ""
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "scripts", map[string]interface{}{
		"scripts": a.GetJsList(),
	})
	if err != nil {
		return ""
	}

	return buf.String()
}

func (a assets) GetPreloadTemplate() string {
	t, err := template.New("scripts").Parse(`
		{{range $val := .scripts}}<link rel="preload" href="{{$val.Path}}?{{$val.Time}}" as="{{$val.Type}}">
		{{end}}
	`)
	if err != nil {
		return ""
	}

	buf := &bytes.Buffer{}
	err = t.ExecuteTemplate(buf, "scripts", map[string]interface{}{
		"scripts": a.GetPreloadList(),
	})
	if err != nil {
		return ""
	}

	return buf.String()
}
