/*
Node contains node classes for linked lists.
*/
package node

import "github.com/KTAtkinson/linked"

/*
Dnode is a doubly linked node.

A doubly linked node keeps track of both the node in front of it and the node behind
it. This is most useful for doubly-linked lists.
*/
type Dnode struct {
	data interface{}
	next linked.Noder
	prev linked.Noder
}

/*
Value returns the struct representing the value of the node.
*/
func (d *Dnode) Data() (i interface{}, err error) {
	return i, err
}

/*
SetValue sets the value of the node to the given struct.
*/
func (d *Dnode) SetData(i interface{}) (err error) {
	return err
}

/*
Next returns to next node in the list, and a boolean indicating if this is the
last thing in the list.
*/
func (d *Dnode) Next() (n *Dnode, is_last bool, err error) {
	return n, has_next, err
}

/*
SetNext sets the next node in the sequence to the given Dnode.
*/
func (d Dnode) SetNext(n *Dnode) (err error) {
	return err
}

/*
Prev returns the previous node and a boolean indicating if this is the first node.
*/
func (d *Dnode) Prev() (d *Dnode, is_first bool, err error) {
	return n, is_first, err
}

/*
SetPrev sets the previous element to the given node interface.
*/
func (d *Dnode) SetPrev(n *Dnode) (err error) {
	return err
}
