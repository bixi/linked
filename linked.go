/*
Linked has basic interfaces for interacting with linked lists.
*/
package linked

type Lister interface {
	Append(Noder, error)
	Push(Noder, error)
	Pop() Noder
	Remove(Noder) error
	Head() Noder
	Tail() Noder
	Length() uint
}

type Noder interface {
	Data() (interface{}, error)
	Next() (Noder, bool, error)
	Prev() (Noder, bool, error)

	SetData(interface{}) error
	SetNext(Noder) error
	SetPrev(Noder) error
}
