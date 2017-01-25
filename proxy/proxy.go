/*
	Pattern Proxy:
		Consiste em você pegar uma classe concreta que já existe,
		criar uma nova classe que irá encapsular a primeira, dando novos recursos a ela.
		Isso fara com que a primeira tenha mudanças minimas.
	Exemplos: 
		Cache de dados, validar permissão de acesso usuário, rotinas de logs, etc ...
 */
package main

import(
	"fmt"
)

type Image interface {
	Display()
}

type ImageReal struct {
	Filename string
}

func NewImageReal(filename string) *ImageReal {
	image := ImageReal{filename}
	image.LoadFromDisk()
	return &image
}

func (this ImageReal) Display() {
	fmt.Println("Displaying: ", this.Filename)
}

func (this ImageReal) LoadFromDisk() {
	fmt.Println("Loading: ", this.Filename)
}

type ImageCache struct {
	ImageReal *ImageReal
	Filename string
}

func NewImageCache(filename string) ImageCache {
	return ImageCache{nil, filename}
}

func (this *ImageCache) Display() {
	if this.ImageReal == nil {
		this.ImageReal = NewImageReal(this.Filename)
	} else {
		fmt.Println("Using cache")
	}

	this.ImageReal.Display()
}

func main() {
	image := NewImageCache("teste.jpg")
	image.Display()

	fmt.Println()

	image.Display()
}