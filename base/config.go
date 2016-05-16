package base

type PoolConfig struct {
  AnalyzerTotal uint32
  DownloaderTotal uint32
}

func NewPoolConfig(analyzerTotal uint32, downloaderTotal uint32) PoolConfig {
  return PoolConfig{AnalyzerTotal: analyzerTotal, DownloaderTotal: downloaderTotal}
}

type ChanConfig struct {
  ReqChanLen uint32
  RespChanLen uint32
}

func NewChanConfig(reqLen uint32, respLen uint32) ChanConfig {
  return ChanConfig{ReqChanLen: reqLen, RespChanLen: respLen}
}
