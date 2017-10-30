# rsvp-lsp-automation
This is my first attempt at automating network provisioning with Python, YAML and Jinja2
I wanted to try something that would help in maintaining the RSVP-TE LSP full mesh.
Without automation, one has to touch every router when a new router joins the MPLS domain.
The tools needs to work with multiple vendors, so Juniper and Cisco are included in the example.
I also wanted to have a data file with routers of differing network roles, P vs PE.  Since P routers
aren't usually ingress or egress for LSPs, I want the logic to limit LSPs to PE routers.
