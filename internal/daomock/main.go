package daomock

//go:generate mockgen -destination=mock.go -package=$GOPACKAGE github.com/AmazingTalker/go-amazing/pkg/dao RecordDAO
