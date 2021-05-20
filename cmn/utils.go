package cmn

import (
	def "license/define"
	gcfg "license/gitconfig"
	hgcfg "license/hgconfig"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
)

//change fileformat from windows to unix
func CFS(name string) string {
	return ChangeFileStyle(name)
}

func ChangeFileStyle(name string) string {
	return strings.Replace(name, "\\", "/", -1)
}

func GName() string {
	if def.Args.FN != "" {
		return def.Args.FN
	}
	n := os.Getenv(def.NameEnv)
	if n != "" {
		return n
	}
	n, err := gcfg.Uname()
	if err == nil {
		return n
	}
	n, err = gcfg.Goba("user.name")
	if err == nil {
		return n
	}
	n, err = hgcfg.Uname()
	if err == nil {
		return n
	}
	usr, err := user.Current()
	if err == nil {
		return usr.Name
	}
	return ""
}

func GYear() string {
	if def.Args.FY != "" {
		return def.Args.FY
	}
	return strconv.Itoa(time.Now().Year())
}
