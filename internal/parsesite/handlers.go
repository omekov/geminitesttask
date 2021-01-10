package parsesite

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/omekov/geminitesttask/internal/parsesite/geocoder"
	"github.com/omekov/geminitesttask/internal/parsesite/scanners"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/branchs", handlerBranch)
	// mux.HandleFunc("/staff", handlerStaff)
	// mux.HandleFunc("/geo", handlerGeo)
	return mux
}

func handlerBranch(w http.ResponseWriter, r *http.Request) {
	branchs, err := scanners.Branch()
	if err != nil {
		log.Println("failed to scanner branch:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, b := range branchs {
		l, err := geocoder.Geocoding(b.Address)
		if err != nil {
			log.Println("failed to serialize response:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(l)
		branchs[i].Location = l
	}
	b, err := json.Marshal(branchs)
	if err != nil {
		log.Println("failed to serialize response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

// test ..
func handlerGeo(w http.ResponseWriter, r *http.Request) {
	l, err := geocoder.Geocoding("1121 Bedford Avenue  Brooklyn, NY 11216")
	if err != nil {
		log.Println("failed to serialize response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(l)
}

// test ..
func handlerStaff(w http.ResponseWriter, r *http.Request) {
	staff, err := scanners.Staff("https://ymcanyc.org/locations/ridgewood-ymca/about")
	if err != nil {
		log.Println("failed to scanner branch:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(staff)
	if err != nil {
		log.Println("failed to serialize response:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}
