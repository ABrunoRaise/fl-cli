// Copyright 2022 Giuseppe De Palma, Matteo Trentin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package admin_deploy_docker

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/funlessdev/fl-cli/pkg"
	"github.com/funlessdev/fl-cli/pkg/homedir"
	"github.com/funlessdev/fl-cli/pkg/log"
	"github.com/funlessdev/fl-cli/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDockerUpRun(t *testing.T) {
	homedirPath, err := os.MkdirTemp("", "funless-test-homedir-")
	require.NoError(t, err)

	defer func() {
		homedir.GetHomeDir = os.UserHomeDir
		os.RemoveAll(homedirPath)
	}()

	up := Up{}
	ctx := context.TODO()

	mockDockerShell := mocks.NewDockerShell(t)
	out, logger := testLogger()

	t.Run("should return error when setup fails", func(t *testing.T) {
		homedir.GetHomeDir = func() (string, error) {
			return "", errors.New("some home error")
		}
		err := up.Run(ctx, mockDockerShell, logger)
		require.Error(t, err)
	})

	t.Run("should complete successfully when no error occurs", func(t *testing.T) {
		homedir.GetHomeDir = func() (string, error) {
			return homedirPath, nil
		}
		mockDockerShell.On("ComposeUp", mock.Anything).Return(nil).Once()
		err := up.Run(ctx, mockDockerShell, logger)
		require.NoError(t, err)

		require.Contains(t, out.String(), "\nDeployment complete!")
	})

	t.Run("should return error when compose up fails", func(t *testing.T) {
		out.Reset()
		mockDockerShell.On("ComposeUp", mock.Anything).Return(errors.New("compose up error")).Once()
		err := up.Run(ctx, mockDockerShell, logger)
		require.Error(t, err)
	})

	t.Run("should modify docker-compose.yml when given custom core/worker", func(t *testing.T) {
		out.Reset()
		mockDockerShell.On("ComposeUp", mock.Anything).Return(nil).Once()
		_, path, err := homedir.ReadFromConfigDir("docker-compose.yml")
		require.NoError(t, err)
		os.Remove(path)

		up.CoreImage = "custom-core"
		up.WorkerImage = "custom-worker"
		err = up.Run(ctx, mockDockerShell, logger)
		require.NoError(t, err)

		require.Contains(t, out.String(), "\nDeployment complete!")
		content, _, err := homedir.ReadFromConfigDir("docker-compose.yml")
		require.NoError(t, err)
		require.Contains(t, string(content), "custom-core")
		require.Contains(t, string(content), "custom-worker")
	})
}

func Test_downloadDockerCompose(t *testing.T) {
	homedirPath, err := os.MkdirTemp("", "funless-test-homedir-")
	require.NoError(t, err)

	homedir.GetHomeDir = func() (string, error) {
		return homedirPath, err
	}
	defer func() {
		homedir.GetHomeDir = os.UserHomeDir
		os.RemoveAll(homedirPath)
	}()

	// Download it for the first time
	path, err := downloadDockerCompose()
	require.NoError(t, err)
	require.FileExists(t, path)

	// Now that it exists it should not give errors
	path, err = downloadDockerCompose()
	require.NoError(t, err)
	require.FileExists(t, path)
}

func Test_downloadPrometheusConfig(t *testing.T) {
	homedirPath, err := os.MkdirTemp("", "funless-test-homedir-")
	require.NoError(t, err)

	homedir.GetHomeDir = func() (string, error) {
		return homedirPath, err
	}
	defer func() {
		homedir.GetHomeDir = os.UserHomeDir
		os.RemoveAll(homedirPath)
	}()

	err = downloadPrometheusConfig()
	require.NoError(t, err)

	require.DirExists(t, filepath.Join(homedirPath, ".fl", "prometheus"))
	require.FileExists(t, filepath.Join(homedirPath, ".fl", "prometheus", "config.yml"))

	err = downloadPrometheusConfig()
	require.NoError(t, err)
}

func Test_replaceImages(t *testing.T) {
	homedirPath, err := os.MkdirTemp("", "funless-test-homedir-")
	require.NoError(t, err)

	homedir.GetHomeDir = func() (string, error) {
		return homedirPath, err
	}
	defer func() {
		homedir.GetHomeDir = os.UserHomeDir
		os.RemoveAll(homedirPath)
	}()

	t.Run("should return error when docker-compose.yml file is not found", func(t *testing.T) {
		err := replaceImages("core-test", "worker-test")
		require.Error(t, err)
	})

	t.Run("should swap core image when different from default", func(t *testing.T) {
		path, err := downloadDockerCompose()
		require.NoError(t, err)
		defer os.Remove(path)

		err = replaceImages("core-test", pkg.WorkerImg)
		require.NoError(t, err)

		content, _, err := homedir.ReadFromConfigDir("docker-compose.yml")
		require.NoError(t, err)

		expected := `  core:
    image: core-test`
		expectedWorker := `  worker:
    image: ghcr.io/funlessdev/worker:latest`
		require.Contains(t, string(content), expected, "core image should be the one provided")
		require.Contains(t, string(content), expectedWorker, "worker image should be the default")
	})

	t.Run("should swap worker image when different from default", func(t *testing.T) {
		path, err := downloadDockerCompose()
		require.NoError(t, err)
		defer os.Remove(path)

		err = replaceImages(pkg.CoreImg, "worker-test")
		require.NoError(t, err)

		content, _, err := homedir.ReadFromConfigDir("docker-compose.yml")
		require.NoError(t, err)

		expected := `  core:
    image: ghcr.io/funlessdev/core:latest`
		expectedWorker := `  worker:
    image: worker-test`
		require.Contains(t, string(content), expected, "core image should be the default")
		require.Contains(t, string(content), expectedWorker, "worker image should be the one provided")
	})

	t.Run("should swap both images when different from default", func(t *testing.T) {
		path, err := downloadDockerCompose()
		require.NoError(t, err)
		defer os.Remove(path)

		err = replaceImages("core-test", "worker-test")
		require.NoError(t, err)

		content, _, err := homedir.ReadFromConfigDir("docker-compose.yml")
		require.NoError(t, err)

		expected := `  core:
    image: core-test`
		expectedWorker := `  worker:
    image: worker-test`
		require.Contains(t, string(content), expected, "core image should be the one provided")
		require.Contains(t, string(content), expectedWorker, "worker image should be the one provided")
	})

}

func testLogger() (*bytes.Buffer, log.FLogger) {
	var outbuf bytes.Buffer
	testLogger, _ := log.NewLoggerBuilder().WithWriter(&outbuf).DisableAnimation().Build()
	return &outbuf, testLogger
}
