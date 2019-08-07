# Prometheus server service

Executing prometheus server start command as service equivalent of `./prometheus --config.file=prometheus_cb.yml` get executed with this service config file.

### Instructions to Add Service in RedHat Systems

* Add this service in your VM with following command: `vi /etc/systemd/system/prometheus.service`
* Paste the specified 'prometheus.service' file's content in above opened file. 
* Change the `ExecStart` command base-path as per your need. 
  ```
  ExecStart=/base-path/prometheus-2.9.2.linux-amd64/prometheus \
    --config.file /base-path/prometheus-2.9.2.linux-amd64/prometheus_cb.yml
  ```
    
### Commands used to execute the service

After, making the changes and saving the `prometheus.service` file. Do execute `systemctl daemon reload` command which will reload the systemd manager configuration. 

* Command to start the service: `service prometheus start` or `systemctl start prometheus.service`.
* Command to know status of the service: `service prometheus status` or `systemctl status prometheus.service`.
* Command to restart the service: `service prometheus restart` or `systemctl restart prometheus.service`.
* Command to  stop the service: `service prometheus stop` or `systemctl stop prometheus.service`.


Similarly, services for `Node-Exporter` and `Alert-Manager` can be added. It's better to use service-based architecture to avoid sudden outages or failures during execution. As it will limit the capabilities of monitoring tool.

---

### Configuration File and Targets

Use `Node Exporter` for dynamically detecting changes in configuration file. We have opted for a design in `targets.json` will be scraped for finding out the targets that are needed to be monitored.

* `config.yml`: It contains information about scraping targets and rules that are needed to imported. Plus, the connection to `AlertManager` that will be monitoring the alerts on given targets.

* `targets.json`: It contains targets with proper labels that can be used to monitor. Here, we have added only label `cluster` this can be used directly with Grafana's json file also in order to get the cluster names that are stored in this label.




