package azuremonitor

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/grafana/grafana/pkg/models"
)

// AzureMonitorQuery is the query for all the services as they have similar queries
// with a url, a querystring and an alias field
type AzureMonitorQuery struct {
	URL           string
	UrlComponents map[string]string
	Target        string
	Params        url.Values
	RefID         string
	Alias         string
}

// AzureMonitorResponse is the json response from the Azure Monitor API
type AzureMonitorResponse struct {
	Cost     int    `json:"cost"`
	Timespan string `json:"timespan"`
	Interval string `json:"interval"`
	Value    []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Name struct {
			Value          string `json:"value"`
			LocalizedValue string `json:"localizedValue"`
		} `json:"name"`
		Unit       string `json:"unit"`
		Timeseries []struct {
			Metadatavalues []struct {
				Name struct {
					Value          string `json:"value"`
					LocalizedValue string `json:"localizedValue"`
				} `json:"name"`
				Value string `json:"value"`
			} `json:"metadatavalues"`
			Data []struct {
				TimeStamp time.Time `json:"timeStamp"`
				Average   float64   `json:"average,omitempty"`
				Total     float64   `json:"total,omitempty"`
				Count     float64   `json:"count,omitempty"`
				Maximum   float64   `json:"maximum,omitempty"`
				Minimum   float64   `json:"minimum,omitempty"`
			} `json:"data"`
		} `json:"timeseries"`
	} `json:"value"`
	Namespace      string `json:"namespace"`
	Resourceregion string `json:"resourceregion"`
}

// ApplicationInsightsResponse is the json response from the Application Insights API
type ApplicationInsightsResponse struct {
	Tables []struct {
		TableName string `json:"TableName"`
		Columns   []struct {
			ColumnName string `json:"ColumnName"`
			DataType   string `json:"DataType"`
			ColumnType string `json:"ColumnType"`
		} `json:"Columns"`
		Rows [][]interface{} `json:"Rows"`
	} `json:"Tables"`
}

// AzureLogAnalyticsResponse is the json response object from the Azure Log Analytics API.
type AzureLogAnalyticsResponse struct {
	Tables []struct {
		Name    string `json:"name"`
		Columns []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"columns"`
		Rows [][]interface{} `json:"rows"`
	} `json:"tables"`
}

// AzureMonitorQueryModel is the JSON data received from frontend
type AzureMonitorQueryModel struct {
	QueryMode string                      `json:"queryMode"`
	Data      map[string]AzureMonitorData `json:"data"`
	AzureMonitorData
}

// AzureMonitorData is a part of JSON containing query data
type AzureMonitorData struct {
	ResourceGroup       string   `json:"resourceGroup"`
	ResourceGroups      []string `json:"resourceGroups"`
	Locations           []string `json:"locations"`
	ResourceName        string   `json:"resourceName"`
	MetricDefinition    string   `json:"metricDefinition"`
	TimeGrain           string   `json:"timeGrain"`
	AllowedTimeGrainsMs []int64  `json:"allowedTimeGrainsMs"`
	MetricName          string   `json:"metricName"`
	MetricNamespace     string   `json:"metricNamespace"`
	Aggregation         string   `json:"aggregation"`
	Dimension           string   `json:"dimension"`
	DimensionFilter     string   `json:"dimensionFilter"`
	Alias               string   `json:"alias"`
	Format              string   `json:"format"`
}

// ResourcesResponse is the json response object from the Azure Monitor resources api.
type ResourcesResponse struct {
	Value []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Type     string `json:"type"`
		Kind     string `json:"kind,omitempty"`
		Location string `json:"location"`
	} `json:"value"`
}

// ResourcesLoader is the interface for Resources loading. Makes resource loading testable
type ResourcesLoader interface {
	Get(azureMonitorData *AzureMonitorData, subscriptions []interface{}, createRequest func(context.Context, *models.DataSource) (*http.Request, error)) ([]resource, error)
}
