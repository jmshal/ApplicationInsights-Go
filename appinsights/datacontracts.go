package appinsights

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type Telemetry interface {
	Timestamp() time.Time
	Context() TelemetryContext
	baseTypeName() string
	baseData() Domain
}

type BaseTelemetry struct {
	timestamp time.Time
	context   TelemetryContext
}

type TraceTelemetry struct {
	BaseTelemetry
	data *messageData
}

func NewTraceTelemetry(message string, severityLevel SeverityLevel) *TraceTelemetry {
	now := time.Now()
	data := &messageData{
		Message: message,
	}

	data.Ver = 2

	item := &TraceTelemetry{
		data: data,
	}

	item.timestamp = now
	item.context = NewItemTelemetryContext()

	return item
}

func (item *TraceTelemetry) Timestamp() time.Time {
	return item.timestamp
}

func (item *TraceTelemetry) Context() TelemetryContext {
	return item.context
}

func (item *TraceTelemetry) baseTypeName() string {
	return "Message"
}

func (item *TraceTelemetry) baseData() Domain {
	return item.data
}

type EventTelemetry struct {
	BaseTelemetry
	data *eventData
}

func NewEventTelemetry(name string) *EventTelemetry {
	now := time.Now()
	data := &eventData{
		Name: name,
	}

	data.Ver = 2

	item := &EventTelemetry{
		data: data,
	}

	item.timestamp = now
	item.context = NewItemTelemetryContext()

	return item
}

func (item *EventTelemetry) Timestamp() time.Time {
	return item.timestamp
}

func (item *EventTelemetry) Context() TelemetryContext {
	return item.context
}

func (item *EventTelemetry) baseTypeName() string {
	return "Event"
}

func (item *EventTelemetry) baseData() Domain {
	return item.data
}

type MetricTelemetry struct {
	BaseTelemetry
	data *metricData
}

func NewMetricTelemetry(name string, value float32) *MetricTelemetry {
	now := time.Now()
	metric := &DataPoint{
		Name:  name,
		Value: value,
		Count: 1,
	}

	data := &metricData{
		Metrics: make([]*DataPoint, 1),
	}

	data.Ver = 2
	data.Metrics[0] = metric

	item := &MetricTelemetry{
		data: data,
	}

	item.timestamp = now
	item.context = NewItemTelemetryContext()

	return item
}

func (item *MetricTelemetry) Timestamp() time.Time {
	return item.timestamp
}

func (item *MetricTelemetry) Context() TelemetryContext {
	return item.context
}

func (item *MetricTelemetry) baseTypeName() string {
	return "Metric"
}

func (item *MetricTelemetry) baseData() Domain {
	return item.data
}

type RequestTelemetry struct {
	BaseTelemetry
	data *requestData
}

func NewRequestTelemetry(name string, timestamp time.Time, duration time.Duration, responseCode string, success bool) *RequestTelemetry {
	now := time.Now()

	var (
		refHours        = int(duration.Hours())
		refMinutes      = int(duration.Minutes())
		refSeconds      = int(duration.Seconds())
		refMilliPlusOne = int(duration.Nanoseconds() / 1e3)
		days            = refHours / 24
		hours           = refHours - days*24
		minutes         = refMinutes - refHours*60
		seconds         = refSeconds - refMinutes*60
		milliPlusOne    = refMilliPlusOne - refSeconds*1e6
	)
	durationString := fmt.Sprintf("%02d.%02d:%02d:%02d.%04d",
		days,
		hours,
		minutes,
		seconds,
		milliPlusOne)

	data := &requestData{
		Id:           uuid.NewV4().String(),
		Name:         name,
		StartTime:    timestamp.Format(time.RFC3339Nano),
		Duration:     durationString,
		ResponseCode: responseCode,
		Success:      success,
	}

	data.Ver = 2

	item := &RequestTelemetry{
		data: data,
	}

	item.timestamp = now
	item.context = NewItemTelemetryContext()

	return item
}

func (item *RequestTelemetry) Timestamp() time.Time {
	return item.timestamp
}

func (item *RequestTelemetry) Context() TelemetryContext {
	return item.context
}

func (item *RequestTelemetry) baseTypeName() string {
	return "Request"
}

func (item *RequestTelemetry) baseData() Domain {
	return item.data
}
