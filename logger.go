package filelogger

import (
	"fmt"
	"log"
	"os"
)

//Filelogger mantém arquivo e metodos para log
type Filelogger struct {
	*log.Logger
	filename string
	tag      string
}

// StartLog em arquivo.
//
// Cria arquivo logger.log ou concatena ao logger.log
func StartLog() (*Filelogger, error) {
	return StartLogWithTag("~/logger.log", "default-tag ")
}

// StartLogWithTag em arquivo e tag.
//
// Inicia arquivo logger.log ou adiciona ao existente, permite
// também especificar string tag padrão no arquivo
func StartLogWithTag(logfile string, tag string) (*Filelogger, error) {

	var arquivoLog *os.File
	arquivoLog, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("Nao foi possivel escrever no arquivo erro1:%s", err.Error())
	}

	logInstance := &Filelogger{
		Logger:   log.New(arquivoLog, tag, log.Ldate|log.Lmicroseconds),
		filename: logfile,
		tag:      tag,
	}
	return logInstance, nil
}

//Warning adiciona nova linha no arquivo de log com rotulo WARNING
//Warning log line "DATETIME TAG WARNING"
func (l *Filelogger) Warning(message string) {
	l.Println("WARNING ", message)
}

//Info adiciona nova linha no log com rotulo INFO
//
//Info log line "DATETIME TAG INFO"
func (l *Filelogger) Info(params ...interface{}) {
	l.Println("INFO ", params)
}

//Error adiciona nova linha no arquivo de log
//
//message é inserida no arquivo de log com rotulo ERROR
func (l *Filelogger) Error(params ...interface{}) {
	l.Println("ERROR", params) ///"ERROR ", params)
}

//Debug adiciona nova linha no arquivo de log
//
//TODO: adicionar uma configuração por variavel de ambiente
//que permite ligar/desligar
func (l *Filelogger) Debug(params ...interface{}) {
	l.Println("DEBUG ", params)
}
