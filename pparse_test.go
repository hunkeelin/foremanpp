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
    ilist,err := Listinit("/data")
    var parr []Ppfile
    if err != nil {
        panic(err)
    }
    for _,i := range ilist{
        p,err := Capturevar(i)
        if err != nil {
            continue
        }
        parr = append(parr,p)
    }
    f := Payload {
        Ppfiles: parr,
        Env: "klintestenv",
    }
    err = json.NewEncoder(w).Encode(f)
    if err != nil {
        panic(err)
    }
    return
}
func TestUparse(t *testing.T){
    f := []string{"init.pp","ntp.pp","fuckyou.pp"}
    var parr []Ppfile
    for _,i := range f {
        p,err := Capturevar(i)
        if err != nil {
            fmt.Println(err)
            continue
        }
        parr = append(parr,p)
    }
    fmt.Println(parr)
}
func TestParse(t *testing.T){
    fmt.Println("testing parse")
    f,e := Listinit("/data")
    if e != nil {
        panic(e)
    }
    for _,i := range f{
        p,err := Capturevar(i)
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
    f,e := Listinit("/data")
    if e != nil {
        panic(e)
    }
    for _,i := range f{
        fmt.Println(i)
    }
}

