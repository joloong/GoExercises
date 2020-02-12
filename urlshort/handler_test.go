package urlshort

import "testing"

func TestParseYaml(t *testing.T) {
	yaml := `
- path: /urlshort
  url: https://github.com/joloong/GoExercises/tree/master/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	pathUrls, err := parseYaml([]byte(yaml))
	if err != nil {
		t.Errorf("parseYaml should not throw an error: %v", err)
	}

	correctPathUrls := []pathUrl{
		pathUrl{"/urlshort", "https://github.com/joloong/GoExercises/tree/master/urlshort"},
		pathUrl{"/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution"},
	}
	for i, p := range pathUrls {
		if p.Path != correctPathUrls[i].Path || p.URL != correctPathUrls[i].URL {
			t.Errorf("parseYaml - got: %v, want: %v.\n", p, correctPathUrls[i])
		}
	}
}

func TestYamlToMap(t *testing.T) {
	pathUrls := []pathUrl{
		pathUrl{"/urlshort", "https://github.com/joloong/GoExercises/tree/master/urlshort"},
		pathUrl{"/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution"},
	}
	pathToUrls := yamlToMap(pathUrls)
	correctPathToUrls := map[string]string{
		"/urlshort":       "https://github.com/joloong/GoExercises/tree/master/urlshort",
		"/urlshort-final": "https://github.com/gophercises/urlshort/tree/solution",
	}
	for k, v := range pathToUrls {
		if v != correctPathToUrls[k] {
			t.Errorf("yamlToMap key: %v - got: %v, want: %v.\n", k, v, correctPathToUrls[k])
		}
	}
}
