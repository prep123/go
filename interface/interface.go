package main

type sing interface {
	sing()
}

func letsSing(p sing) {
	p.sing()
}

