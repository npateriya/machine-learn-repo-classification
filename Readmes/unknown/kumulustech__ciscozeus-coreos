# ciscozeus-coreos Kubernetes with Docker engine support

This project provides a container (Dockerfile or on docker hub: kumulustec/td-agent-cz),
with a modified fluent-cadvisor-plugin script.  The script is already loaded into
the container (along with other normally deployed components for log aggregation), and
a run.sh script is set as the entrypoint so that the agent configuration can be
fine tuned for certain parameters.

Specifically, due to changes to how the Kubelet embedded cadvisor runs, it is necessary
to provide the local cadvisor hostname or ip to the container on start (via an
environment variable in the pod .yaml file), and we will start the agent container
via the kubelet mechanism, by placing the .yaml description in
/etc/kubernetes/manifests on each of the managed coreos machines.

In addition, the container needs the CiscoZeus user and token defined, and to that
end, there is a script 'launch_cz_agent.sh' which expects, in order:

    ./launch_cz_agent.sh cz-username cz-token core-os-target-ip

The script will update an agent.yaml pod description, and try to copy it to the
core-os-target-ip machine as the core user, and then move that file to the
/etc/kubernetes/manifests directory to launch the agent service.

Once this is done, the kubernetes UI or kubectl --namespace=kube-system get pods
should show a ciscozeus-logforwarder-TARGET_IP_ADDR_OR_NAME pod as up and running
and your CiscoZeus entity should start to receive data from both the system containers
and user containers launched in the Kubernetes environment.

## Build your own agent containers
There is a Dockerfile provided in case one would prefer to build their own container,
or wants to further modify how the service is started.  Note that the run.sh script
that is set as the default endpoint for the container, also modifies the td-agent.conf
script in order to define the cadvisor endpoint, and to incorporate the ciscozeus
user and token information.

In addition to the run script, there is a modified version of the
fluent-cadvisor-plugin script that works with the current CoreOS/Kubernetes/Docker
environment.  It is current configured to leverage the cAdvisor 1.3 API, which
provides a mechanism for discovering and monitoring Docker based container
deployments.  Note that the logic currently queries docker via the socket
interface in order to discover the actual container IDs and then querries the
cAdvisor API to determine information about the related PODs.  The script attemptes
to query all containers in a pod based on the docker discovered container mappings.

This script is not currently bundled as a gem, and instead is simply loaded as a
script during the container build process.

Files expected by the current build process:

Dockerfile : build description for 'docker build -t org/name:latest .'
in_cadvisor.rb : fluent-cadvisor-plugin script to be placed in /etc/td-agent/plugins
  directory in the contianer (handled by Dockerfile)
td-agent.repo : YUM repository file to support installing td-agent on a CentOS
  based container
run.sh : Container entrypoint script that loads the service description and
  does environment variable substitution on contaienr launch

td-agent.conf : an exmpale configuration file (note that this file doesn't get used
  by the current build process, it is only an example)
agent.yml : YAML POD description for the ciscozeus-logforwarder PODs, not installed
  in the container.


## CiscoZeus data discovery

Data can be discovered in the normal way in CiscoZeus, with tag searches for logs.cadvisor* and logs.kubernetes* providing most of the data that is currently
being captured.   From there, the normal visualizations and other analyses can be
accomplished.

## in_cadvisor.rb
Copyright (c) 2014 Alex

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

The in_cadvisor.rb code comes from the fluent-cadvisor-plugin project:
https://github.com/Woorank/fluent-plugin-cadvisor

It is licensed under the MIT license as described above.

## the rest of the work is released under the Apache, Version 2.0 License
   Copyright 2016 Kumulus Technologies
   Copyright 2016 Chad Merritt
   Copyright 2016 Robert Starmer

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
