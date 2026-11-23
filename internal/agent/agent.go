package agent

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Projectoutlast/ypracticum-metrics/internal/config"
	models "github.com/Projectoutlast/ypracticum-metrics/internal/model"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/httpclient"
	"github.com/Projectoutlast/ypracticum-metrics/pkg/logger"
)

const (
	_defaultPollInterval   = 2 * time.Second
	_defaultReportInterval = 10 * time.Second
)

type Collector interface {
	Collect(ctx context.Context) error
	Consume(ctx context.Context, id string, sent int64)
}

type Metric interface {
	Get(ctx context.Context) map[string]*models.Metrics
}

type Agent struct {
	*httpclient.HTTPClient

	serverAddr string

	reportInterval time.Duration
	pollInterval   time.Duration

	c      Collector
	logger logger.Interface
	m      Metric
}

func New(
	cfg *config.Config,
	client *httpclient.HTTPClient,
	l logger.Interface,
	c Collector,
	m Metric,
) *Agent {
	pollInterval := _defaultPollInterval
	reportInterval := _defaultReportInterval

	if cfg.Agent.PollInterval > 0 {
		pollInterval = cfg.Agent.PollInterval
	}

	if cfg.Agent.ReportInterval > 0 {
		reportInterval = cfg.Agent.ReportInterval
	}

	return &Agent{
		HTTPClient:     client,
		serverAddr:     cfg.Agent.ServerAddr,
		pollInterval:   pollInterval,
		reportInterval: reportInterval,
		c:              c,
		logger:         l,
		m:              m,
	}
}

func (a *Agent) Run(ctx context.Context) {
	go a.pollLoop(ctx)

	a.reportLoop(ctx)
}

func (a *Agent) pollLoop(ctx context.Context) {
	t := time.NewTicker(a.pollInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			a.logger.Debug("agent - Run - ctx.Done: context done")
			return
		case <-t.C:
			if err := a.collect(ctx); err != nil {
				a.logger.Error("agent - Run - a.collect: ", "error", err)

				continue
			}

			a.logger.Info("agent - Run - a.collect: success")
		}
	}
}

func (a *Agent) reportLoop(ctx context.Context) {
	t := time.NewTicker(a.reportInterval)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			a.logger.Debug("agent - Run - ctx.Done: context done")
			return
		case <-t.C:
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
		var (
			value string
			sent  int64
		)

		switch v.MType {
		case models.Counter:
			if v.Delta == nil || *v.Delta == 0 {
				continue
			}

			sent = *v.Delta
			value = strconv.FormatInt(sent, 10)
		case models.Gauge:
			if v.Value == nil {
				continue
			}

			value = strconv.FormatFloat(*v.Value, 'f', -1, 64)
		default:
			continue
		}

		url := fmt.Sprintf(
			"http://%s/update/%s/%s/%s",
			a.serverAddr, v.MType, v.ID, value,
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

			return fmt.Errorf("agent - update: status %d", resp.StatusCode)
		}

		if v.MType == models.Counter {
			a.c.Consume(ctx, v.ID, sent)
		}
	}

	return nil
}
