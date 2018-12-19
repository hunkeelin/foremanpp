package foremanpp
type Payload struct {
    Ppfiles []Ppfile  `json:"files"`
    Env string `json:"name"` // name of microservice
}

type Ppfile struct {
    Classes []Pclass `json:"classes"`
    Path string `json:"path"`
}

type Pclass struct {
    Name   string        `json:"name"` // name of the class like klin::install
    Params []interface{} `json:"params"`
}

type PParams struct {
    Literal interface{} `json:"default_literal"`
    Source string  `json:"default_source"`
    Name string `json:"name"`
}
type Varparams struct {
    Source string  `json:"default_source"`
    Name string `json:"name"`
}
