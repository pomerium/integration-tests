package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/calebdoxsey/pomerium-integration-tests/internal/cluster"
	"github.com/onsi/gocleanup"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const localHTTPSPort = 9443

var (
	mainCtx     context.Context
	testcluster *cluster.Cluster
)

func TestMain(m *testing.M) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	mainCtx = context.Background()
	var cancel func()
	mainCtx, cancel = context.WithCancel(mainCtx)
	var clearTimeout func()
	mainCtx, clearTimeout = context.WithTimeout(mainCtx, time.Minute*10)
	defer clearTimeout()

	_, mainTestFilePath, _, _ := runtime.Caller(0)
	testcluster = cluster.New(filepath.Dir(mainTestFilePath))
	if err := testcluster.Setup(mainCtx); err != nil {
		log.Fatal().Err(err).Send()
	}

	status := m.Run()
	cancel()
	gocleanup.Cleanup()
	os.Exit(status)
}
