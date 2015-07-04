package ikafes

import (
	"http"
)

type Info struct{}

func (i *Info) Help() string {
	return ""
}

func (i *Info) Run(args []string) int {
	return 0
}

func (i *Info) Synopsis() string {
	return "fes information"
}

func getInfo() {
	r, e := http.Get(jsonUrl + "fes_info.json?" + Now())
	defer r.Body.Close()
}
