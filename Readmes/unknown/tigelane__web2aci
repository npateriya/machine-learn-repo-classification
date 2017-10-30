# web2aci
Web interface into Cisco ACI

You can either use these as they are or in a docker container.  If you want to install this on your own system without Docker.com, take look at the Dockerfile first.  This will show you all of the things I do to create a Docker image and should help you imensily in creating your own server.


If you want to get this working fast, head over to Docker and do it that way.  It's much quicker.  https://registry.hub.docker.com/u/tigelane/web2aci/

After you isntall Docker, use this line to get the repository:
docker pull dockercisco/aci

To run the new container on port 80 (http) run this line:
sudo docker run -p 80:80 tigelane/web2aci /usr/sbin/apache2ctl -D FOREGROUND &

If you already have a web server running, this this line to change to port 8080:
sudo docker run -p 8080:80 tigelane/web2aci /usr/sbin/apache2ctl -D FOREGROUND &
