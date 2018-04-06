package lib

type (
	Bucket struct {
		Experiment Experiment
	}

	Experiment struct {
		Weight float32
		Name   string
	}

	User struct {
		ID string
		// User.profile is a map of experiments assigned to this user
		Profile map[string]Experiment
	}
)
