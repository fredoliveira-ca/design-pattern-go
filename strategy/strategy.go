package main

import "fmt"

type SoundInterface interface {
	Execute()
}

type SoundMP3 struct{}

func (this *SoundMP3) Execute() {
	fmt.Println("Executando MP3")
}

type SoundWAV struct{}

func (this *SoundWAV) Execute() {
	fmt.Println("Executando WAV")
}

type buttonAbstract struct {
	Text  string
	Sound SoundInterface
}

func (this *buttonAbstract) GetText() string {
	return this.Text
}

func (this *buttonAbstract) SoundClick() {
	this.Sound.Execute()
}

func (this *buttonAbstract) ChangeSoundType(sound SoundInterface) {
	this.Sound = sound
}

type buttonCommon struct {
	buttonAbstract
}

func NewButtonCommun(text string, sound SoundInterface) *buttonCommon {
	return &buttonCommon{buttonAbstract{text, sound}}
}

func main() {
	buttonMP3 := NewButtonCommun("* Botão Comum MP3", &SoundMP3{})
	fmt.Println(buttonMP3.GetText())
	buttonMP3.SoundClick()

	buttonWAV := NewButtonCommun("* Botão Comum WAV", &SoundWAV{})
	fmt.Println(buttonWAV.GetText())
	buttonWAV.SoundClick()

	buttonChanged := NewButtonCommun("* Botão Misto", &SoundMP3{})
	fmt.Println(buttonChanged.GetText())
	buttonChanged.SoundClick()
	buttonChanged.ChangeSoundType(&SoundWAV{})
	buttonChanged.SoundClick()
}
