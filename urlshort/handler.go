package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if we can match a path, redirect to it
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		// else
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the yaml
	pathUrls, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}

	// 2. Convert YAML array into map
	pathToUrls := yamlToMap(pathUrls)

	return MapHandler(pathToUrls, fallback), nil
}

func parseYaml(yml []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func yamlToMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for i := range pathUrls {
		pathToUrls[pathUrls[i].Path] = pathUrls[i].URL
	}
	return pathToUrls
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// 1. Parse the JSON
	pathUrls, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}

	pathToUrls := jsonToMap(pathUrls)
	return MapHandler(pathToUrls, fallback), nil
}

func parseJSON(jsn []byte) ([]jsonPathUrl, error) {
	var jsonPathUrls []jsonPathUrl
	err := json.Unmarshal(jsn, &jsonPathUrls)
	if err != nil {
		return nil, err
	}
	return jsonPathUrls, nil
}

func jsonToMap(pathUrls []jsonPathUrl) map[string]string {
	pathToUrls := make(map[string]string)
	for i := range pathUrls {
		pathToUrls[pathUrls[i].Path] = pathUrls[i].URL
	}
	return pathToUrls
}

type jsonPathUrl struct {
	Path string `json:"path,omitempty"`
	URL  string `json:"url,omitempty"`
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
