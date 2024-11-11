package usageDAL

import (
	"go-commons/config"
	"log"
)

func GetWebBuffer() uint32 {
	webBuffer := config.AppConfig.SourceWebBuffer
	log.Printf("Retrieved Web Buffer: %d", webBuffer)
	return webBuffer
}
