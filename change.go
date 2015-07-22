package gompatible

import (
	"golang.org/x/tools/go/types"
)

type ChangeKind int

const (
	ChangeUnchanged ChangeKind = iota
	ChangeAdded
	ChangeRemoved
	ChangeCompatible
	ChangeBreaking
)

func (ck ChangeKind) String() string {
	switch ck {
	case ChangeUnchanged:
		return "ChangeUnchanged"
	case ChangeAdded:
		return "ChangeAdded"
	case ChangeRemoved:
		return "ChangeRemoved"
	case ChangeCompatible:
		return "ChangeCompatible"
	case ChangeBreaking:
		return "ChangeBreaking"
	}

	return ""
}

type Change interface {
	TypesObject() types.Object
	ShowBefore() string
	ShowAfter() string
	Kind() ChangeKind
}

func ShowChange(c Change) string {
	switch c.Kind() {
	case ChangeAdded:
		return "+ " + c.ShowAfter()
	case ChangeRemoved:
		return "- " + c.ShowBefore()
	case ChangeUnchanged:
		return "= " + c.ShowBefore()
	case ChangeCompatible:
		return "* " + c.ShowBefore() + " -> " + c.ShowAfter()
	case ChangeBreaking:
		fallthrough
	default:
		return "! " + c.ShowBefore() + " -> " + c.ShowAfter()
	}
}
