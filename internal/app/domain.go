package app

import (
	"github.com/nachoconques0/ethtxfetcher/internal/controller"
	"github.com/nachoconques0/ethtxfetcher/internal/parser"
	"github.com/nachoconques0/ethtxfetcher/internal/service"
)

// setupDomain uses dependency injections required for start up our business domain, the outcome
// will be a ready domain service or an error
func (a *Application) setupDomain() error {
	ethParser := parser.NewParser(a.nodeURL)
	parserSvc, err := service.NewParserSvc(ethParser)
	if err != nil {
		return err
	}
	a.parserSvc = parserSvc

	parserCtrl := controller.NewParserCtrl(a.parserSvc)
	a.parserCtrl = parserCtrl
	return nil
}
