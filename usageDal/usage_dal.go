package usageDAL

import (
	"log"

	"github.com/Aashirwad-Chauhan/go-commons/config"
)

func GetWebBuffer() uint32 {
	webBuffer := config.AppConfig.SourceWebBuffer
	log.Printf("Retrieved Web Buffer: %d", webBuffer)
	return webBuffer
}
