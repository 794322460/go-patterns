package proxy

import "fmt"

/*
proxy pattern provide object control access to another object, intercepting all calls

在本例中， 代理模式有助于实现延迟初始化， 并对低效的第三方 YouTube 集成程序库进行缓存。
*/

type ThirdPartyYouTubeLib interface {
	GetVideo(videoId int) Video
}

type Video struct {
	Id    int
	Title string
	Url   string
}

func NewVideo(id int, title string) Video {
	return Video{Id: id, Title: title, Url: "http://random.mp4"}
}

type ThirdPartyYouTubeLibImpl struct{}

func NewThirdPartyYouTubeLib() ThirdPartyYouTubeLib {
	return &ThirdPartyYouTubeLibImpl{}
}

func (lib *ThirdPartyYouTubeLibImpl) GetVideo(videoId int) Video {
	return NewVideo(videoId, "random Title")
}

// =========================================================================== //

type ThirdPartyYouTubeLibProxy struct {
	ThirdPartyYouTubeLib
	cacheVideos map[int]Video
}

func NewThirdPartyYouTubeLibProxy() ThirdPartyYouTubeLib {
	return &ThirdPartyYouTubeLibProxy{ThirdPartyYouTubeLib: NewThirdPartyYouTubeLib(), cacheVideos: make(map[int]Video, 100)}
}

func (p *ThirdPartyYouTubeLibProxy) GetVideo(videoId int) Video {
	fmt.Println("代理：先尝试从缓存获取... videoId=", videoId)
	if video, exists := p.cacheVideos[videoId]; exists {
		fmt.Println("代理：命中缓存，直接返回... video=", video)
		return video
	}
	video := p.ThirdPartyYouTubeLib.GetVideo(videoId)
	fmt.Println("代理：没有拿到数据，rpc调用获取，并放入缓存... video=", video)
	p.cacheVideos[videoId] = video
	return video
}
