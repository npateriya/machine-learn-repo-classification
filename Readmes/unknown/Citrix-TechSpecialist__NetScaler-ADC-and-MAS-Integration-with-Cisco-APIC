# Cisco Application Policy Infrastructure Controller 2.2 with Citrix Integration v1 #

## About This Demonstration ##

This guide for the preconfigured Cisco Application Policy Infrastructure Controller (Cisco APICTM) demonstration includes:

  * [**Scenario 1**](./Scenario1): APIC System Overview and Operations
  * [**Scenario 2**](./Scenario2): Building a Single Tenant with a Single-Node Graph within the APIC via the Northbound API
  * [**Scenario 3**](./Scenario3): Building a Single Tenant with a Multi-Node Graph within the APIC via the Northbound API
  * [**Scenario 4**](./Scenario4): Removing Objects within the APIC via the Northbound API
  * [**Scenario 5**](./Scenario5): Building a Single Tenant with a Single-Node Graph within the APIC using Service Manager mode via the Northbound API
  * [**Scenario 6**](./Scenario6): Building a Single Tenant with a Multi-Node Graph within the APIC using Service Manager mode via the Northbound API
  * [**Appendix**](./Appendix) (not to be performed as stand-alone procedures)
    * [Reset APIC Simulator](./Appendix-A/)
    * [Fabric Discovery](./Appendix-B/)
    * [Add VMware Hosts to APIC DVS Manually](./Appendix-C/)
    * [Create a Device Cluster](./Appendix-D/)
    * [Attach Service Graphs to Tenants](./Appendix-E/)
    * [Add the Citrix vpx1 VM](./Appendix-F/)
    * [Create VMM Domain](./Appendix-G/)

## Customization ##

It is not necessary to run all of the scenarios in this demonstration. There are several possible combinations. For example, to demonstrate: 

### Building a single tenant with a single- vs a multi-node graph and applying application policies within APIC ###

Run [**Scenario 2**](./Scenario2) and [**Scenario 3**](./Scenario3). Between the two scenarios, run [**Scenario 4**](./Scenario4) to remove the APIC objects.

### Building a single tenant with a single- vs a multi-node graph and applying application policies with Citrix Netscaler MAS ###

Run [**Scenario 5**](./Scenario5) and [**Scenario 6**](./Scenario6). Between the two scenarios, run [**Scenario 4**](./Scenario4)to remove the APIC objects.

### Difference between applying application policies in APIC vs using Netscaler MAS ###

Run [**Scenario 2**](./Scenario2) and [**Scenario 5**](./Scenario5) (single-node graph) or [**Scenario 3**](./Scenario3) and [**Scenario 6**](./Scenario6) (multi-node graph). Between scenarios, run [**Scenario 4**](./Scenario4) to remove the APIC objects.

## Limitations ##

Certain features of Cisco APIC 2.2 are outside the scope of this demonstration, because the demonstration uses a simulated fabric rather than a physical fabric:

  * All configuration will be lost after a reboot of the APIC simulator
  * No traffic will pass between devices connected to the simulated fabric
  * Screen refresh may take slightly longer than expected
  * In [Scenario 1](./Scenario1), traceroute will only show from the Spines, not from each Leaf

## Requirements ##

Below are the requirements for this preconfigured demonstration.

