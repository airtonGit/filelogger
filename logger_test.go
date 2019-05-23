package filelogger

import "testing"

func TestInfo(t *testing.T) {
	logger, err := StartLogWithTag("arquiv-log-teste.log", "ms-teste-tag ")
	if err != nil {
		panic("Não foi possivel inicial filelogger")
	}
	logger.Info("Exemplo de linha de info")
}
