package actions

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Version: %s", "v201703191120")
}
