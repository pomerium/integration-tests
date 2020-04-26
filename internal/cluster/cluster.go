package cluster

import (
	"net/http"
)

type Cluster struct {
	workingDir string

	client *http.Client
	certs  *TLSCerts
}

func New(workingDir string) *Cluster {
	return &Cluster{
		workingDir: workingDir,
	}
}

func (cluster *Cluster) Do(req *http.Request) (res *http.Response, err error) {
	return cluster.client.Do(req)
}
