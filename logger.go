package filelogger

import (
	"fmt"
	"log"
	"os"
)

// IniciaLog em arquivo.
//
// Cria arquivo logger.log ou concatena ao logger.log
func IniciaLog() *log.Logger {
	return IniciaLogWithTag("default-tag ")
}

// IniciaLogWithTag em arquivo e tag.
//
// Inicia arquivo logger.log ou adiciona ao existente, permite
// também especificar string tag padrão no arquivo
func IniciaLogWithTag(tag string) *log.Logger {

	var arquivoLog *os.File
	arquivoLog, err := os.OpenFile("logger.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Nao foi possivel escrever no arquivo erro1:", err)
	}

	logger := log.New(arquivoLog, tag, log.Ldate|log.Lmicroseconds)

	return logger
}
