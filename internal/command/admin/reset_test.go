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

package admin

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/funlessdev/fl-cli/pkg/log"
	"github.com/funlessdev/fl-cli/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestResetRun(t *testing.T) {
	reset := reset{}
	ctx := context.TODO()

	var outbuf bytes.Buffer
	testLogger, _ := log.NewLoggerBuilder().WithWriter(&outbuf).DisableAnimation().Build()

	deployer := mocks.NewDevDeployer(t)

	t.Run("print error when setup client fails", func(t *testing.T) {
		deployer.On("Setup", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(func(ctx context.Context, coreImg string, workerImg string) error {
				return errors.New("error")
			}).Once()

		_ = reset.Run(ctx, deployer, testLogger)

		expectedOutput := []string{
			"Removing local funless deployment...\n",
			"\n",
			"",
		}

		assertOutput(t, expectedOutput, &outbuf)
	})

	t.Run("print error when docker networks setup fails", func(t *testing.T) {
		deployer.On("Setup", ctx, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
			Return(func(ctx context.Context, coreImg string, workerImg string) error {
				return nil
			})
		deployer.On("RemoveCoreContainer", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_ = reset.Run(ctx, deployer, testLogger)

		expectedOutput := []string{
			"Removing local funless deployment...\n",
			"\n",
			"Removing Core container... ☠️\n",
			"failed\n",
			"",
		}

		assertOutput(t, expectedOutput, &outbuf)
	})

	t.Run("print error when pulling core image fails", func(t *testing.T) {
		deployer.On("RemoveCoreContainer", ctx).Return(func(ctx context.Context) error {
			return nil
		})
		deployer.On("RemoveWorkerContainer", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_ = reset.Run(ctx, deployer, testLogger)

		expectedOutput := []string{
			"Removing local funless deployment...\n",
			"\n",
			"Removing Core container... ☠️\n",
			"done\n",
			"Removing Worker container... 🔪\n",
			"failed\n",
			"",
		}

		assertOutput(t, expectedOutput, &outbuf)
	})

	t.Run("print error when pulling worker image fails", func(t *testing.T) {
		deployer.On("RemoveWorkerContainer", ctx).Return(func(ctx context.Context) error {
			return nil
		})
		deployer.On("RemoveFLNetwork", ctx).Return(func(ctx context.Context) error {
			return errors.New("error")
		}).Once()

		_ = reset.Run(ctx, deployer, testLogger)

		expectedOutput := []string{
			"Removing local funless deployment...\n",
			"\n",
			"Removing Core container... ☠️\n",
			"done\n",
			"Removing Worker container... 🔪\n",
			"done\n",
			"Removing fl network... ✂️\n",
			"failed\n",
			"",
		}

		assertOutput(t, expectedOutput, &outbuf)
	})

	t.Run("successful prints when everything goes well", func(t *testing.T) {
		deployer.On("RemoveFLNetwork", ctx).Return(func(ctx context.Context) error {
			return nil
		})

		_ = reset.Run(ctx, deployer, testLogger)

		expectedOutput := []string{
			"Removing local funless deployment...\n",
			"\n",
			"Removing Core container... ☠️\n",
			"done\n",
			"Removing Worker container... 🔪\n",
			"done\n",
			"Removing fl network... ✂️\n",
			"done\n",
			"\n",
			"All clear! 👍\n",
			"",
		}

		assertOutput(t, expectedOutput, &outbuf)
	})

}
