package api

import (
	"fmt"

	"github.com/divyanshkul/temp-mod/constants"
)

func CallAPI() {
	fmt.Println("API called successfully")
	fmt.Println("Log Level", constants.LOG_LEVEL)
}
