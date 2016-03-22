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
	if DEBUG {
		log.Println("body: ", string(prettyJSON.Bytes()))
	}

	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusAccepted {
		return errors.New(resp.Status)
	}

	err := json.Unmarshal(body, result)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return err
	}
	return nil
}

func _warning(v ...interface{}) {
	log.Println(fmt.Sprintln("[go-wordpress]", v))
}

func _log(v ...interface{}) {
	log.Println(fmt.Sprintln("[go-wordpress]", v))
}

// UnmarshallServerError A helper function to unmarshall error response from server
func UnmarshallServerError(body []byte) ([]GeneralError, error) {
	var resp []GeneralError
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Println("JSON parse error: ", err)
		return nil, err
	}
	return resp, nil
}
