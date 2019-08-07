# Dashboard creation process in Grafana with JSON files

This grafana dashboard json file gets created when `dashboard.jsonnet` file gets built while building couchbase-exporter. This building from _jsonnet_ does require dependencies from jsonnet and grafonnet libraries. Instead, here we manipulate this json directly, the manipulation part is intuitive with similar elements being pasted and edited over to produce new dashboard elements. 

Advantage of using direct manipulation is that we skip extra build part with grafonnet like libraries with jsonnet backend. But, it demands more careful work and doesn't leverage capabilities of jsonnet language.

## Description

Both `instance-dashboard.json` & `cluster-dashboard.json` are our modified dashboards containing some extra elements for observing newer metrics that might be of interest. Also, we have `base-dashboard.json` which is OOB grafana JSON file for observation and getting started from original [couchbase-blog](https://github.com/couchbaselabs/blog-source-code/blob/master/Meliani/GrafanaDashboard.json)

* `instance-dashboard.json`: Sticks to standard `instance` variable based naming convention like _localhost:9420_ etc. like options for listening stats over an instance. It is there to get started with simple targets from Prometheus end with no modifications to its configuration file.

* `cluster-dashboard.json`: It utilizes the `targets.json` file's from Prometheus configuration file `cluster` label to select over couchbase VMs over which cb-exporter is listening. Hence, the name of cluster that is fetched from ReST endpoint `pools/default` matters a lot otherwise the VMs will be pushed onto orphan cluster tag.

Hence, in this repository we provide options to use both if any needed. Also, you can do simple `diff` comparison b/w the two also to understand the manipulating techniques of these large json files.

## Instructions to Use

* Open Grafana -> Click on `+` symbol -> Click on `import` button -> Then, paste the JSON below on `Or paste JSON`  label or upload .json file.
  - Your dashboard is ready to use.


## Reading through this JSON file

* All the new changes with respect to new variables like `node`, `bucket` or other dashboard graph elements added can be observed with 6 extra newlines below & above these elements.  

  - Read along and observe the `query` part of these elements for editing them. These PromQL queries contain all the neccessary patterns to learn about writing newer dashboard elements.
  
  - While modifying and adding newer metrics make sure that the newer added metric is having a newer `id` in case of duplicate ids issue with respect to dashboarding will come.
