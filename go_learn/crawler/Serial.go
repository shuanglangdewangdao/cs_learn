package crawler

import "sync"

func Serial(url string, fetcher Fetcher, fetched map[string]bool) {
	if fetched[url]{
		return
	}
	fetched[url]=true
	url,err:=fetcher.Fetch(url)
	if err!=nil {
		return
	}
	for _.u:=range urls{
		Serial(u,fetcher,fetched)
	}
	return 
}

type fetchState struct {
	mutex sync.Mutex
	fetched map[string]bool
}

func ConcurrentMutex(url string,fetcher Fetcher,f *fetchState){
	f.mutex.Lock()
	already:=f.fetched[url]
	f.fetched[url]=true
	f.mutex.Unlock()
	if already{
		return
	}
}