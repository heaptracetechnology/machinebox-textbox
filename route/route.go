package route

import (
    "github.com/gorilla/mux"
    textbox "github.com/heaptracetechnology/machinebox-textbox/textbox"
    "log"
    "net/http"
)

//Route struct
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//Routes array
type Routes []Route

var routes = Routes{
    Route{
        "TextAnalyze",
        "POST",
        "/textAnalyze",
        textbox.TextAnalyze,
    },
}

//NewRouter route
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
