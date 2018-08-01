package main

type singer interface {
	sing()
}

type dancer interface {
	dance()
}

type singAndDance interface {
	singer
	dancer
}