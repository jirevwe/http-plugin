package http_plugin

import (
	"fmt"
	"github.com/jirevwe/plug"
	"time"
)

func init() {
	plug.RegisterModule(Http{})
}

const ID plug.ModuleID = "http.logger"

type Http struct{}

func (h Http) Emit(value any) error {
	fmt.Printf("[HTTP]: %v", value.(time.Time).Unix())
	return nil
}

func (h Http) ModuleInfo() plug.ModuleInfo {
	return plug.ModuleInfo{
		ID:  ID,
		New: func() plug.Module { return new(Http) },
	}
}

// var _ plug.Validator = (Http)(nil)
var _ plug.Emitter = (*Http)(nil)
