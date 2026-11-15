package agent

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Projectoutlast/ypracticum-metrics/internal/service"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/httpclient"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

const (
	_defaultPollInterval   = 2 * time.Second
	_defaultReportInterval = 10 * time.Second
)

type Agent struct {
	*httpclient.HTTPClient

	c      service.Collector
	logger logger.Interface
	m      service.Metric
}

func New(
	client *httpclient.HTTPClient,
	l logger.Interface,
	c service.Collector,
	m service.Metric,
) *Agent {
	return &Agent{HTTPClient: client, c: c, logger: l, m: m}
}

func (a *Agent) Run(ctx context.Context) {
	go a.pollLoop(ctx)

	a.reportLoop(ctx)
}

func (a *Agent) pollLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			a.logger.Debug("agent - Run - ctx.Done: context done")
			return
		case <-time.After(_defaultPollInterval):
			if err := a.collect(ctx); err != nil {
				a.logger.Error("agent - Run - a.collect: ", "error", err)

				continue
			}

			a.logger.Info("agent - Run - a.collect: success")
		}
	}
}

func (a *Agent) reportLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			a.logger.Debug("agent - Run - ctx.Done: context done")
			return
		case <-time.After(_defaultReportInterval):
			if err := a.update(ctx); err != nil {
				a.logger.Error("agent - Run - a.update: ", "error", err)
			}

			a.logger.Info("agent - Run - a.update: success")
		}
	}
}

func (a *Agent) collect(ctx context.Context) error {
	if err := a.c.Collect(ctx); err != nil {
		return err
	}

	return nil
}

func (a *Agent) update(ctx context.Context) error {
	m := a.m.Get(ctx)

	for _, v := range m {
		url := fmt.Sprintf(
			"http://localhost:8080/update/%s/%s/%s",
			v.MType, v.ID, strconv.Itoa(int(*v.Value)),
		)

		req, err := http.NewRequest(http.MethodPost, url, nil)
		if err != nil {
			a.logger.Error("agent - update - http.NewRequest: ", "error", err)

			return err
		}

		req.Header.Set("Content-Type", "text/plain")

		resp, err := a.Client.Do(req)
		if err != nil {
			a.logger.Error("agent - update - a.Client.Do: ", "error", err)

			return err
		}

		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			a.logger.Error("agent - update - response: code %d", resp.StatusCode)
		}
	}

	return nil
}
