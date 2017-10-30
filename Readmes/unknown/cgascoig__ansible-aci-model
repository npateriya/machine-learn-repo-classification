# ansible-aci-model
Ansible module and roles to declaratively define Cisco ACI policy model 

## Requirements
- Ansible (tested with 2.2.1.0)
- requests module (tested with 2.8.1)

## Usage
### aci_model module
The aci_model module (located in the aci-model role) allows the specification of a base distinguished name (DN) and desired model. For example, the following specifies that there should be an application profile named test-ap under the tenant ACILab:
```yaml

tasks:
      - name: test aci_model
        aci_model:
          username: "{{ username }}"
          password: "{{ password }}"
          url: "http://{{ inventory_hostname }}/"
          dn: uni/tn-ACILab/ap-test-simple
          body:
            fvAp:
              attributes:
                rn: ap-test-simple
                name: test-simple
                descr: test simple application profile
```

See the `test-simple-playbook.yml` for a full example of this. Use `ansible-playbook -i hosts test-simple-playbook.yml` to execute it. 

### Abstracted constructs
In order to simplify usage, abstracted constructs for any commonly used ACI construct can be built on top of the aci_model module. These constructs are built using action_plugins that are simply wrappers around aci_model. 

Currently example action_plugins exist for the following:
- aci_application_profile
- aci_bd
- aci_contract
- aci_epg
- aci_filter
- aci_tenant
- aci_vrf

An example of using these action_plugins can be found in `test-playbook.yml`. Use `ansible-playbook -i hosts test-playbook.yml` to execute it. 
