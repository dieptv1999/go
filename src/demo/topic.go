package main
//
//var (
//	TopicList map[string]*Topic
//)
//
//type Topic struct {
//	Name     string
//	Channels []chan string
//}
//
//func Init() {
//	TopicList = make(map[string]*Topic)
//
//	TopicList["golang"] = &Topic{"golang", []chan string{}}
//	TopicList["python"] = &Topic{"python", []chan string{}}
//	TopicList["java"] = &Topic{"java", []chan string{}}
//	TopicList["react"] = &Topic{"react", []chan string{}}
//}
//
//func Pub(Topics []string, content string) {
//	for _, v := range Topics {
//		if u, ok := TopicList[v]; ok {
//			for _, c := range u.Channels {
//				c <- content + v
//			}
//		}
//	}
//}
//
//func Sub(Topics []string) []chan string {
//	var chann []chan string
//	for _, v := range Topics {
//		if u, ok := TopicList[v]; ok {
//			c := make(chan string, 5)
//			u.Channels = append(u.Channels, c)
//			chann = append(chann, c)
//		}
//	}
//	return chann
//}
