package dto

type ActiveClient struct {
	Metadata     *ClientMetadata
	InactiveTime int64
}
