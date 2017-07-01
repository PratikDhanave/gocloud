package amazonstorage

import (
	"encoding/xml"
	"fmt"
	auth "github.com/scorelab/gocloud-v2/auth"
	awsauth "github.com/scorelab/gocloud-v2/awsauth"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var timeNow = time.Now

func prepareVolume(params map[string]string, volume CreateVolume) {
	params["AvailabilityZone"] = volume.AvailZone
	if volume.SnapshotId != "" {
		params["SnapshotId"] = volume.SnapshotId
	}
	if volume.VolumeType != "" {
		params["VolumeType"] = volume.VolumeType
	}
	if volume.VolumeSize > 0 {
		params["Size"] = strconv.FormatInt(int64(volume.VolumeSize), 10)
	}
	if volume.IOPS > 0 {
		params["Iops"] = strconv.FormatInt(volume.IOPS, 10)
	}
	if volume.Encrypted {
		params["Encrypted"] = "true"
	}
}

func makeParams(action string) map[string]string {
	return makeParamsWithVersion(action, legacyAPIVersion)
}

func makeParamsVPC(action string) map[string]string {
	return makeParamsWithVersion(action, vpcAPIVersion)
}

//add version to params
func makeParamsWithVersion(action, version string) map[string]string {
	params := make(map[string]string)
	params["Action"] = action
	params["Version"] = version
	return params
}

func (amazonstorage *Amazonstorage) PrepareSignatureV2query(params map[string]string, Region string, resp interface{}) error {

	EC2Endpoint := "https://ec2." + Region + ".amazonaws.com"

	req, err := http.NewRequest("GET", EC2Endpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", timeNow().In(time.UTC).Format(time.RFC3339))
	req.URL.RawQuery = query.Encode()

	auth := map[string]string{"AccessKey": auth.Config.AWSAccessKeyID, "SecretKey": auth.Config.AWSSecretKey}

	awsauth.SignatureV2(req, auth)

	fmt.Println(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))

	return xml.NewDecoder(r.Body).Decode(resp)
}
