package actions

import (
"fmt"
"net/http"
)

// To show all the usage for REST api.
func Help(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Version: %s", "v201703191120")
}

