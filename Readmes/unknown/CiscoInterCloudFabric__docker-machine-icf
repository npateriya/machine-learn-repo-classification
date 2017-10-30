# docker-machine-icf
Cisco ICF Docker Machine plugin

# Build

make -C bin

sudo mv bin/docker-machine-driver-icf $(dirname `which docker-machine`)

# RUN

## Parameters

1. --icf-server [$ICF_SERVER]

   IP address of Cisco Intercloud For Business
2. --icf-username [$ICF_USERNAME]

   Tenant User name on Cisco Intercloud For Business
3. --icf-password [$ICF_PASSWORD]

   Tenant User password on Cisco Intercloud For Business
4. --icf-server-cert [$ICF_SERVER_CERT]

   HTTPS Server Certificate of Cisco Intercloud For Business
   
>   export ICF_SERVER_CERT=$(keytool -printcert --rfc -sslserver $ICF_SERVER:443)

5. --icf-vdc [$ICF_VDC]

   UUID of VDC assigned to the tenant
6. --icf-catalog [$ICF_CATALOG]

   UUID of Catalog (VM Template) published for the tenant
7. --icf-network [$ICF_NETWORK]

   UUID of Network published for the tenant
8. --icf-provider-access [$ICF_PROVIDER_ACCESS]

   Does the VM require Provider Access (Native service access)
9. --icf-ssh-user [$ICF_SSH_USER]

   Linux User name of the VM (based on Catalog)
10. --icf-ssh-password [$ICF_SSH_PASSWORD]

   Linux User password of the VM (based on Catalog)

## Create VM

   Set all the environment variables described above

   docker-machine create --driver icf <name of vm>

## Delete VM

   docker-machine delete <name of VM>

# Requirements

Requires ssh client and sshpass utilities to be present
