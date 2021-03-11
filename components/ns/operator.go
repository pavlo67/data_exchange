package ns

import (
	"github.com/pavlo67/common/common/strlib"
)

type Item struct {
	Host     string `json:"host,omitempty"     bson:"host,omitempty"`
	Path     string `json:"path,omitempty"     bson:"path,omitempty"`
	Fragment string `json:"fragment,omitempty" bson:"fragment,omitempty"`
}

func (item *Item) IsValid() bool {
	if item == nil {
		return false
	}
	return strlib.ReSpaces.ReplaceAllString(item.Host, "") != "" &&
		strlib.ReSpaces.ReplaceAllString(item.Path, "") != "" &&
		strlib.ReSpaces.ReplaceAllString(item.Fragment, "") != ""
}

func (item *Item) URN() URN {
	if item == nil {
		return ""
	}

	host := strlib.ReSpaces.ReplaceAllString(item.Host, "")
	path := strlib.ReSpaces.ReplaceAllString(item.Path, "")
	id := strlib.ReSpaces.ReplaceAllString(item.Fragment, "")

	if len(id) > 0 {
		return URN(host + PathDelim + path + IDDelim + id)
	} else if len(path) > 0 {
		return URN(host + PathDelim + path)
	} else if len(host) > 0 {
		return URN(host)
	} else {
		return ""
	}
}

func (item *Item) String() string {
	return string(item.URN())
}

//func FromURLRaw(urlRaw string) Item {
//	urlWithoutProto := reProto.ReplaceAllString(strings.TrimSpace(urlRaw), "")
//	domain := reHostDelim.ReplaceAllString(urlWithoutProto, "")
//
//	// TODO!!! clean more
//
//	return Item{
//		Host: domain,
//		Path:   urlWithoutProto[len(domain):],
//	}
//}
