package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/info", handleInfo)
	http.HandleFunc("/epreuves", handleEpreuves)
	http.HandleFunc("/solution", handleSolution)
	http.HandleFunc("/classement", handleClassement)

	http.HandleFunc("/epreuves/epreuve1", handleEpreuve1)
	http.HandleFunc("/epreuves/epreuve2", handleEpreuve2)
	http.HandleFunc("/epreuves/epreuve3", handleEpreuve3)
	http.HandleFunc("/epreuves/epreuve4", handleEpreuve4)
	http.HandleFunc("/epreuves/epreuve5", handleEpreuve5)
	http.HandleFunc("/epreuves/epreuve6", handleEpreuve6)

	http.HandleFunc("/solutions/solution1", handleSolution1)
	http.HandleFunc("/solutions/solution2", handleSolution2)
	http.HandleFunc("/solutions/solution3", handleSolution3)
	http.HandleFunc("/solutions/solution4", handleSolution4)
	http.HandleFunc("/solutions/solution5", handleSolution5)
	http.HandleFunc("/solutions/solution6", handleSolution6)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))

	fmt.Println("Serveur démarré sur http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	renderTemplate(w, "html/index.html")
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/info.html")
}

func handleEpreuves(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/menuepreuve.html")
}

func handleSolution(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/menusolution.html")
}

func handleClassement(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/classement.html")
}

func handleEpreuve1(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve1.html")
}

func handleEpreuve2(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve2.html")
}

func handleEpreuve3(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve3.html")
}

func handleEpreuve4(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve4.html")
}

func handleEpreuve5(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve5.html")
}

func handleEpreuve6(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/epreuves/epreuve6.html")
}

func handleSolution1(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution1.html")
}

func handleSolution2(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution2.html")
}

func handleSolution3(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution3.html")
}

func handleSolution4(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution4.html")
}

func handleSolution5(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution5.html")
}

func handleSolution6(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "html/solutions/solution6.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
