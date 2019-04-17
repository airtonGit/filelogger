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

// StartLog em arquivo.
//
// Cria arquivo logger.log ou concatena ao logger.log
func StartLog() {
	StartLogWithTag("~/logger.log", "default-tag ")
}

// StartLogWithTag em arquivo e tag.
//
// Inicia arquivo logger.log ou adiciona ao existente, permite
// também especificar string tag padrão no arquivo
func StartLogWithTag(logfile string, tag string) {

	var arquivoLog *os.File
	arquivoLog, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
