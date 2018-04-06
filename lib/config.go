package lib

type (
	Config struct {
		Experiments map[string]Experiment
		// Parse contents of X-Bucketeer-Id header and return a bucket
		Parser func(string) (Bucket, error)
	}
)
