package video

import (
	"encoding/json"
	"fmt"
	error2 "github.com/pkg/errors"
	"net/http"
	"tryon01/net/once/service/video"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	result, err := video.QueryAll()
	if err != nil {
		fmt.Printf("%+v", error2.Unwrap(err))
		return
	}

	bs, err := json.Marshal(result)
	if err != nil {
		fmt.Println(error2.Cause(err))
		return
	}

	fmt.Fprintf(w, string(bs))
}
