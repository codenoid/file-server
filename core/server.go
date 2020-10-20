package core

import (
	"net/http"
	"strings"
)

var conf = config{}

func serveFile(w http.ResponseWriter, r *http.Request) {

	identity := ""

	switch conf.IdentityType {
	case "uri-segment", "urisegment":
		switch conf.IdentitySource {
		case "last":
			arrOfURI := strings.Split(r.URL.Path, "/")
			identity = arrOfURI[len(arrOfURI)-1]
		}
	case "query", "query-param":
		identity = r.URL.Query().Get(conf.IdentitySource)
	}

	// handle error ?
	file := make([]byte, 0)
	for _, source := range conf.Sources {
		if len(file) > 0 {
			break
		}

		switch source.Type {
		case "fs", "filesystem", "localfile", "disk":
			file, _ = ReadFile(source.Source, identity)
		case "url", "web":
			file, _ = ReadURL(source.Source, identity)
		}
	}

	if len(file) == 0 {
		w.WriteHeader(404)
	}
	w.Write(file)
}

// StartServer start file-server server
func StartServer(bind, configFile string) {

	conf = loadConfig(configFile)

	http.HandleFunc("/", serveFile)
	http.ListenAndServe(bind, nil)
}
