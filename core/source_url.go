package core

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// ReadURL source are url, identity more likely filename
func ReadURL(source, identity string) (val []byte, err error) {

	u, err := url.Parse(source)
	if err != nil {
		return val, err
	}

	u.Path = path.Join(u.Path, identity)
	resp, err := http.Get(u.String())
	if err != nil {
		return val, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		val, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return val, err
		}
	}

	return val, nil
}
