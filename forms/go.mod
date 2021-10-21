module github.com/OutBoulderCounty/icc-services/forms

go 1.17

// require github.com/OutBoulderCounty/icc-services/database v0.0.0-unpublished

require (
	github.com/aws/aws-lambda-go v1.27.0
	github.com/aws/aws-sdk-go v1.41.2
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

// replace github.com/OutBoulderCounty/icc-services/database v0.0.0-unpublished => ../database
