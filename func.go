package foremanpp
import(
    "bytes"
    "io/ioutil"
    "fmt"
    "strings"
    "strconv"
    "regexp"
    "path/filepath"
    "errors"
)
func Listinit(rootdir string)([]string,error){ // list the init in each directory
    var toreturn []string
    if len(rootdir) == 0 {
        return toreturn,errors.New("Please specify puppet env dir")
    }
    last := rootdir[len(rootdir)-1:]
    if last != "/" {
        rootdir = rootdir + "/"
    }
    f,err := filepath.Glob(rootdir+"*")
    if err != nil {
        return toreturn,errors.New("Unable to glob the rootdir, does the rootdir exist?" + rootdir)
    }
    for idx,ele := range f {
        f[idx] = ele + "/manifests/init.pp"
    }
    return f,nil
}
func removecomment(f []byte) []byte{
    g := bytes.Split(f,[]byte("\n"))
    for i,ele := range g {
        if len(ele) > 1 {
            if ele[0] == 35 {
                g[i] = []byte("")
            }
        }
    }
    return bytes.Join(g,[]byte("\n"))
}
func Capturevar(s string)(Ppfile,error){
    var toreturn Ppfile
    var returnclass []Pclass
    var returnparams []PParams
    var returnparamsp []Varparams
    var tmPclass Pclass
    g,err := ioutil.ReadFile(s)
    if err != nil {
        return toreturn,errors.New("No such file " + s)
    }
    f := removecomment(g)
    reducespace := bytes.Replace(f,[]byte(" "),[]byte(""),-1)
    reducelines := bytes.Replace(reducespace,[]byte("\n"),[]byte(""),-1)
    capture := regexp.MustCompile("class.*?\\((.*?)\\)")
    class := regexp.MustCompile("class(.*?)\\(")
    capturecontent := capture.FindSubmatch(reducelines)
    captureclass := class.FindSubmatch(reducelines)
    if len(captureclass) != 2 {
        return toreturn,errors.New("no aval class params")
    }
    if len(capturecontent) != 2 {
        return toreturn,errors.New("no aval content")
    }
    if string(captureclass[1]) == "" {
        return toreturn,errors.New("No class info for "+s)
    }
    contents := bytes.Split(capturecontent[1],[]byte(","))
    var a,b int
    var ab,bb bool
    for i,ele := range reducelines {
        if ele == 123 && !ab{ //123 = {
            a = i
            ab = true
        }
        if ele == 40 && !bb { //40 = (
            b = i
            bb = true
        }
    }
    if a < b {
        class = regexp.MustCompile("class(.*?)\\{")
        captureclass = class.FindSubmatch(reducelines)
        tmPclass.Name = string(captureclass[1])
        fakeint := make([]interface{},1)
        fakeparams := Varparams {
            Source: "N/A",
            Name: "N/A",
        }
        fakeint[0] = fakeparams
        tmPclass.Params = fakeint
        returnclass = append(returnclass,tmPclass)
        toreturn.Classes = returnclass
        fmt.Println("No class info aval for this class " + s)
        return toreturn,nil
    }
    tmPclass.Name = string(captureclass[1])
    for _,i := range contents {
        param := bytes.Split(i,[]byte("="))
        if len(param) != 2 {
            continue
        }
        var sswitch bool
        var slength int
        if strings.HasPrefix(string(param[1]),"$"){
            var tmpparam Varparams
            tmpparam.Name = string(param[0][1:])
            tmpparam.Source  = string(param[1])
            returnparamsp = append(returnparamsp,tmpparam)
            sswitch = true
            slength = len(returnparamsp)
        } else{
            var tmpparam PParams
            tmpparam.Name = string(param[0][1:])
            tmpparam.Source  = string(param[1])
            switch{
            case tmpparam.Source == strings.ToLower("true"):
                tmpparam.Literal = true
            case tmpparam.Source == strings.ToLower("false"):
                tmpparam.Literal = false
            default:
                if ij,err := strconv.Atoi(tmpparam.Source); err == nil {
                    tmpparam.Literal = ij
                } else {
                    tmpparam.Literal = string(bytes.Trim(param[1],"\""))
                }
            }
            returnparams = append(returnparams,tmpparam)
            sswitch = false
            slength = len(returnparams)
        }
        if sswitch {
            rparamsInt := make([]interface{},slength)
            for i := range returnparamsp {
                rparamsInt[i] = returnparamsp[i]
            }
            tmPclass.Params = rparamsInt
        } else {
            rparamsInt := make([]interface{},slength)
            for i := range returnparams {
                rparamsInt[i] = returnparams[i]
            }
            tmPclass.Params = rparamsInt
        }
    }
    if len(tmPclass.Params) == 0 {
        fakeint := make([]interface{},1)
        fakeparams := Varparams {
            Source: "N/A",
            Name: "N/A",
        }
        fakeint[0] = fakeparams
        tmPclass.Params = fakeint
    }
    returnclass = append(returnclass,tmPclass)
    toreturn.Classes = returnclass
    return toreturn,nil
}
func shit(){
    fmt.Println("asdf")
}
/*
func Smartclassvar(rootdir string)([]Puppetclass,error){
    var toreturn []Puppetclass
    f,err := Listinit(rootdir)
    if err != nil {
        return toreturn,err
    }
    for _,i := range f{
        p,err := Capturevar(i)
        if err == nil {
            toreturn = append(toreturn,p)
        }else {
            fmt.Println(err)
        }
    }
    return toreturn,nil
}
*/
