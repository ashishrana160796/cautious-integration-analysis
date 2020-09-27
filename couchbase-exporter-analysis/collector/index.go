package collector

import (
    "sync"
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/common/log"
    "github.com/totvslabs/couchbase-exporter/client"
)

 

type indexCollector struct {
    mutex  sync.Mutex
    client client.Client

    up             *prometheus.Desc
    scrapeDuration *prometheus.Desc

    indexesStorageMode *prometheus.Desc
    indexesPartioned *prometheus.Desc
    indexesInstID *prometheus.Desc
    // indexesHosts *prometheus.Desc
    indexesProgress *prometheus.Desc
    // indexesDefinition *prometheus.Desc
    indexesStatus *prometheus.Desc
    // indexesBucket *prometheus.Desc
    // indexesIndx *prometheus.Desc
    indexesID *prometheus.Desc

    indexVersion *prometheus.Desc
   //  indexWarnings *prometheus.Desc

}


// NewIndexCollector index collector
func NewIndexCollector(client client.Client) prometheus.Collector {

    const subsystem = "index"
    // nolint: lll
    return &indexCollector{
        client: client,
        up: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "up"),
            "Couchbase cluster API is responding",
            nil,
            nil,
        ),
        scrapeDuration: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "scrape_duration_seconds"),
            "Scrape duration in seconds",
            nil,
            nil,
        ),
        indexesStorageMode: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_storage_mode"),
            "Mode of Index Storage",
            nil,
            nil,
        ),
        indexesPartioned: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_partioned"),
            "Partitioned Indexes",
            nil,
            nil,
        ),
        indexesInstID: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_inst_id"),
            "Inst Id of Index",
            nil,
            nil,
        ),
        indexesProgress: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_progress"),
            "Progress of Index",
            nil,
            nil,
        ),
        indexesStatus: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_status"),
            "Status of Index",
            nil,
            nil,
        ),
        indexesID: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_id"),
            "Id of Index",
            nil,
            nil,
        ),
        indexVersion: prometheus.NewDesc(
            prometheus.BuildFQName(namespace, subsystem, "indexes_version"),
            "Id of Index",
            nil,
            nil,
        ),
    }
}

 

// Describe all metrics
func (c *indexCollector) Describe(ch chan<- *prometheus.Desc) {
    ch <- c.up
    ch <- c.scrapeDuration

    ch <- c.indexesStorageMode
    ch <- c.indexesPartioned
    ch <- c.indexesInstID
    // ch <- c.indexesHosts
    ch <- c.indexesProgress
    // ch <- c.indexesDefinition
    ch <- c.indexesStatus
    // ch <- c.indexesBucket
    // ch <- c.indexesIndx
    ch <- c.indexesID
    ch <- c.indexVersion
    // ch <- c.indexWarnings
}

 

// Collect all metrics

// nolint: lll

func (c *indexCollector) Collect(ch chan<- prometheus.Metric) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    start := time.Now()
    log.Info("Collecting index metrics...")
    index, err := c.client.Index()

    if err != nil {
        ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)
        log.With("error", err).Error("failed to scrape cluster")
        return
    }

    ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
    ch <- prometheus.MustNewConstMetric(c.scrapeDuration, prometheus.GaugeValue, time.Since(start).Seconds())

    ch <- prometheus.MustNewConstMetric(c.indexesStorageMode, prometheus.GaugeValue, fromBool(index.Indexes.StorageMode == "plasma"))
    ch <- prometheus.MustNewConstMetric(c.indexesPartioned, prometheus.GaugeValue, fromBool(index.Indexes.Partitioned))
    ch <- prometheus.MustNewConstMetric(c.indexesInstID, prometheus.GaugeValue, float64(index.Indexes.InstID))
    // ch <- prometheus.MustNewConstMetric(c.indexesHosts, prometheus.GaugeValue, index.Indexes.Hosts)
    ch <- prometheus.MustNewConstMetric(c.indexesProgress, prometheus.GaugeValue, float64(index.Indexes.Progress))
    // ch <- prometheus.MustNewConstMetric(c.indexesDefinition, prometheus.GaugeValue, index.Indexes.Definition)

    ch <- prometheus.MustNewConstMetric(c.indexesStatus, prometheus.CounterValue, fromBool(index.Indexes.Status == "Ready"))
    // ch <- prometheus.MustNewConstMetric(c.indexesBucket, prometheus.CounterValue, index.Indexes.Bucket)
    // ch <- prometheus.MustNewConstMetric(c.indexesIndx, prometheus.CounterValue, index.Indexes.Indx)
    ch <- prometheus.MustNewConstMetric(c.indexesID, prometheus.CounterValue, float64(index.Indexes.ID))
    ch <- prometheus.MustNewConstMetric(c.indexVersion, prometheus.CounterValue, float64(index.Version))
    // ch <- prometheus.MustNewConstMetric(c.indexWarnings, prometheus.CounterValue, index.Warnings)
}
