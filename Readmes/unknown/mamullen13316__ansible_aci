Example Ansible roles to configure ACI and VMware ESX.

create_vmm.yml - Creates a VMM domain in ACI

create_vlan_pool.yml - Creates a VLAN Pool in ACI

vmm_to_aep.yml - Assigns a VMM domain to an AEP in ACI

create_tenant_and_anp.yml - Creates a tenant and Application Network Profile in ACI

add_host_to_dvs.yml - Adds a host to the DVS created by ACI in vCenter

change_vm_portgroup.yml - Changes the vNIC on virtual machines to use the port group that was pushed by ACI

Requirements:

1. install Cobra SDK:  http://www.cisco.com/c/en/us/td/docs/switches/datacenter/aci/apic/sw/1-x/api/python/install/b_Install_Cisco_APIC_Python_SDK_Standalone.html

2. clone git repository.  https://github.com/jedelman8/aci-ansible

3. Copy aci-ansible/library/aci_rest to Ansible library directory (found in ansible.cfg,  with default of /usr/share/my_modules/).  Make sure to un-comment the library variable in ansible.cfg. Also may need to add a .py extension, depending on the host machine configuration.

4. export PYTHONPATH="${PYTHONPATH}:/usr/share/my_modules"

Global variables used across all roles are located in the /roles/common/vars/main.yml file. 





 



