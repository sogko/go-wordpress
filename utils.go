package wordpress

import (
	"fmt"
	"log"
)

func _warning(v ...interface{}) {
	log.Println(fmt.Sprintln("[go-wordpress]", v))
}
