# Deploying Multi-OS applications to Docker EE
Docker EE 17.06 is the first Containers-as-a-Service platform to offer production-level support for the integrated management and security of Linux AND Windows Server Containers.

In this lab we'll build a Docker EE cluster comprised of Windows and Linux nodes. Then we'll deploy both a Linux and Windows web app, as well as a multi-service application that includes both Windows and Linux components.

> **Difficulty**: Intermediate (assumes basic familiarity with Docker)

> **Time**: Approximately 60 minutes

> **Tasks:**
>
> * [Prerequisites](#prerequisites)
> * [Task 1: Build a Docker EE Cluster](#task1)
>   * [Task 1.1: Install the UCP manager](#task1.1)
>   * [Task 1.2: Install a Linux worker node](#task1.2)
>   * [Task 1.3: Install a Windows worker node](#task1.3)
>   * [Task 1.4: Install DTR and Create Two Repositories](#task1.4)
>   * [Task 1.5: Install Self Signed Certs on All Nodes](#task1.5)
> * [Task 2: Deploy a Linux Web App](#task2)
>   * [Task 2.1: Clone the Demo Repo](#task2.1)
>   * [Task 2.2: Build and Push the Linux Web App Image](#task2.2)
>   * [Task 2.3: Deploy the Web App using UCP](#task2.3)
> * [Task 3: Deploy a Windows Web App](#task3)
>   * [Task 3.1: Create the Dockerfile with Image2Docker](#task3.1)
>   * [Task 3.2: Build and Push Your Image to Docker Trusted Registry](#task3.2)
>   * [Task 3.3: Deploy the Windows Web App](#task3.3)

## Document Conventions

- When you encounter a phrase between `<` and `>` you are meant to substitue a different value.
	For example: if you see `<linux vm name>` you would actually type something like `pod-1-lin01`

- When you see the Linux Penguin, all of the following instructions should be executed in one of your Linux VMs.

	![](./images/linux75.png)
	
- When you see the Windows Flag, all of the subsequent instructions should be completed in your Windows VM.

	![](./images/windows75.png)
	
### Virtual Machine Naming Conventions
Your VMs are named in the following convention pod-X-[OS]YY.csc.pittsburgh.cisco.com
Where X is the pod number [OS] is the Operating System type and YY is the Node number.  

> For example: `pod-3-lin02` or `pod-3-win01`

### Virtual Machine Roles
This lab uses a total of five virtual machines

The Docker EE cluster you will be building will be comprised of five nodes - 4 Linux nodes running Centos7 and 1 Windows node running Server 2016.

![](./images/vm_roles.png)

## <a name="prerequisites"></a>Prerequisites

You will be provided a set of five virtual machines (four Linux and 1 Windows), which we will be configuring with Docker	. You do not need Docker running on your laptop, but you will need a Remote Desktop client to connect to the Windows VM, and an SSH client to connect into the Linux one.
### 1. RDP Client

- Windows - use the built-in Remote Desktop Connection app.
- Mac - install [Microsoft Remote Desktop](https://itunes.apple.com/us/app/microsoft-remote-desktop/id715768417?mt=12) from the app store.
- Linux - install [Remmina](http://www.remmina.org/wp/), or any RDP client you prefer.

### 2. SSH Client

- Windows - [Download Putty](http://www.chiark.greenend.org.uk/~sgtatham/putty/download.html)
- Linux - Use the built in SSH client
- Mac - Use the built in SSH client

> **Note**: When you connect to the Windows VM, if you are prompted to run Windows Update, you should cancel out. The labs have been tested with the existing VM state and any changes may cause problems.

## <a name="task1"></a>Task 1: Build a Docker EE Cluster

In this first step we're going to install Docker Universal Control Plane (UCP) and Docker Trusted Registry. UCP is a web-based control plane for Docker containers that can deploy and manage Docker-based applications across Windows and Linux nodes.  Docker Trusted Registry is a private registry server for storing your Docker images.

### <a name="task1.1"></a>Task 1.1: Install the UCP manager

> **Note**: In the current version of UCP manager nodes must be Linux. Worker nodes can be Windows or Linux

![](./images/linux75.png)

1. Either in a terminal window (Mac or Linux) or using Putty (Windows) SSH into Linux node **01** using the IP Address. The IP Address can be found on your lab information paper.  The username is `docker` and the password is `Docker2017`

	`ssh docker@<pod-X-lin01>`

Start the UCP installation on the current node, and make it your UCP manager.

Run the following command to pull the image and install UCP.

```
docker container run --rm -it --name ucp \
-v /var/run/docker.sock:/var/run/docker.sock \
docker/ucp:2.2.2 install \
--admin-username docker \
--admin-password Docker2017 \
--host-address <pod-X-lin01 Ip Address> \
--interactive
```

You will see output similar to the following:

```
Unable to find image 'docker/ucp:latest' locally
latest: Pulling from docker/ucp
88286f41530e: Pull complete 
054affee0b99: Pull complete 
1247b4599b73: Pull complete 
Digest: sha256:fb8935c92876a5255bd5ab26f2a9910918490b7eb8457323d46e416ac2b4d226
Status: Downloaded newer image for docker/ucp:latest
INFO[0000] Verifying your system is compatible with UCP 2.2.2 (94d1abdf3) 
INFO[0000] Your engine version 17.06.1-ee-2, build 8e43158 (3.10.0-514.26.2.el7.x86_64) is compatible 
WARN[0000] Your system uses devicemapper.  We can not accurately detect available storage space.  Please make sure you have at least 3.00 GB available in /var/lib/docker
```

This will take a few minutes while the images are downloaded and deployed.
When prompted for something similar to the below press enter.

```
You may enter additional aliases (SANs) now or press enter to proceed with the above list.
Additional aliases:
```

1. Download the [license file](https://drive.google.com/file/d/0ByQd4O58ibOEazFVRUJYNDYxMjg/view?usp=sharing) to your local laptop.

2. Open a web browser and type in the ip address of `pod-X-lin01` in the address bar.  Ignore any certificate warnings and continue.

	![](./images/ucp_login.png)

3. Login to UCP using the credentials we defined during the installation (docker/Docker2017).

	![](./images/ucp_license.png)
	
4. Click `Upload License`, navigate to your license location, and double-click the license file

Congratulations, you have installed the UCP manager node.

### <a name="task1.2"></a>Task 1.2: Install a Linux worker nodes

Now that we have a manager node, we'll add the Linux worker nodes to our cluster. Worker nodes are the servers that actually run our Docker-based applications.

1. From the main UCP dashboard click `Nodes` and then `Add Nodes`.
2. Copy the text in the box to your clipboard.
	> **Note** There is an icon in the upper right corner of the box that you can click to copy the text to your clipboard
	![](./images/add_node.png)
3.  SSH into the remaining Linux nodes (pod-X-lin02 - pod-X-lin04)
	`ssh docker@<linux node IP Address>`
4. Paste the text from Step 2 at the command prompt, and press enter.

	You should see the message `This node joined a swarm as a worker.` indicating you've successfully joined the node to the cluster.

5. Switch back to the UCP console in your web browser and click the `x` in the upper right corner to close the `Add Node` window

6. You should be taken to the `Nodes` screen will will see 4 nodes listed at the bottom of your screen. Your  lin01 node is the manager, and the lin02-04 node are your workers.

Congratulations on adding your worker nodes.

In the next step we'll install and configure a Windows worker node.

### <a name="task1.3"></a>Task 1.3: Install a Windows worker node

![](./images/windows75.png)

Let's add our 5th node to the cluster, a Windows Server 2016 worker node. The process is basically exactly the same as it was for Linux

1. From the Nodes screen, click the blue `Add node` button in the middle of the screen on the right hand side.

> **Note**: You may notice that there is a UI component to select `Linux` or `Windows`. The lab VMs already have the Windows components pre installed, so you do NOT need to select `Windows`. Just leave the selecton on `Linux` and move on to step 2

2. Copy the text from the dark box shown on the `Add Node` screen.

	> **Note** There is an icon in the upper right corner of the box that you can click to copy the text to your clipboard
	
	![](./images/add_node.png)

3. Use RDP to log in to `pod-X-win05`.

4. From the Start menu open a Powershell window (Start > Powershell (right click and run as administrator)

4. Paste the text from Step 2 at the command prompt, and press enter.

	You should see the message `This node joined a swarm as a worker.` indicating you've successfully joined the node to the cluster.

5. Switch back to your web browser and click the `x` in the upper right corner to close the `Add Node` window

6. You should be taken to the `Nodes` screen will will see 5 nodes listed at the bottom of your screen. Your Linux nodes and Windows nodes.

Congratulations on building your UCP cluster. Next up we'll install and configure DTR.

### <a name="task1.4"></a>Task 1.4: Install DTR and Create Two Repositories

![](./images/linux75.png)

Like UCP, DTR uses a single Docker container to bootstrap the install process. In the first step we'll kick off that container to install DTR, and then we'll create two repositories that we'll use later for our Tweet apps we're going to deploy.

1.  Run the bootstrap to deploy DTR.
	You will need to supply three inputs:
	
	* **--dtr-external URL**: `<pod-X-lin01 IP Address>`
	* **--ucp-node**: The hostname of `lin02` (For example: pod-X-lin02)
	* **--ucp-url**: The IP address of the UCP Manager (pod-X-lin01)
	
	```
	docker run -it --rm docker/dtr install \
	--dtr-external-url <pod-X-lin02 IP Address \
	--ucp-node <pod-X-lin02> \
	--ucp-username docker \
	--ucp-password Docker2017 \
	--ucp-url <pod-X-lin01 IP Address> \
	--ucp-insecure-tls
	```

	You will see a lot of output scroll by while DTR installs finishing with:

	```
	<output deleted>
	INFO[0160] Successfully registered dtr with UCP
	INFO[0161] Establishing connection with Rethinkdb
	INFO[0162] Background tag migration started
	INFO[0162] Installation is complete
	INFO[0162] Replica ID is set to: 8ec3809e352e
	INFO[0162] You can use flag '--existing-replica-id 8ec3809e352e' when joining other replicas to your Docker Trusted Registry Cluster
	INFO[0185] finished reading output
	```
5. Point your web browser to `https://<pod-X-lin02 IP Address>` and log in with the username `docker` and the password `Docker2017`.

	> **Note**: You need to use `https:` NOT `http`

	> **Note**: Because UCP uses self-signed SSL certs, your web browser may warn you that your connection is not secure. You will need to click through that warning. An example from Chrome is shown below.

Now that DTR is installed, let's go ahead and create a couple of repositories to hold our tweet app images. Repositories are how images are organized within a DTR server. Each image gets pushed to its own repository. Multiple images can be stored within a repository by supplying a different tag to each version.

2. Click the green `New repository` button on the right hand side of the screen. This brings up the new repository dialog

3. Under `REPOSITORY NAME` type `linux_tweet_app` and click `Save`

Let's repeat this process to create a repository for our Windows tweet app.

4. Click the green `New repository` button on the right hand side of the screen. This brings up the new repository dialogue.

5. Under `REPOSITORY NAME` type `windows_tweet_app` and click `Save`

Congratulations you've installed Docker Trusted Registry, and have created two new repositories.

### <a name="task1.5"></a>Task 1.5: Install Self Signed Certs on All Nodes

Docker uses TLS to ensure the identity of the Docker Trusted Registry. In a production environment you would use certs that come from a trusted certificate authority (CA). However, by default when you install UCP and DTR they use self-signed certs. These self-signed certs are not automatically trusted by the Docker engine. In order for them to be trusted, we need to copy down the root CA cert from the DTR server onto each node in the cluster. There is a script on each of your nodes that will do this for you

> **Note**: This step is only necessary in POC environments where trusted 3rd party certs are not used

Perform the following steps on all 4 of your Linux nodes (pod-X-lin01, pod-X-lin02, pod-X-lin03, pod-X-lin04)

![](./images/linux75.png)

1. SSH into the Linux node

2. At the command prompt run the `copy_certs` script passing it the ip address of your linux lin02 node.

	`$ ./copy_certs.sh <pod-X-lin02 Ip address>`

	You should see the following output


		% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
		100  1988  100  1988    0     0   7216      0 --:--:-- --:--:-- --:--:--  7229
		Updating certificates in /etc/ssl/certs...
		1 added, 0 removed; done.
		Running hooks in /etc/ca-certificates/update.d...
		done.

> **Note**: In some cases you may see some Perl warnings in addition to the above output, these can be safely ignored.

3. Log into the DTR server from the command line to ensure the cert was copied correctly. The username should be `docker` and the password `Docker2017`

	> **Note**: Be sure to substitute the FQDN of your **B** Linux node

	> **Note**: If you see an x509 certificate is from an unknown source error the cert didn't copy correctly, just rerun the above command.

	```
	$ docker login <pod-X-lin02 IP Address>
	Username: docker
	Password: Docker2017
	```

	You should see a `Login succeeded` message

> **Note**: Remember to repeat these steps on all 4 Linux nodes.

![](./images/windows75.png)

Now we need to do something similar on the two Windows nodes. Perform the following steps on each of the two Windows nodes.

1. RDP into the Windows node

2. From the start menu, open a Powershell window

3. Execute the `copy_certs` script

	`c:\copy_certs.ps1 <pod-X-lin02 Ip Address>`

3. Log into the DTR server from the command line to ensure the cert was copied correctly. The username should be `docker` and the password `Docker2017`

	> **Note**: Be sure to substitute the IP Address of pod-X-lin02 Linux node

	> **Note**: If you see an x509 certificate is from an unknown source error the cert didn't copy correctly, just rerun the above command.

	```
	$ docker login <pod-X-lin02 Ip Address>
	Username: docker
	Password: Docker2017
	```

	You should see a `Login succeeded` message
>
Congratulations, your nodes are now configured to work with your DTR instance.


## <a name="task2"></a>Task 2: Deploy a Linux Web App

Now that we've built our cluster, let's deploy a couple of web apps. These are simple web pages that allow you to send a tweet. One is built on Linux using NGINX and the other is build on Windows Server 2016 using IIS.  

Let's start on the Linux node.

### <a name="task2.1"></a> Task 2.1: Clone the Demo Repo

![](./images/linux75.png)

1. SSH in to your Linux **A** node

2. Make sure you're in your home directory on your Linux VM

	`$ cd ~`

2. Use git to clone the workshop repository.

	```
	$ git clone https://github.com/mikegcoleman/hybrid-workshop.git
	```

	You should see something like this as the output:

	```
	Cloning into 'hybrid-workshop'...
	remote: Counting objects: 13, done.
	remote: Compressing objects: 100% (10/10), done.
	remote: Total 13 (delta 1), reused 10 (delta 1), pack-reused 0
	Unpacking objects: 100% (13/13), done.
	Checking connectivity... done.
	```

	You now have the necessary demo code on your Linux VM.

### <a name="task2.2"></a> Task 2.2: Build and Push the Linux Web App Image

![](./images/linux75.png)

1. Change into the `linux_tweet_app` directory.

	`$ cd ~/hybrid-workshop/linux_tweet_app/`

2. Use `docker build` to build your Linux tweet web app Docker image.

	`$ docker build -t <pod-X-lin02 IP Address>/docker/linux_tweet_app .`

	The `-t` tells Docker that you're going to store this image in the `docker` repo on your Docker Trusted Registry node

	> **Note**: Feel free to examine the Dockerfile in this directory if you'd like to see how the image is being built.

	Your output should be similar to what is shown below

	```
	Sending build context to Docker daemon  4.096kB
	Step 1/4 : FROM nginx:latest
	latest: Pulling from library/nginx
	ff3d52d8f55f: Pull complete
	b05436c68d6a: Pull complete
	961dd3f5d836: Pull complete
	Digest: sha256:12d30ce421ad530494d588f87b2328ddc3cae666e77ea1ae5ac3a6661e52cde6
	Status: Downloaded newer image for nginx:latest
	 ---> 3448f27c273f
	Step 2/4 : COPY index.html /usr/share/nginx/html
	 ---> 72d22997a765
	Removing intermediate container e262b9220942
	Step 3/4 : EXPOSE 80 443
	 ---> Running in 54e4ff1b39a6
	 ---> 2b5bd87894cd
	Removing intermediate container 54e4ff1b39a6
	Step 4/4 : CMD nginx -g daemon off;
	 ---> Running in 54020cdec942
	 ---> ed5f550fc339
	Removing intermediate container 54020cdec942
	Successfully built ed5f550fc339
	Successfully tagged  pod-1-lin02/docker/linux_tweet_app:latest
	```

4. Use `docker push` to upload your image up to Docker Trusted Registry.

	> **Note**: You should still be logged into DTR from the previous steps, but if not you will need to log in again.

	```
	$ docker push <pod-X-lin02 IP Address>/docker/linux_tweet_app
	```

	The output should be similar to the following:

	```
	The push refers to a repository [pod-1-lin02/docker/linux_tweet_app]
	feecabd76a78: Pushed
	3c749ee6d1f5: Pushed
	af5bd3938f60: Pushed
	29f11c413898: Pushed
	eb78099fbf7f: Pushed
	latest: digest: sha256:9a376fd268d24007dd35bedc709b688f373f4e07af8b44dba5f1f009a7d70067 size: 1363
	```

4. In your web browser head back to your DTR server and click `View Details` next to your `linux_tweet_app` repo to see the details of the repo.

5. Click on `Images` from the horizontal menu. Notice that your newly pushed image is now on your DTR.

### <a name="task2.3"></a> Task 2.3: Deploy the Web App using UCP

Now let's run our application by by creating a new service.

Services are application building blocks (although in many cases an application will only have one service, such as this example). Services are based on a single Docker image. Tasks are the individual Docker containers that execute the application. When you create a new service you instantiate at least one task automatically, but you can scale the number of tasks up to meet the needs of your service.

1. In your web browser navigate to your UCP server (`https://<pod-X-lin01 IP Address`)

2. In the left hand menu click `Services`

3. In the upper right corner click `Create Service`

4. Enter `linux_tweet_app` for the name.

4. Under `Image` enter the path to your image which should be `<linux node b fdqn/docker/linux_tweet_app`

5. From the left hand menu click `Scheduling`

6. A few lines down the screen click `Add Constraint+`.

	Constraints are used to tell UCP where to run workloads. They are based on labels - in this specific case we're using a built in label that tells us what OS a given node is running (`node.platform.os`). Since this is a Linux-based container we need to make sure it ends up on a Linux node.

7. Enter `node.platform.os == Linux` into the text field

	![](./images/linux_constraint.png)

8. From the left hand menu click `Network`

9. Click `Publish Port+`

	We need to open a port for our web server. Since port 80 is already used by UCP on one node, and DTR on the other, we'll need to pick an alternate port. We'll go with 8088.

10. Fill out the port fields as shown below

	![](./images/linux_ports.png)

11. Click `Confirm`

12. Click `Create` near the bottom right of the screen.

After a few seconds you should see a green dot next to your service name. Once you see you green dot you can  point your web browser to `http://<linux node c fqdn>:8088` to see your running website.

> **Note**: You want to go to `http://` not `https://`

## <a name="task3"></a>Task 3: Deploy a Windows Web App

Now we'll deploy the Windows version of the tweet app.

### <a name="task3.1"></a> Task 3.1: Create the dockerfile with Image2Docker

There is a Windows Server 2016 VHD that contains our Windows Tweet App stored in `c:\` on your cloud-based VM. We're going to use Image2Docker to scan the VHD, and create a Dockerfile. We'll build the Dockerfile as we did in the previous step, push it to DTR, and then deploy our Windows tweet app.

![](./images/windows75.png)

1. Move back to Windows node **A**, and open a PowerShell window.

2. Use Image2Docker's `ConvertTo-Dockerfile` command to create a dockerfile from the VHD.

	Copy and paste the command below into your Powershell window.

	```
	ConvertTo-Dockerfile -ImagePath c:\ws2016.vhd `
	                     -Artifact IIS `
	                     -OutputPath C:\windowstweetapp `
	                     -Verbose
	```


	* `ImagePath` specifies where the VHD can be found

	* `Artifact` specifies what feature or code to look for

	* `OutputPath` specifies where to write the dockerfile and other items

	* `Verbose` instructs the script to provide extra output.

When the process completes you'll find a dockerfile in `c:\windowstweetapp`


### <a name="task3.2"></a> Task 3.2: Build and Push Your Image to Docker Trusted Registry

![](./images/windows75.png)

1. CD into the directory where your Image2Docker files have been placed.

	`PS C:\Users\docker> cd c:\windowstweetapp\ `


2. Use `docker build` to build your Windows tweet web app Docker image.

	`$ docker build -t <pod-X-lin02 IP Address>/docker/windows_tweet_app .`

	> **Note**: Be sure to use the FQDN for Linux node **B** when you tag the image.

	> **Note**: Feel free to examine the Dockerfile in this directory if you'd like to see how the image is being built.

	Your output should be similar to what is shown below

	```
	PS C:\Users\docker\scm\hybrid-workshop\windows_tweet_app\docker build -t <linux node b fqdn>/docker/windows_tweet_app .

	Sending build context to Docker daemon  6.144kB
	Step 1/10 : FROM microsoft/windowsservercore
	 ---> 590c0c2590e4

	<output snipped>

	Step 10/10 : HEALTHCHECK CMD powershell -command     try {      $response = Invoke-WebRequest http://localhost -UseBasic
	Parsing;      if ($response.StatusCode -eq 200) { return 0}      else {return 1};     } catch { return 1 }
	 ---> Running in ab4dfee81c7e
	 ---> d74eead7f408
	Removing intermediate container ab4dfee81c7e
	Successfully built d74eead7f408
	Successfully tagged <linux node b fqdn>/docker/windows_tweet_app:latest
	```
	> **Note**: It will take a few minutes for your image to build. If it takes more than 5 minutes move into your powershell window and press `Enter`. Sometimes the Powershell window will not update the current status of the build process.

4. Log into Docker Trusted Registry

	```
	PS C:\Users\docker> docker login <pod-X-lin02 IP Address>
	Username: docker
	Password: Docker2017
	Login Succeeded
	```

5. Push your new image up to Docker Trusted Registry.

	```
	PS C:\Users\docker> docker push <pod-X-lin02 IP Address>/docker/windows_tweet_app
	The push refers to a repository [<pod-X-lin02 IP Address>/docker/windows_tweet_app]
	5d08bc106d91: Pushed
	74b0331584ac: Pushed
	e95704c2f7ac: Pushed
	669bd07a2ae7: Pushed
	d9e5b60d8a47: Pushed
	8981bfcdaa9c: Pushed
	25bdce4d7407: Pushed
	df83d4285da0: Pushed
	853ea7cd76fb: Pushed
	55cc5c7b4783: Skipped foreign layer
	f358be10862c: Skipped foreign layer
	latest: digest: sha256:e28b556b138e3d407d75122611710d5f53f3df2d2ad4a134dcf7782eb381fa3f size: 2825
	```
### <a name="task3.3"></a> Task 3.3: Deploy the Windows Web App
Now that we have our Windows Tweet App up on the DTR server, let's deploy it. It's going to be almost identical to how did the Linux version with a couple of small exceptions:

* We will use a constraint to put the workload on a Windows node instead of Linux
* Windows does not currently support ingress load balancing, so we'll exposer the ports in `host` mode using `dnsrr`

1. In your web browser navigate to your UCP server (`https://<pod-X-lin02 IP Address>`)

2. In the left hand menu click `Services`

3. In the upper right corner click `Create Service`

4. Enter `windows_tweet_app` for the name.

4. Under `Image` enter the path to your image which should be `<pod-X-lin02 IP Address>/docker/windows_tweet_app`

5. From the left hand menu click `Scheduling`

6. A few lines down the screen click `Add Constraint+`.

	Constraints are used to tell UCP where to run workloads. They are based on labels - in this specific case we're using a built in label that tells us what OS a given node is running (`node.platform.os`). Since this is a Linux-based container we need to make sure it ends up on a Linux node.

7. Enter `node.platform.os == windows` into the text field

8. From the left hand menu click `Network`

8. Set the `ENDPOINT SPEC` to `DNS Round Robin`. This tells the service to load balance using DNS. The alternative is VIP, which uses IPVS.

9. Click `Publish Port+`

	We need to open a port for our web server. This app runs on port 80 which is used by DTR so let's use 8082.

10. Fill out the port fields as shown below. **Be sure to set the `Publish Mode` to  `Host`**

	![](./images/windows_ports.png)

11. Click 'Confirm'

12. Click `Create` near the bottom right of the screen.

After a few seconds you should see a green dot next to your service name. Once you see you green dot you can  point your web browser to `http://<pod-X-win05 IP Address>:8082` to see your running website.

> **Note**: You want to go to your Windows node b, and use `http:// ` not `https:// `

## <a name="task4"></a> Deploying a Multi-OS Application

For our last exercise we'll use a docker compose file to deploy an application that uses a Java front end designed to be deployed on Linux, with a Microsoft SQL Server back end running on windows.

### <a name="task4.1"></a> Task 4.1: Examine the Docker Compose file

We'll use a Docker Compose file to instantiate our application. With this file we can define all our services and their parameters, as well as other Docker primatives such as networks.

Let's look at the Docker Compose file:

```
version: "3.2"

services:

  database:
    image: sixeyed/atsea-db:mssql
    ports:
      - mode: host
        target: 1433
    networks:
     - atsea
    deploy:
      endpoint_mode: dnsrr
      placement:
        constraints:
          - 'node.platform.os == windows'

  appserver:
    image: sixeyed/atsea-app:mssql
    ports:
      - target: 8080
        published: 8080
    networks:
      - atsea
    deploy:
      placement:
        constraints:
          - 'node.platform.os == linux'

networks:
  atsea:
```

There are two services. `appserver` is our web frontend written in Java, and `database` is our Microsoft SQL Server database. The rest of the commands should look familiar as they are very close to what we used when we deployed our tweet services manually.

One thing that is new is the creation of an overlay network (`atsea`). Overlay networks allow containers running on different hosts to communicate over a private software-defined network. In this case, the web frontend on our Linux host will use the `atsea` network to communicate with the database.

You may have used Docker Compose before to deploy multi-service applications, but with swarm we use a slightly different command: `docker stack`.

### <a name="task4.2"></a> Task 4.2: Deploy the Application

1. Move to the UCP console in your web browser

2. In the left hand menu click `stacks`

3. In the upper right click `Create Stack`

4. Enter `atsea` under `NAME`

5. Select `Services` under `MODE`

6. Select `SHOW VERBOSE COMPOSE OUTPUT`

7. Paste the compose file from above into the `COMPOSE.YML` box

8.  Click `Create`

	You will see some output to show the progress of your deployment, and then a banner will pop up at the bottom indicating your deployment was successful.

9. Click `Done`

The UI shows your stack (`atsea`) and that it's comprised of 2 services and 1 network.

1. Click on the `atsea` stack in the list

2. From the right side of the screen choose `Services` under `Inspect Resource`

	![](./images/inspect_resource.png)

	Here you can see your two services running. It may take a few minutes for the databased service to come up (the dot to turn green). Once it does move on to the next section.


### <a name="task4.3"></a> Task 4.3: Verify the Running Application

1. To see our running web site (an art store) visit `http://<pod-X-lin01:8080>`

	> **Note**: You want to go to `http:// ` not `https:// `

	The thumbnails you see displayed are actually pulled from the SQL database. This is how you know that the connection is working between the database and web front end.


This concludes our workshop, thanks for attending.