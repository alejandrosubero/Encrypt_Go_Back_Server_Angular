package conmons

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	openssl "github.com/Luzifer/go-openssl/v4"
	"github.com/jasonlvhit/gocron"
)

/*
LAS FUNCIONES EN ESTE CONTROLLER PERMITEN GENERAR UNA KEY PARA ENCRIPTAR LA LLAMADA DESDE EL FRONT AL BACK,
RECIBIR LA LLAMADA CON EL CUERPO ENCRIPTADO EL CUAL SE DESENCRIPTA CON LOS ALGORIHTMOS ASE, PARA LUEGO
SERELIZARLO AL OBJETO.

ESTO SE UNE CON EL MODELO DE SINGLETON QUE PREMITIRA MENTENER UNA SOLA LLAVE PARA TODOAS LAS REQUEST LA LLAVE CAMBIARA
CADA DIA.
*/

var bodyRecibe string
var Passphrasekey string

func CodeRest(writer http.ResponseWriter, request *http.Request) {
	log.Println("inicio de la llamada al server ")

	DataRecive := Data{}

	DataRecive = CodeDecoRequest(writer, request)

	log.Printf(" $$$ Decodificado en Objeto %v: ", DataRecive)

	jsonx, _ := json.Marshal(DataRecive)

	fmt.Printf("objeto : => %s", jsonx)

	SendResponse(writer, http.StatusCreated, jsonx)

	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR: IN NEWDECODER")
			log.Println("Panic to Recover occurred:", err)
			errorSistenResponse, er := json.Marshal(err)
			if er != nil {
				fmt.Println("error:", er)
			}
			SendErrorMensaje(writer, http.StatusBadRequest, errorSistenResponse)
		}
		log.Println("")
		log.Println("***********************************************************************")
	}()
}

func CodeDecoRequest(writer http.ResponseWriter, request *http.Request) Data {
	log.Println("inicio de la llamada al server ")

	DataRecive := Data{}
	// Passphrasekey = "8439db0f0c86d30e9bc6d43743e73b605008160b9af9cf973ead223644a3fa5c"

	err := json.NewDecoder(request.Body).Decode(&bodyRecibe)
	if err != nil {
		log.Println("ERROR: IN NEWDECODER")
		log.Println(err)
		SendError(writer, http.StatusBadRequest)
	}

	log.Printf("Request Body %v", bodyRecibe)

	decryptingMensaje := DecryptOpensslEncrypted(bodyRecibe, Passphrasekey)

	log.Println("decodificado: " + decryptingMensaje)

	errObjete := json.Unmarshal([]byte(decryptingMensaje), &DataRecive)

	log.Printf("Decodificado en Objeto %v: ", DataRecive)

	if errObjete != nil {
		log.Println("ERROR: json.Marshal(DataRecive)")
		log.Println(errObjete)
		SendError(writer, http.StatusBadRequest)
	}
	return DataRecive
}

func DecryptOpensslEncrypted(opensslEncrypted string, passphrase string) string {

	//go get github.com/Luzifer/go-openssl/v4
	// go get github.com/Luzifer/go-openssl

	o := openssl.New()
	dec, err := o.DecryptBytes(passphrase, []byte(opensslEncrypted), openssl.BytesToKeyMD5)
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
	decrypt := string(dec)
	log.Println("decrypt")
	return decrypt
}

// A Golang Job Scheduling Package.// go get github.com/jasonlvhit/gocron
func Scheduler() {
	// go get github.com/jasonlvhit/gocron

	s := gocron.NewScheduler()
	//s.Every(30).Seconds().Do(task)
	// s.Every(2).Minutes().Do(task)
	s.Every(2).Day().Do(task)
	<-s.Start()
}

//Scheduling task
func task() {
	//generate a random 32 byte key for AES-256
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	Passphrasekey = key

	objetoKey := GetInstancia()
	objetoKey.SetKey()

	log.Println("")
	log.Printf("key ==>: %s", key)
}

func Passphrase(writer http.ResponseWriter, request *http.Request) {

	if Passphrasekey == "" {
		task()
		log.Println("corre task desde func Passphrase key")
	}

	keyEncryptkey := "8439db0f0c86d30e9bc6d43743e73b605008160b9af9cf973ead223644a3fa5c"
	objetoKey := GetInstancia()
	// log.Printf("GENERADA:========> %s", objetoKey.key)

	encrit := encrypt(keyEncryptkey, objetoKey.key)
	// jsonx, err := json.Marshal(objetoKey.key)
	jsonx, err := json.Marshal(encrit)
	if err != nil {
		log.Println("ERROR: objetoKey.key")
		log.Println(err)
		SendError(writer, http.StatusBadRequest)
	}
	//	log.Printf("func Passphrase key:=> %s", jsonx)

	SendResponse(writer, http.StatusCreated, jsonx)
}

func encrypt(passphrase string, plaintext string) string {
	o := openssl.New()
	enc, err := o.EncryptBytes(passphrase, []byte(plaintext), openssl.BytesToKeyMD5)
	if err != nil {
		//	fmt.Printf("An error occurred: %s\n", err)
	}
	// fmt.Printf("Encrypted text: %s\n", string(enc))
	return string(enc)
}
