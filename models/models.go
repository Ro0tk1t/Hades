package models


type Target struct{
    IP string
    Port int
    URL string
    Protocol string
}


type Result struct{
    User string
    Pass string
    Target Target
}
