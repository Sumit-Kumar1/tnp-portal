package models

type DateTime struct {
	Value    *string `json:"value,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}
