package dto

type ExecutionTimes struct {
	// in nanoseconds
	TotalTime  uint64 `json:"totalTime"`
	TotalCalls uint32 `json:"totalCalls"`
}
