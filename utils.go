package wordpress

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
	"net/http"
	"os"
)

var DEBUG bool = (os.Getenv("DEBUG") == "1")

func unmarshallResponse(resp gorequest.Response, body []byte, result interface{}) error {

	var prettyJSON bytes.Buffer
	err2 := json.Indent(&prettyJSON, body, "", "  ")
	if err2 != nil {
		log.Println("JSON parse error: ", err2)

		if DEBUG {
			log.Println("body: ", string(body))
		}
	}

	//	log.Println(string(prettyJSON.Bytes()))
	//	log.Println(resp.StatusCode)
	if DEBUG {
		log.Println(string(prettyJSON.Bytes()))
	}

	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusAccepted {
		log.Println("unmarshallResponse: ", resp.StatusCode)
		return errors.New(resp.Status)
	}

	err := json.Unmarshal(body, result)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return err
	}
	return nil
}

func warningNotImplemented(url string) {
	_warning("API not implemented yet: ", url)
}

func _warning(v ...interface{}) {
	_log(fmt.Sprintln("Warning: ", v))
}

func _log(v ...interface{}) {
	log.Println("[go-wordpress]", fmt.Sprintln(v))
}
