package main

import (
	"fmt"
)

type l interface {
	p() string
	j() int
}

type w interface {
	p() string
	j() int
}

type h interface {
	w() w
}

func aa(h h) {
	fmt.Println(h.w().p())
}

type ll struct {
}

func (l *ll) p() string {
	return ""
}

func (l *ll) j() int {
	return 0
}

type hh struct {
	l *ll
}

func (h *hh) w() l {
	return h.l
}

func init() {
	kk := hh{}
	aa(kk)

}
