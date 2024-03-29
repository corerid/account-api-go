package main
import (
    "github.com/gorilla/mux"
    "net/http"
)
// Struct for a Route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
// Array of Route
type Routes []Route
// initiate the router
func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    // declare routes
    var routes = Routes{
        Route{
            "getIndex",
            "GET",
            "/accounts/{id}",
            accountFindById,
        },
        Route{
            "accounts",
            "GET",
            "/accounts",
            accounts,
        },
        Route{
            "insert",
            "POST",
            "/accounts/insert",
            insert,
        },
        Route{
            "update",
            "PUT",
            "/accounts/update",
            update,
        },
        Route{
            "delete",
            "DELETE",
            "/accounts/delete/{id}",
            delete,
        },
    }
    // bind routes
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }
    return router
}