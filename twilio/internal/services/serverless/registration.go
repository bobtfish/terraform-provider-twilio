package serverless

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "Serverless"
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"twilio_serverless_environment": resourceServerlessEnvironment(),
		"twilio_serverless_service":     resourceServerlessService(),
		"twilio_serverless_variable":    resourceServerlessVariable(),
	}
}
