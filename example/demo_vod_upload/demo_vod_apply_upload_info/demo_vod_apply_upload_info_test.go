package demo_vod_upload

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/volcengine/volc-sdk-golang/base"
	"github.com/volcengine/volc-sdk-golang/models/vod/request"
	"github.com/volcengine/volc-sdk-golang/service/vod"
)

func TestVod_ApplyUploadInfo(t *testing.T) {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := vod.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")

	space := "your space name"

	applyRequest := &request.VodApplyUploadInfoRequest{SpaceName: space}

	resp, _, err := instance.ApplyUploadInfo(applyRequest)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	if resp.GetResponseMetadata().GetError() != nil {
		fmt.Println(resp.ResponseMetadata.Error)
		return
	}

	bts, _ := json.Marshal(resp)
	fmt.Printf("\nresp = %s\n", bts)

	fmt.Println(resp.GetResult().GetData().UploadAddress.SessionKey)
	fmt.Println(resp.GetResponseMetadata().GetRequestId())
}
