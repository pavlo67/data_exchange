package ns

import (
	"regexp"
	"strings"

	"github.com/pavlo67/common/common"
)

// TODO!!! format URN according to RFC

type URN common.IDStr

// DEPRECATED
// type URN = URN

//type URN common.IDStr

// https://www.ietf.org/rfc/rfc2141.txt
//
// <URN>         ::= 1*<URN chars>
// <URN chars>   ::= <trans> | "%" <hex> <hex>
// <trans>       ::= <upper> | <lower> | <number> | <other> | <reserved>
// <hex>         ::= <number> | "A" | "B" | "C" | "D" | "E" | "F" |
// "a" | "b" | "c" | "d" | "e" | "f"
// <other>       ::= "(" | ")" | "+" | "," | "-" | "." |
// ":" | "=" | "@" | ";" | "$" |
// "_" | "!" | "*" | "'"
//
// Depending on the rules governing a namespace, valid identifiers in a namespace might contain characters that are not members
// of the URN character set above (<URN chars>).  Such strings MUST be translated into canonical URN format before using them
// as protocol elements or otherwise passing them on to other applications. Translation is done by encoding each character
// outside the URN character set as a sequence of one to six octets using UTF-8 encoding [5], and the encoding of each of those
//  octets as "%" followed by two characters from the <hex> character set above. The two characters give the hexadecimal
// representation of that octet.
//
// 2.3 Reserved characters
//
// The remaining character set left to be discussed above is the reserved character set, which contains various characters
// reserved from normal use.  The reserved character set follows, with a discussion on the specifics of why each character
// is reserved.
//
// The reserved character set is:
//
// <reserved>    ::= '%" | "/" | "?" | "#"

// -----------------------------------------------------------------------------------------------

const PathDelim = `/`
const IDDelim = `#`

var reProto = regexp.MustCompile(`^https?://`)
var reHostDelim = regexp.MustCompile(PathDelim + `.*`)
var rePathDelimFirst = regexp.MustCompile(`^(` + PathDelim + `)+`)
var rePathDelim = regexp.MustCompile(IDDelim + `.*`)
var reIDDelimFirst = regexp.MustCompile(`^(` + IDDelim + `)+`)

func (urn URN) Item() *Item {
	idStr := strings.TrimSpace(string(urn))
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
