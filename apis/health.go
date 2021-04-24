package apis

import (
	"fmt"
	"net/http"
)

func GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey, who's there?")
}
