// Created: 2025-04-24
package net

import (
	"time"

	"github.com/lunar-parklife/did"
)

type Cache []CacheResult

type CacheResult struct {
	Document  *did.Document
	Handle    string
	Retrieved time.Time
}
