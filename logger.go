package filelogger

import (
	"fmt"
	"log"
	"os"
)

type logHandler struct {
	*log.Logger
	filename string
	tag      string
}

//File Logger handler
var logInstance *logHandler

// IniciaLog em arquivo.
//
// Cria arquivo logger.log ou concatena ao logger.log
func IniciaLog() {
	IniciaLogWithTag("default-tag ")
}

// IniciaLogWithTag em arquivo e tag.
//
// Inicia arquivo logger.log ou adiciona ao existente, permite
// também especificar string tag padrão no arquivo
func IniciaLogWithTag(tag string) {

	var arquivoLog *os.File
	arquivoLog, err := os.OpenFile("logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Nao foi possivel escrever no arquivo erro1:", err)
	}

	logInstance = &logHandler{
		Logger:   log.New(arquivoLog, tag, log.Ldate|log.Lmicroseconds),
		filename: "logger.log",
		tag:      tag,
	}
}

//Warning adiciona nova linha no arquivo de log com rotulo WARNING
//
//Warning log line "DATETIME TAG WARNING"
func Warning(message string) {
	logInstance.Println("WARNING ", message)
}

//Info adiciona nova linha no log com rotulo INFO
//
//Info log line "DATETIME TAG INFO"
func Info(params ...interface{}) {
	logInstance.Println("INFO ", params)
}

//Error adiciona nova linha no arquivo de log
//
//message é inserida no arquivo de log com rotulo ERROR
func Error(params ...interface{}) {
	// var config struct {
	// 	itens interface{}
	// 	tag   string
	// }
	// config.itens = params
	// config.tag = "ERROR"
	// msgs := struct {
	// 	tag   string
	// 	itens interface{}
	// }{
	// 	"ERROR",
	// 	params,
	// }

	logInstance.Println("ERROR", params) ///"ERROR ", params)
}

//Debug adiciona nova linha no arquivo de log
//
//message é inserida no arquivo de log com rotulo DEBUG
//TODO: adicionar uma configuração por variavel de ambiente
//que permite ligar/desligar
func Debug(params ...interface{}) {
	logInstance.Println("DEBUG ", params)
}
