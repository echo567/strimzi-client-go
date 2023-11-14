// Code generated by schema-generate. DO NOT EDIT.

package v1beta2

// AutoRestart The auto restart status.
type AutoRestart struct {

  // The name of the connector being restarted.
  ConnectorName string `json:"connectorName,omitempty"`

  // The number of times the connector or task is restarted.
  Count int `json:"count,omitempty"`

  // The last time the automatic restart was attempted. The required format is 'yyyy-MM-ddTHH:mm:ssZ' in the UTC time zone.
  LastRestartTimestamp string `json:"lastRestartTimestamp,omitempty"`
}

// ConditionsItems 
type ConditionsItems struct {

  // Last time the condition of a type changed from one status to another. The required format is 'yyyy-MM-ddTHH:mm:ssZ', in the UTC time zone.
  LastTransitionTime string `json:"lastTransitionTime,omitempty"`

  // Human-readable message indicating details about the condition's last transition.
  Message string `json:"message,omitempty"`

  // The reason for the condition's last transition (a single word in CamelCase).
  Reason string `json:"reason,omitempty"`

  // The status of the condition, either True, False or Unknown.
  Status string `json:"status,omitempty"`

  // The unique identifier of a condition, used to distinguish between other conditions in the resource.
  Type string `json:"type,omitempty"`
}

// Config The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.
type Config struct {
}

// ConnectorStatus The connector status, as reported by the Kafka Connect REST API.
type ConnectorStatus struct {
}

// KafkaConnector 
type KafkaConnector struct {

  // The specification of the Kafka Connector.
  Spec *Spec `json:"spec,omitempty"`

  // The status of the Kafka Connector.
  Status *Status `json:"status,omitempty"`
}

// Spec The specification of the Kafka Connector.
type Spec struct {

  // Automatic restart of connector and tasks configuration.
  AutoRestart *AutoRestart `json:"autoRestart,omitempty"`

  // The Class for the Kafka Connector.
  Class string `json:"class,omitempty"`

  // The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.
  Config *Config `json:"config,omitempty"`

  // Whether the connector should be paused. Defaults to false.
  Pause bool `json:"pause,omitempty"`

  // The maximum number of tasks for the Kafka Connector.
  TasksMax int `json:"tasksMax,omitempty"`
}

// Status The status of the Kafka Connector.
type Status struct {

  // The auto restart status.
  AutoRestart *AutoRestart `json:"autoRestart,omitempty"`

  // List of status conditions.
  Conditions []*ConditionsItems `json:"conditions,omitempty"`

  // The connector status, as reported by the Kafka Connect REST API.
  ConnectorStatus *ConnectorStatus `json:"connectorStatus,omitempty"`

  // The generation of the CRD that was last reconciled by the operator.
  ObservedGeneration int `json:"observedGeneration,omitempty"`

  // The maximum number of tasks for the Kafka Connector.
  TasksMax int `json:"tasksMax,omitempty"`

  // The list of topics used by the Kafka Connector.
  Topics []string `json:"topics,omitempty"`
}
