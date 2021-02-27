package ns

import (
	"regexp"
	"strings"

	"github.com/pavlo67/common/common"
)

// TODO??? use URI / Item / URN according to RFC

type ID common.IDStr

// -----------------------------------------------------------------------------------------------

const PathDelim = `/`
const IDDelim = `#`

var reProto = regexp.MustCompile(`^https?://`)
var reHostDelim = regexp.MustCompile(PathDelim + `.*`)
var rePathDelimFirst = regexp.MustCompile(`^(` + PathDelim + `)+`)
var rePathDelim = regexp.MustCompile(IDDelim + `.*`)
var reIDDelimFirst = regexp.MustCompile(`^(` + IDDelim + `)+`)

func (id ID) Item() *Item {
	idStr := strings.TrimSpace(string(id))
	if len(idStr) < 1 {
		return nil
	}

	host := reHostDelim.ReplaceAllString(idStr, "")
	rest := rePathDelimFirst.ReplaceAllString(strings.TrimSpace(idStr[len(host):]), "")

	path := rePathDelim.ReplaceAllString(rest, "")
	fragment := reIDDelimFirst.ReplaceAllString(strings.TrimSpace(rest[len(path):]), "")

	return &Item{
		Host:     host,
		Path:     path,
		Fragment: fragment,
	}
}

//func (id ID) Short(domain string) ID {
//	if len(id) > len(domain) && string(id[:len(domain)]) == domain && id[len(domain):len(domain)+1] == PathDelim {
//		return id[len(domain):]
//	}
//	return id
//}
//
//func (id ID) Full(domain string) ID {
//	if len(id) > 0 && id[:1] == PathDelim {
//		return ID(domain + string(id))
//	}
//	return id
//}
//
//func IsEqual(identity *Item, is ID, domain string) bool {
//	return identity != nil && is == identity.ID()
//}
