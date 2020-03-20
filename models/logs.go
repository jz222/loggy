package models

type Error struct {
	Message     string            `json:"message" bson:"message"`
	Stacktrace  string            `json:"stacktrace" bson:"stacktrace"`
	Path        string            `json:"path" bson:"path"`
	Line        string            `json:"line" bson:"line"`
	Type        string            `json:"type" bson:"type"`
	Fingerprint string            `json:"fingerprint" bson:"fingerprint"`
	Badges      map[string]string `json:"badges" bson:"badges"`
	Snippet     map[string]string `json:"snippet" bson:"snippet"`
	Logs        []string          `json:"logs" bson:"logs"`
	Ticket      string            `json:"ticket" bson:"ticket"`
	Host        string            `json:"host" bson:"host"`
	UserAgent   string            `json:"userAgent" bson:"userAgent"`
	ClientIP    string            `json:"clientIp" bson:"clientIp"`
	Count       int               `json:"count,omitempty" bson:"count,omitempty"`
	Timestamp   string            `json:"timestamp" bson:"timestamp"`
}
