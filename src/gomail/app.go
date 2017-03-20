package main

import (
	"wc-server/actions"
	"github.com/drone/routes"
	"net/http"
	"os"
	"flag"
	"fmt"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}


func main() {

	argsWithProg := os.Args[1]

	mux := routes.New()

	mux.Get("/index", actions.Index)
	mux.Get("/upload", actions.Upload)
	mux.Post("/upload", actions.Upload)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	http.Handle("/", mux)


	http.ListenAndServe(":"+argsWithProg, nil)
	//s := &http.Server{
	//	Addr: ":"+argsWithProg,
	//	Handler: mux,
	//	TLSConfig: &tls.Config{
	//		//ClientCAs:  utils.LoadCA("/etc/pki/tls/certs/ca-bundle.crt"),
	//		//ClientCAs:  utils.LoadCA("/Users/wangpeng/keys/gpfs/ca.crt"),
	//		//ClientAuth: tls.RequireAnyClientCert,
	//	},
	//}
	//
	//e := s.ListenAndServeTLS("/etc/httpd/conf.d/ssl.crt", "/etc/httpd/conf.d/ssl.key")
	////e := s.ListenAndServeTLS("/Users/wangpeng/keys/gpfs/server.crt", "/Users/wangpeng/keys/gpfs/server.key")
	//if e != nil {
	//	log.Fatal("ListenAndServeTLS: ", e)
	//}


}
