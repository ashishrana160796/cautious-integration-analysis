# Create functions as utility for ReST HTTP server.
import subprocess, random
import requests
import json
from requests.auth import HTTPBasicAuth

# finding cluster name & feeding into the targets.json file for labels of prometheus
def clust_name(hostname,username,password):
    port='8091'
    URL='http://'+hostname+':'+port

    try:
        resp = requests.get(URL + '/pools/default', auth=HTTPBasicAuth(username, password))
        if resp.status_code != 200:
            return 0
        else:
            data=resp.json()
            cluster_name=data['clusterName']
            if len(cluster_name.replace(" ","")) == 0:
                return "Orphan VM. No, Cluster Name."
            else:
                return cluster_name
    except requests.exceptions.RequestException as err:
        print (hostname + " - OOps: Something Else",err)
    except requests.exceptions.HTTPError as errh:
        print (hostname + " - Http Error:",errh)
    except requests.exceptions.ConnectionError as errc:
        print (hostname + " - Error Connecting:",errc)
    except requests.exceptions.Timeout as errt:
        print (hostname + " - Timeout Error:",errt)

# Each VM for exporting stats by couchbase-exporter requires it's own specific port number.
def get_empty_port():
    # Assumption empty is available in the given VM.
    port_no = -1
    mssg_dgst = '-1 Imaginary Port'
    while( len(mssg_dgst) != 0 ):
        port_no = random.randint(16384,65535)
        eff_command = "netstat -anp | grep ':"+str(port_no)+"'"
        # giving shell=True is a bad code design, for insecure input: hazardous pointed by documentation
        p = subprocess.Popen(eff_command, stdout=subprocess.PIPE, shell=True)
        mssg_dgst = p.communicate()[0]
    return port_no

import json

# write VM related information into the reference file.
def write_util(hostname,username, password, port_no):
    with open("targets.json", "r") as read_file:
        data=json.load(read_file)
    str_val = username+","+password+","+str(port_no)
    data[hostname] = str_val
    print("Inserted Data: for key"+" hostname "+data[hostname])
    with open("targets.json", "w+") as write_file:
        data=json.dump(data,write_file, indent=4, separators=(',', ': '))

# return information related to particular host.
def get_util(hostname):
    with open("targets.json", "r") as read_file:
        data=json.load(read_file)
    if hostname in data:
        return data[hostname]
    return "Key Not Present. Please Check."

# return all information of reference file.
def get_view():
    with open("targets.json", "r") as read_file:
        data=json.load(read_file)
    return data

# Delete given Key entry from reference file.
def del_util(hostname):
    with open("targets.json", "r") as read_file:
        data=json.load(read_file)
    str_val = ""
    port_no = "-1"
    if hostname in data:
        str_val = data.pop(hostname)
        val_lst = str_val.split(',')
        port_no = val_lst[2]
        str_val = "Values Deleted: "+str_val+" for Key: "+hostname
    else:
        str_val = "Key Not Present. Please Check."
    with open("targets.json", "w+") as write_file:
        json.dump(data,write_file, indent=4, separators=(',', ': '))
    return str_val, port_no

# Start the couchbase exporter with executing the command.
def cbexport_start(hostname):
    with open("targets.json", "r") as read_file:
        data=json.load(read_file)
    if hostname in data:
        str_vals=data[hostname]
        lst_vals=str_vals.split(",")
        eff_command="nohup ./../CBExporter/couchbase-exporter --couchbase.username "+lst_vals[0]+" --couchbase.password "+lst_vals[1]+" --web.listen-address="+ \
                '":'+lst_vals[2]+'" --couchbase.url="http://'+hostname+':8091" &'
    else:
        eff_command="echo Malformed Command Exception: Key Not Found."
    p = subprocess.Popen(eff_command, stdout=subprocess.PIPE, shell=True)
    mssg_dgst = p.communicate()[0]
    return eff_command, mssg_dgst



import re

# Delete or end the couchbase exporter process.
def cbexport_del(port_no):
    # Need to be 'root' to know the process id information.
    port_command = 'netstat -nlp | grep ":'+str(port_no)+'"'
    p = subprocess.Popen(port_command, stdout=subprocess.PIPE, shell=True)
    mssg_dgst = p.communicate()[0]
    if len(mssg_dgst) != 0:
        mssg_lst = mssg_dgst.split(" ")
        mssg_lst[:] = [item for item in mssg_lst if item != '']
        rgx = re.compile(r'[0-9]+\/\w+')
        pid_list = list(filter(rgx.match, mssg_lst))
        pid_list = map(lambda x: x.split("/")[0], pid_list)
        if len(pid_list) != 0:
            for x in pid_list:
                kill_command = 'kill -SIGTERM '+str(x)
                p = subprocess.Popen(kill_command, stdout=subprocess.PIPE, shell=True)
                mssg_dgst = p.communicate()[0]
                print(mssg_dgst)
    else:
        print("Port number not in use.")

# Write into json file of prometheus running server. For dynamic detection by Node Exporter.
def write_targets(port_no,hostname,username,password):
    targ_str = "localhost:"+str(port_no)
    # These two entries can be fetched from global repository
    # get this data with another rest query to server
    clus_str = clust_name(hostname, username, password)
    # eff_net_str for inputting into targets.json
    eff_net_str = '{ "targets": [ "localhost:'+str(port_no)+'" ], "labels": { "cluster": "'+clus_str+'"} }'
    data_appnd = json.loads(eff_net_str)
    with open('../prometheus-2.9.2.linux-amd64/targets.json','r') as read_file:
        data = json.load(read_file)
    data.append(data_appnd)
    with open('../prometheus-2.9.2.linux-amd64/targets.json','w+') as write_file:
        data=json.dump(data,write_file, indent=4, separators=(',', ': '))

# Delete value from json file of prometheus running server. For dynamic detection by Node Exporter.
def del_targets(port_no):
    with open('../prometheus-2.9.2.linux-amd64/targets.json','r') as read_file:
        data = json.load(read_file)
    i=0
    for x in data:
        str_ports = ''.join(x['targets'])
        str_lst = re.findall(r'\d+', str_ports)
        str_ports = ''.join(str_lst)
        # print(str_ports)
        if str_ports == str(port_no):
            # print(re.findall(r'\d+', str(x['targets'][0])))
            # print(str(port_no))
            s=data.pop(i)
            # print(s)
        i=i+1
    # print(data)
    with open('../prometheus-2.9.2.linux-amd64/targets.json','w+') as write_file:
        data=json.dump(data,write_file, indent=4, separators=(',', ': '))
