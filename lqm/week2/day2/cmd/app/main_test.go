package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tuannguyenandpadcojp/go-training/lqm/week2/day2/config"
)

func Test_loadConfig(t *testing.T) {
	tests := []struct {
		name             string
		prepareVariables func()
		want             config.Config
	}{
		{
			name:             "default config",
			prepareVariables: func() {},
			want: config.Config{
				PoolSize:              2,
				MaxJobs:               2,
				BannedNames:           map[string]struct{}{},
				WorkerPoolNonBlocking: false,
				HTTPPort:              8080,
			},
		},
		{
			name: "custom config",
			prepareVariables: func() {
				_ = os.Setenv("POOL_SIZE", "10")
				_ = os.Setenv("MAX_JOBS", "20")
				_ = os.Setenv("WORKER_POOL_NON_BLOCKING", "true")
				_ = os.Setenv("HTTP_PORT", "8081")
			},
			want: config.Config{
				PoolSize:              10,
				MaxJobs:               20,
				BannedNames:           map[string]struct{}{},
				WorkerPoolNonBlocking: true,
				HTTPPort:              8081,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got config.Config
			if tt.prepareVariables != nil {
				tt.prepareVariables()
				got = loadConfig("")
			} else {
				got = loadConfig("../../.env")
			}
			want := tt.want
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("loadConfig() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
