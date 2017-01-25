/*
	Pattern Observer:
		O padrão Observer funciona como um canal no YouTube, onde sempre que o canal publica um novo vídeo,
		as pessoas inscritas recebem uma notificação dessa mudança. Assim como na vida real,
		uma pessoa pode se inscrever no canal para ficar por dentro de todas atualizações ou 
		se desinscrever para não receber notificações.
		Seguindo esse exemplo, temos o YouTube sendo o (Observable) e as pessoas sendo o (Observer)
	Exemplos: 
		Tenho um serviço que verifica novos arquivos em uma pasta e notifica o sistema que é preciso
		processar, mover arquivo e notificar o usuário na tela sem a necessidade de uma consulta.
		Todas essas classes iram assina o verificador para receberem atualizações de mudanças.
 */
package main

import(
	"fmt"
)

type Observer interface {
	Update(message string)
}

type Observable interface {
	RegisterObserver(ob Observer)
	RemoveObserver(ob Observer)
	NotifyObservers()
}

type Person struct {
	Name string
}

func (this Person) Update(message string) {
	fmt.Printf("%s recebeu uma nova mensagem do YouTube: %s \n", this.Name, message)
}

type YouTube struct {
	Observers []Observer
}

func (this *YouTube) RegisterObserver(ob Observer) {
	fmt.Printf("YouTube recebeu um inscrito: %s \n", ob)
	this.Observers = append(this.Observers, ob)
}

func (this *YouTube) RemoveObserver(ob Observer) {
	for i, observer := range this.Observers {
		if observer == ob {
			fmt.Printf("YouTube perdeu um inscrito: %s \n", ob)
			this.Observers = append(this.Observers[:i], this.Observers[i+1:]...)
		}
	}
}

func (this YouTube) NotifyObservers() {
	for _, ob := range this.Observers {
		ob.Update("Saiu um novo vídeo hoje")
	}
}

func main() {
	wilson := Person{"Wilson"}
	alex := Person{"Alex"}
	youTube := YouTube{}
	
	youTube.RegisterObserver(wilson)
	youTube.RegisterObserver(alex)
	
	fmt.Println(len(youTube.Observers))
	
	youTube.RemoveObserver(wilson)
	youTube.RemoveObserver(alex)
	
	fmt.Println(len(youTube.Observers))

	youTube.RegisterObserver(wilson)
	youTube.RegisterObserver(alex)
	
	fmt.Println(len(youTube.Observers))

	youTube.NotifyObservers()
}