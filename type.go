package foremanpp
type payload struct {
    Ppfiles []ppfile  `json:"files"`
    Env string `json:"name"` // name of microservice
}

type ppfile struct {
    Classes []pclass `json:"classes"`
    Path string `json:"path"`
}

type pclass struct {
    Name   string        `json:"name"` // name of the class like klin::install
    Params []interface{} `json:"params"`
}

type pParams struct {
    Literal interface{} `json:"default_literal"`
    Source string  `json:"default_source"`
    Name string `json:"name"`
}
type varParams struct {
    Source string  `json:"default_source"`
    Name string `json:"name"`
}
