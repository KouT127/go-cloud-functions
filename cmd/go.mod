module github.com/KouT127/go-cloud-functions

go 1.13

require (
	github.com/KouT127/go-cloud-functions/health v0.0.0
	github.com/KouT127/go-cloud-functions/image v0.0.0
)

replace (
	github.com/KouT127/go-cloud-functions/common => ../common
	github.com/KouT127/go-cloud-functions/health => ../health
	github.com/KouT127/go-cloud-functions/image => ../image
)
