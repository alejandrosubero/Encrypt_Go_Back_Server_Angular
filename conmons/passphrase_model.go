package conmons

import (
	"log"
	"sync"
	"time"
)

// Singleton
type Singleton struct {
	Tiempo int64
	key    string
}

// Creador "estático"
var instancia *Singleton
var once sync.Once

func GetInstancia() *Singleton {
	once.Do(func() {
		instancia = &Singleton{
			key:    Passphrasekey,
			Tiempo: time.Now().Unix(),
		}
	})
	return instancia
}

func (s *Singleton) SetKey() {
	s.key = Passphrasekey
	log.Println("Actualiza la key en Passphrase")
}
