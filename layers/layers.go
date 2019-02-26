/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package layers

import (
	"fmt"
	"path/filepath"

	"github.com/buildpack/libbuildpack/internal"
	"github.com/buildpack/libbuildpack/logger"
)

// Layers represents the layers for an application.
type Layers struct {
	// Root is the path to the root directory for the layers.
	Root string

	logger logger.Logger
}

// Layer creates a Layer with a specified name.
func (l Layers) Layer(name string) Layer {
	metadata := filepath.Join(l.Root, fmt.Sprintf("%s.toml", name))
	return Layer{filepath.Join(l.Root, name), metadata, l.logger}
}

// String makes Layers satisfy the Stringer interface.
func (l Layers) String() string {
	return fmt.Sprintf("Layers{ Root: %s, logger: %s }", l.Root, l.logger)
}

// WriteApplicationMetadata writes application metadata to the filesystem.
func (l Layers) WriteApplicationMetadata(metadata Metadata) error {
	f := filepath.Join(l.Root, "launch.toml") // TODO: Remove once launch.toml removed from lifecycle

	l.logger.Debug("Writing application metadata: %s <= %s", f, metadata)
	if err := internal.WriteTomlFile(f, 0644, metadata); err != nil {
		return err
	}

	f = filepath.Join(l.Root, "app.toml")

	l.logger.Debug("Writing application metadata: %s <= %s", f, metadata)
	return internal.WriteTomlFile(f, 0644, metadata)
}

// NewLayers creates a new Logger instance.
func NewLayers(root string, logger logger.Logger) Layers {
	return Layers{root, logger}
}
