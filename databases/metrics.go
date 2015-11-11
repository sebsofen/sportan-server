package databases

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type MetricApi struct {
	host     string
	database string
	password string
	user     string
}

func NewMetricApi(host string, database string, user string, password string) *MetricApi {
	return &MetricApi{
		host:     host,
		database: database,
		password: password,
		user:     user,
	}
}

//PostDataPoint post new points to series
func (m *MetricApi) PostDataPoint(series string, value int) {
	var jsonStr = []byte(`[{"name":"` + series + `","columns":["value"],"points":[[` + strconv.Itoa(value) + `]]}]`)
	url := m.host + "/db/" + m.database + "/series?u=" + m.user + "&p=" + m.password + ""
	fmt.Println(url)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body)
}
