package broker

import (
	"fmt"
	"github.com/adrg/xdg"
	"github.com/itsabgr/go-handy"
	"github.com/valyala/fasthttp"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

const Version = "0.1.1"
const appName = "ilam"

var CertsCacheDir = ""

func init() {
	var err error
	CertsCacheDir, err = xdg.CacheFile(fmt.Sprintf("%s/%s/certs", appName, Version))
	handy.Throw(err)
}

var HostName = ""

func init() {
	var err error
	HostName, err = os.Hostname()
	handy.Throw(err)
}

var accessControlMaxAge = time.Hour * 24
var accessControlAllowMethods = []string{
	fasthttp.MethodGet,
	fasthttp.MethodPost,
	fasthttp.MethodOptions,
}

var accessControlMaxAgeString = strconv.Itoa(int(math.Floor(accessControlMaxAge.Seconds())))

var accessControlAllowMethodsString = strings.Join(accessControlAllowMethods, ", ")
