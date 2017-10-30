css_lb_configurator
===================

### Simple python script that helps create Cisco CSS 11500 load balancers

This script is being developed by Fernando Paci @ The Warranty Group

All the modifications of the code to use classes by [ignacionf](https://github.com/ignacionf)

Requirements:

+ Python 2. Tested with Python 2.6 $ 2.7. 
+ File with the following structure:

```
[Global]
redirect_url=dev.example.com
vip_ip=10.10.0.1
app_name=example
env=dev
service_ip=192.168.1.10, 192.168.1.11

[Template]
template_service=./templates/ecom/service.tmpl
template_keepalive=./templates/ecom/keepalive.tmpl
template_group=./templates/ecom/group.tmpl
template_owner=./templates/ecom/owner.tmpl
```

Usage:

`python gen_lb_config.py dev.example.com.cfg`

Output:

```
$ python gen_lb_config.py dev.example.com.cfg 
!************************* KEEPALIVE *************************
keepalive 10_dev1_ka_example
  description "keepalive for 10_dev1_example"
  ip address 192.168.1.10
  port 8080
  type http
  uri /JsSf3LoWJn8aM3sWQ2LBsx4mM6/index.jsp
  active


!************************** SERVICE **************************
service 10_dev1_example
  ip address 192.168.1.10
  keepalive type named 10_dev1_ka_example
  protocol tcp
  active


!************************* KEEPALIVE *************************
keepalive 11_dev2_ka_example
  description "keepalive for 11_dev2_example"
  ip address 192.168.1.11
  port 8080
  type http
  uri /JsSf3LoWJn8aM3sWQ2LBsx4mM6/index.jsp
  active


!************************** SERVICE **************************
service 11_dev2_example
  ip address 192.168.1.11
  keepalive type named 11_dev2_ka_example
  protocol tcp
  active


!*************************** GROUP ***************************
group 1_dev_example
  vip address 10.10.0.1
  add destination service 10_dev1_example
  add destination service 11_dev2_example
  active

!*************************** OWNER ***************************
owner 1_dev_example

  content 1_dev_example
    vip address 10.10.0.1
    add service 10_dev1_example
    add service 11_dev2_example
    add dns dev.example.com
    sticky-inact-timeout 480
    port 443
    protocol tcp
    advanced-balance sticky-srcip
    active

  content 1_dev_example_redir
    vip address 10.10.0.1
    sticky-inact-timeout 480
    port 80
    protocol tcp
    redirect "https://dev.example.com"
    active
```
