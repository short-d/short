package instrumentation

type SSO struct {
	event        string
	providerName string
}

func (o SSO) GetMetricName() string {
	return o.event
}

func NewSSO(providerName string) SSO {
	return SSO{
		event:        providerName + "-SSO",
		providerName: providerName,
	}
}