* **Required** : Laptop
* **Optional** : [Cisco AnyConnect](http://www.cisco.com/c/en/us/products/security/anyconnect-secure-mobility-client/index.html)

## About This Solution

The **Cisco Application Policy Infrastructure Controller** ([Cisco APIC](http://www.cisco.com/c/en/us/products/cloud-systems-management/application-policy-infrastructure-controller-apic/index.html)) is the unifying point of automation and management for the Application Centric Infrastructure (ACI) fabric. The Cisco APIC provides centralized access to all fabric information, optimizes the application lifecycle for scale and performance, supporting flexible application provisioning across physical and virtual resources.

The **Citrix NetScaler Management and Analytics System** ([Citrix NetScaler MAS](https://www.citrix.com/products/netscaler-management-and-analytics-system/)) works together with the Cisco APIC, enabling the separation of Network (L2-L3) and Application (L4-L7) operations. In the following guide, demo scenarios 2 and 3 consolidate management and operations within the Cisco APIC. Conversly, scenarios 5 and 6 separate management and enable operational flexibility by allowing a Networking team to have exclusive control of L2-L3 policies with Cisco APIC in Service Manager mode, and the Server or Application team control of L4-L7 application policies with Citrix NetScaler MAS operating as an L4-L7 Services Device Manager.

For additional information, visit [www.cisco.com/go/apic](www.cisco.com/go/apic), and [www.citrix.com/netscaler](www.citrix.com/netscaler).

## Topology

The following is the virtual demonstration topology, which consists of the following virtual machines:

* APIC Simulator (version 2.2(1n))
	* APIC1, APIC2 and APIC3
	* Leaf1 and Leaf2
	* Spine1 and Spine2
* VMware Virtual Center 6.0 Server Appliance
	* ASAv – version 9.6(2)
  * vpx1 – version 11.1 (50.10.nc)
  * VMware ESXi 6.0 Host 1
* VMware ESXi 6.0 Host 2
* Windows 10 Workstation
* EMC vVNX Storage Device
* Linux Tools Repository (Red Hat Enterprise Linux Server 7.1)

**Figure 1.** Demonstration Topology

![Figure1](images/Figure1.png)

### Before Presenting ###

Cisco dCloud strongly recommends that you perform the tasks in this document with an active session before presenting in front of a live audience. This will allow you to become familiar with the structure of the document and content.
It may be necessary to schedule a new session after following this guide in order to reset the environment to its original configuration.

### Preparation is Key to a Successful Presentation ###

Follow the steps to schedule a session of the content and configure your presentation environment.

  1. Initiate your dCloud session. [Show Me How](https://dcloud-cms.cisco.com/help/ initiate-your-dcloud-session)
  
    >**NOTE**: It may take up to 10 minutes for your session to become active.
  
  2. For best performance, connect to the workstation with **Cisco AnyConnect VPN** [Show   Me How](https://dcloud-cms.cisco.com/help/install_anyconnect_pc_mac) and the **local RDP  client on your laptop** [Show Me How](https://dcloud-cms.cisco.com/help/ local_rdp_mac_windows)
  
    * Workstation 1: **198.18.133.36**, Username: DCLOUD/demouser,Password: **C1sco12345**
  
    > **NOTE**: You can also connect to the workstation using the Cisco dCloud Remote     Desktop client [Show Me How](https://dcloud-cms.cisco.com/help/access_demo_wkstn).  The  dCloud Remote Desktop client works best for accessing an active session with  minimal   interaction. However, many users experience connection and performance   issues with this   method.
  
  3. The fabric discovery is automatically started at demo setup. Double-click the **APIC   Login icon**, accept the SSL Privacy error by clicking **advanced** followed by **  Proceed to 198.18.133.200** (unsafe), and finally login with the following credentials:   Username: **admin**, Password: **C1sco12345**.
  
  4. Select **Fabric** from the top menu.
  
  5. Select **Inventory** from the top sub-menu.
  
  6. In the left menu, click **Fabric Membership** and check that you have the 4 devices  populated as shown in Figure 2. (IP addresses may vary.)
  
    > **NOTE**: The fabric discovery can take up to 15 minutes to complete. If you login  before 15 minutes have passed, all devices may not be fully discovered._
  
    **Figure 2.** Completed Fabric Membership
  
    ![Figure1](images/Figure2.png)
  
    > **NOTE:** To demonstrate Fabric Discovery, reset the APIC Simulator (see [Appendix](/ Appendix)) If only TEP-1-101 is present at login, see [Appendix](/Appendix) to   discover the Fabric.
  
  7. Start vSphere from the Task Bar by clicking the icon, and make sure the Use Windows  Credentials checkbox is checked. **Click Login.**
  
  8. Check that the vpx1 and mas1 virtual machines are present and running as below.
  
**Figure 3.** Virtual Center Inventory

![Figure1](images/Figure3.png)

  >**NOTE:** If the **vpx1** or the **mas1** VM is not present in the L4-L7 Services Resource Pool, [run the APIC Fabric Discovery script](/Appendix/Appendix-B).




