package process

import (
	"fmt"

	"github.com/pulpfree/univsales-pdf-url/awsservices"
	"github.com/pulpfree/univsales-pdf-url/config"
	"github.com/pulpfree/univsales-pdf-url/model"
)

// Process struct
type Process struct {
	cfg     *config.Config
	request *model.Request
}

// New function
func New(req *model.Request, cfg *config.Config) (p *Process, err error) {
	p = &Process{
		cfg:     cfg,
		request: req,
	}
	return p, err
}

// CreateURL method
func (p *Process) CreateURL() (url string, err error) {
	prefix, err := p.assemblePrefix()
	fmt.Printf("prefix %s\n", prefix)

	s3Svc, err := awsservices.NewS3(p.cfg)
	if err != nil {
		return "", err
	}

	return s3Svc.GetSignedURL(prefix)
}

//
// ======================== Helper Methods =============================== //
//

func (p *Process) assemblePrefix() (prefix string, err error) {
	tp := p.request.Type
	if tp == "quote" {
		prefix = fmt.Sprintf("/%s/qte-%d-r%d.pdf", tp, p.request.Number, p.request.Version)
	} else if tp == "invoice" {
		prefix = fmt.Sprintf("/%s/inv-%d.pdf", tp, p.request.Number)
	}
	return prefix, err
}
