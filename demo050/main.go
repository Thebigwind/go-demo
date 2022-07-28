package main

func main() {
	XLoggerInit("./temp.log", 0)

	StoneLogger.Infof("abcd Info:%s", "haha")
	StoneLogger.Infof("abcdef:%s", "xxx")
	StoneLogger.Debugf("abcde Debug:%s", "xxxDebug")
	StoneLogger.Errorf("abce Errorf:%s,code:%d,desc:%s", "xxx", 1001, "detail err")
	StoneLogger.Warnf("abcd Warf:%s", "xxxx")
	StoneLogger.Info("xxxxxx")
	StoneLogger.Error("xxxxxxxxxx Error:%s", "xx")
}
