package snapshot

import (
 "time"
 "MWF/internal/actions"
)
type Metadata struct { ID string `json:"id"`; Timestamp time.Time `json:"timestamp"`; Profile string `json:"profile"`; Actions []actions.Action `json:"actions"` }
// State records the value observed before MWF made each managed change.
type State struct { Records []Record `json:"records"` }
type Record struct { Action actions.Action `json:"action"`; Previous string `json:"previous,omitempty"`; Existed bool `json:"existed"`; Reversible bool `json:"reversible"` }
