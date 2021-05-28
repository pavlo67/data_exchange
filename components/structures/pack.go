package structures

import "fmt"

//const InPack selectors.Key = "in_pack"

type Data interface {
	Value() interface{}
	IsEqualTo(interface{}) bool
	Stat() *ItemsStat
}

type Pack interface {
	SetDescription(ItemDescription) error
	Description() *ItemDescription
	Data() Data
}

type ItemsStat struct {
	Total    int
	NonEmpty int
	Errored  int
}

func (itemsStat *ItemsStat) String() string {
	if itemsStat == nil {
		return "nil"
	}
	return fmt.Sprintf("Total:%4d, NonEmpty:%4d, Errored:%4d", itemsStat.Total, itemsStat.NonEmpty, itemsStat.Errored)

	//bytes, _ := json.Marshal(itemsStat)
	//return string(bytes)
}
