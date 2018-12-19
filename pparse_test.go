package foremanpp
import(
    "testing"
    "encoding/json"
    "github.com/hunkeelin/mtls/klinserver"
    "net/http"
    "fmt"
)
func TestServer(t *testing.T){
    con := http.NewServeMux()
    con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        testweb(w, r)
    })
    j := &klinserver.ServerConfig{
        BindPort: "2018",
        BindAddr: "",
        ServeMux: con,
    }
    err := klinserver.Server(j)
    if err != nil {
        panic(err)
    }
}
func testweb(w http.ResponseWriter, r *http.Request) {
    ilist,err := listinit("/data")
    var parr []ppfile
    if err != nil {
        panic(err)
    }
    for _,i := range ilist{
        p,err := capturevar(i)
        if err != nil {
            continue
        }
        parr = append(parr,p)
    }
    f := payload {
        Ppfiles: parr,
        Env: "klintestenv",
    }
    err = json.NewEncoder(w).Encode(f)
    if err != nil {
        panic(err)
    }
    return
}
func TestParse(t *testing.T){
    fmt.Println("testing parse")
    f,e := listinit("/data")
    if e != nil {
        panic(e)
    }
    for _,i := range f{
        p,err := capturevar(i)
        if err != nil {
            continue
        }
        fmt.Println(p)
    }
}
/*
func TestSmart(t *testing.T){
    p,err := Smartclassvar("/data")
    if err != nil {
        panic(err)
    }
    for _,i := range p{
        fmt.Println(i.Class)
    }
}
*/
func TestList(t *testing.T){
    f,e := listinit("/data")
    if e != nil {
        panic(e)
    }
    for _,i := range f{
        fmt.Println(i)
    }
}

