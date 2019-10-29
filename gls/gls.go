// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gls implements a loader of OpenGL functions for the platform
// and a Go binding for selected OpenGL functions. The binding maintains
// some cached state to minimize the number of C function calls.
// The OpenGL function loader is generated by the "glapi2go" tool by
// parsing the OpenGL "glcorearb.h" header file
//
// This package also contains abstractions for some OpenGL object such as Program,
// Uniform, VBO and others.
package gls

import (
	"github.com/hecate-tech/engine/util/logger"
	"math"
	"unsafe"
)

// Package logger
var log = logger.New("GLS", logger.Default)

// Stats contains counters of WebGL resources being used as well
// the cumulative numbers of some WebGL calls for performance evaluation.
type Stats struct {
	Shaders    int    // Current number of shader programs
	Vaos       int    // Number of Vertex Array Objects
	Buffers    int    // Number of Buffer Objects
	Textures   int    // Number of Textures
	Caphits    uint64 // Cumulative number of hits for Enable/Disable
	UnilocHits uint64 // Cumulative number of uniform location cache hits
	UnilocMiss uint64 // Cumulative number of uniform location cache misses
	Unisets    uint64 // Cumulative number of uniform sets
	Drawcalls  uint64 // Cumulative number of draw calls
}

const (
	capUndef    = 0
	capDisabled = 1
	capEnabled  = 2
	uintUndef   = math.MaxUint32
	intFalse    = 0
	intTrue     = 1
)

const (
	FloatSize = int32(unsafe.Sizeof(float32(0)))
)
