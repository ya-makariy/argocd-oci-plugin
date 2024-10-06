package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func VerboseToStdErr(format string, message ...interface{}) {
	if viper.GetBool("verboseOutput") {
		log.Printf(fmt.Sprintf("%s\n", format), message...)
	}
}
