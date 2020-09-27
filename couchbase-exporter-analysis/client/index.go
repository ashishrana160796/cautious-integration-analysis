package client
import "github.com/pkg/errors"

// Index returns the results of /pools/indexStatus
func (c Client) Index() (Index, error) {
              var index Index
              err := c.get("/indexStatus", &index)
              return index, errors.Wrap(err, "failed to get cluster")
}

// Index (/indexStatus)
type Index struct {
    Indexes struct {
        StorageMode string   `json:"storageMode"`
        Partitioned bool     `json:"partitioned"`
        InstID      uint64    `json:"instId"`
        Hosts       []string `json:"hosts"`
        Progress    int      `json:"progress"`
        Definition  string   `json:"definition"`
        Status      string   `json:"status"`
        Bucket      string   `json:"bucket"`
        Indx       string   `json:"index"`
        ID          uint64    `json:"id"`
    } `json:"indexes"`
    Version  int           `json:"version"`
    Warnings []interface{} `json:"warnings"`
}
