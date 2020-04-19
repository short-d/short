package instrumentation

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/keygen"
)

// Factory initializes instrumentation code.
type Factory struct {
	serverEnv   fw.ServerEnv
	keyGen      keygen.KeyGenerator
	logger      fw.Logger
	tracer      fw.Tracer
	timer       fw.Timer
	metrics     fw.Metrics
	analytics   fw.Analytics
	geoLocation fw.GeoLocation
	network     fw.Network
}

// NewHTTPRequest creates and initializes Instrumentation tied to the given HTTP
// request.
func (f Factory) NewHTTPRequest(req *http.Request) Instrumentation {
	requestID, err := f.keyGen.NewKey()
	if err != nil {
		f.logger.Error(err)
	}

	connection := f.network.FromHTTP(req)
	clientIP := connection.ClientIP

	location, err := f.geoLocation.GetLocation(clientIP)
	if err != nil {
		f.logger.Error(err)
	}
	f.logger.Info(fmt.Sprintf("%v", location))

	ctx := fw.ExecutionContext{
		RequestID:      string(requestID),
		RequestStartAt: f.timer.Now(),
		Location:       location,
	}

	return Instrumentation{
		logger:      f.logger,
		tracer:      f.tracer,
		timer:       f.timer,
		metrics:     f.metrics,
		analytics:   f.analytics,
		geoLocation: f.geoLocation,
		ctx:         ctx,
	}
}

// NewFactory creates Instrumentation factory.
func NewFactory(
	serverEnv fw.ServerEnv,
	logger fw.Logger,
	tracer fw.Tracer,
	timer fw.Timer,
	metrics fw.Metrics,
	analytics fw.Analytics,
	geoLocation fw.GeoLocation,
	keyGen keygen.KeyGenerator,
	network fw.Network,
) Factory {
	return Factory{
		serverEnv:   serverEnv,
		logger:      logger,
		tracer:      tracer,
		timer:       timer,
		metrics:     metrics,
		analytics:   analytics,
		geoLocation: geoLocation,
		keyGen:      keyGen,
	}
}
