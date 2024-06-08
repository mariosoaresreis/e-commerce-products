package config

type GCP struct {
	GoogleApplicationCredentials string   `json:"google_application_credentials"`
	ProjectID                    string   `json:"gcp_project_id"`
	ServicePubTopic              string   `json:"service_pub_topic"`
	SubscriptionTopics           []string `json:"subscription_topics"`
	IsLocalPubsub                string   `json:"gcp_is_local_pubsub"`
	PubsubHost                   string   `json:"gcp_pubsub_host"`
}
