package zoo

import "testing"

func TestAppSmokeTest(t *testing.T) {

	zooConfig := NewZooAppConfig(
		"testdata/availabilities",
		"testdata/mail.tmpl",
		"testdata/sendlist",
	)

	statusCode := Run(zooConfig)

	if statusCode != Ok {

		t.Fatalf("app reported a failure")

	}

}
