package proxy

import (
	"fmt"
	"testing"
)

func TestProxyObject(t *testing.T) {

	proxy := NewThirdPartyYouTubeLibProxy()
	video := proxy.GetVideo(1)
	fmt.Println("video=", video)
	video = proxy.GetVideo(1)
	fmt.Println("video=", video)
}
