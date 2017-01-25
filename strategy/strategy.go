package main

import "fmt"

type Sound interface {
	Execute()
}

type SoundMP3 struct{}

func (this SoundMP3) Execute() {
	fmt.Println("Executando MP3")
}

type SoundWAV struct{}

func (this SoundWAV) Execute() {
	fmt.Println("Executando WAV")
}

type buttonAbstract struct {
	Text  string
	Sound Sound
}

func (this buttonAbstract) GetText() string {
	return this.Text
}

func (this buttonAbstract) SoundClick() {
	this.Sound.Execute()
}

func (this *buttonAbstract) ChangeSoundType(sound Sound) {
	this.Sound = sound
}

type button struct {
	buttonAbstract
	Icon string
}

func (this button) GetIcon() string {
	return this.Icon
}

func NewButton(text string, sound Sound, icon string) button {
	return button{buttonAbstract{text, sound}, icon}
}

type menuItem struct {
	buttonAbstract
	Parent string
}

func (this menuItem) GetParent() string {
	return this.Parent
}

func NewMenuItem(text string, sound Sound, parent string) menuItem {
	return menuItem{buttonAbstract{text, sound}, parent}
}

func main() {
	button := NewButton("Cancelar", &SoundMP3{}, "X")
	fmt.Printf("Botão: %s | Icone: %s \n", button.GetText(), button.GetIcon())
	button.SoundClick()
	button.ChangeSoundType(&SoundWAV{})
	button.SoundClick()

	fmt.Println()

	menuItem := NewMenuItem("Pessoas", &SoundMP3{}, "Cadastro")
	fmt.Printf("Menu Item: %s | Menu: %s \n", menuItem.GetText(), menuItem.GetParent())
	menuItem.SoundClick()
	menuItem.ChangeSoundType(&SoundWAV{})
	menuItem.SoundClick()
}
