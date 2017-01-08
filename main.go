package sendingEmail

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
	"bufio"
	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/mail"
	netMail "net/mail"
)


func init() {

	http.HandleFunc("/",serveStaticPages)
	http.HandleFunc("/public/css/bootstrap.css",serveResource)
	http.HandleFunc("/public/js/jquery.js",serveResource)
	http.HandleFunc("/public/js/bootstrap.js",serveResource)
	http.HandleFunc("/send",serveSendEmail)
}

// hello is an HTTP handler that prints "Hello Gopher!"
func serveStaticPages(w http.ResponseWriter, r *http.Request) {
	template:= template.Must(template.ParseFiles("bin/pages/index.html"))
	if err := template.ExecuteTemplate(w,"index.html",nil);err !=nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}

}
func serveResource(w http.ResponseWriter, r *http.Request)  {
	//fmt.Println("Hello from serve function -------------------------------------------")
	path:= "./bin"+r.URL.Path
	var contenttype  string
	if strings.HasSuffix(path, ".css") {
		contenttype = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contenttype = "image/png; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contenttype = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contenttype = "application/javascript; charset=utf-8"
	} else {
		contenttype = "text/plain; charset=utf-8"
	}
	//fmt.Println("path of resources = -----",path)
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contenttype)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		fmt.Println(err.Error())
		w.WriteHeader(404)
	}

}
func serveSendEmail(w http.ResponseWriter, r *http.Request)  {
	c := appengine.NewContext(r)
	email := r.FormValue("email")
	subject := "HELLO THIS IS  FIRST MAIL"
	message := "we try to send cv's"
	msg := &mail.Message{
		Sender:  "mostafamowafy93@gmail.com",
		To:      []string{"m_mowafy_1993@hotmail.com"},
		ReplyTo: email,
		Subject: subject,
		Body:    message,
		Headers: netMail.Header{
			"On-Behalf-Of": []string{email},
		},
	}
	fmt.Println(msg)
	if err := mail.Send(c, msg); err != nil {
		fmt.Fprint(w, "Mail NOT send! Error")
	}else{
		fmt.Fprint(w, "Mail send.")
	}
}