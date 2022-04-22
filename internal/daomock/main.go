package daomock

//go:generate mockgen -destination=mock.go -package=$GOPACKAGE github.com/AmazingTalker/bevis-chang/pkg/dao RecordDAO,MemberDAO
