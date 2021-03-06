package digiocean

import (
	droplet "github.com/cloudlibz/gocloud/compute/droplet"
	digioceandns "github.com/cloudlibz/gocloud/dns/digioceandns"
	digioceanloadbalancer "github.com/cloudlibz/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanstorage "github.com/cloudlibz/gocloud/storage/digioceanstorage"
	digioceancontainer "github.com/cloudlibz/gocloud/container/digioceancontainer"
)

// DigitalOcean  struct represents Digital Ocean cloud provider.
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.Digioceanloadbalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
}
