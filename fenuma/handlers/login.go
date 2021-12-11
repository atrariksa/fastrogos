package handlers

// import (
// 	"net/http"
// 	"text/template"

// 	"github.com/sirupsen/logrus"
// )

// type LoginPage struct {
// 	Title string
// }

// type LoginHandler struct {
// 	log *logrus.Logger
// }

// func NewLoginHandler(log *logrus.Logger) *LoginHandler {
// 	return &LoginHandler{
// 		log: log,
// 	}
// }

// func (lp *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
// 	loginPage := LoginPage{Title: "Login"}
// 	err := lp.renderTemplate(w, "templates/login/login", loginPage)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 	}
// }

// func (lp *LoginHandler) renderTemplate(w http.ResponseWriter, tmpl string, p LoginPage) error {
// 	t, err := template.ParseFiles(tmpl + ".html")
// 	if err != nil {
// 		return err
// 	}

// 	return t.Execute(w, p)
// }
