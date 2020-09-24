## Dataset preparation methodology, description and assumptions taken

This individual file in this dataset is prepared from aggregating multiple files in the `realAWSCloudwatch/` folder from NAB benchmark open source project. The univariate dataset files are combined in a manner where feature values with closest dates or timestamp values are combined together to create new multivariate dataset files representing the hypothetical cluster topologies. The purpose to create these cluster topologies is to generate multivariate dataset where different clusters don't interact and run independently. These cluster topologies don't exist originally but are safely assumed as metrics of different machines from a cluster are not considered to be dependent. These can be thought of as mirror database instances running independently in AWS and are all active for high availability. This methodology has resulted in some interesting topologies for us to analyze and work on.

The technique that we have used to combine multiple files in AWS Cloudwatch data exactly involves first scanning all the files and identifying files that captures a stream of data under similar date or timestamp windows doesn't have to be an exact match. The assumption is the data monitored across same or similar date windows belong to that given cluster only and the reasoning for the given assumption is provided above. Hence, all the EC2 machines combine to form that cluster and accordingly different features of that cluster are captured like network stats, CPU utilization, request count etc. We believe it to be a fairly reasonable assumption to create our multivariate dataset as all machines are assumed to have interacted independently during monitoring period and we are not interested in studying correlated variations amongst variables as well.

``` 
-- Logic for time frame approximate matching used to merge different univariate datasets.
-- In SQL

SELECT UNITED_DATA.*
(
    SELECT one.*, two.*
    , RANK() OVER ( ORDER BY ABS(one.timestamp - two.timestamp) DESC ) as rank_var
    FROM table_one one
    LEFT JOIN table_two two
) UNITED_DATA
WHERE UNITED_DATA.rank_var = 1
```

---

## Assumed Topology Description

* __multi-var-five.csv:__ It is a four VM EC2 cluster with one of the VM having a redshift instance loaded as a database onto it while for others we are not aware what's keeping them occupied.
  * `realAWSCloudwatch/`  _CSV Files Used:_ ec2_cpu_utilization_24ae8d, ec2_cpu_utilization_53ea38, rds_cpu_utilization_cc0c53, ec2_cpu_utilization_5f5533, ec2_cpu_utilization_fe7f93
* __multi-var-four-one.csv:__ It is a three VM EC2 cluster with one of them having their disk utilization description also being available to us.
  * `realAWSCloudwatch/`  _CSV Files Used:_ ec2_cpu_utilization_c6585a, ec2_cpu_utilization_ac20cd, ec2_disk_write_bytes_c0d644, ec2_cpu_utilization_77c1ca
* __multi-var-four-two.csv__ It is the most interesting dataset out of all with one single VM of EC2 having its memory utilization, network utilization, request count number and redshift metric data available to us.
  * `realAWSCloudwatch/`  _CSV Files Used:_ ec2_network_in_257a54, ec2_cpu_utilization_825cc2, elb_request_count_8c0756, rds_cpu_utilization_e47b3b

