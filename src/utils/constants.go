package utils

//FIXME, this will be removed later
const GLOBAL_SEED = "top secret"

var API_PROXY_MAP = map[string]string{
	"PROXY_DATASOURCE_DJANGO_API": "http://192.168.7.11:8000",
	"PROXY_DATASHARE_NODEJS_API":  "http://192.168.7.162:8090",
}

//FIXME. need to fix
var ROLEACC_DATASOURCE_DJANGO = map[string]bool{
	"admin": true,
	"user":  true,
}

var ROLEACC_DATASHARE_NODEJS = map[string]bool{
	"admin": true,
	"user":  false,
}
