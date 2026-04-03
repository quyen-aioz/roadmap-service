package roadmapmodel

type Status string

const (
	StatusCompleted  Status = "completed"
	StatusInProgress Status = "in-progress"
	StatusComingSoon Status = "coming-soon"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusCompleted, StatusInProgress, StatusComingSoon:
		return true
	}
	return false
}

type GroupID string

const (
	GroupIDAiozNetwork  GroupID = "aioz-network"
	GroupIDAiozDepin    GroupID = "aioz-depin"
	GroupIDAiozAi       GroupID = "aioz-ai"
	GroupIDAiozStream   GroupID = "aioz-stream"
	GroupIDAiozStorage  GroupID = "aioz-storage"
	GroupIDAiozPin      GroupID = "aioz-pin"
	GroupIDAiozWallet   GroupID = "aioz-wallet"
	GroupIDAiozAds      GroupID = "aioz-ads"
	GroupIDAiozAiAgents GroupID = "aioz-ai-agents"
	GroupIDAiozBridge   GroupID = "aioz-bridge"
	GroupIDAiozDex      GroupID = "aioz-dex"
	GroupIDAiozExplorer GroupID = "aioz-explorer"
	GroupIDAiozTransfer GroupID = "aioz-transfer"
	GroupIDAiozVault    GroupID = "aioz-vault"
)

func (s GroupID) IsValid() bool {
	switch s {
	case GroupIDAiozNetwork,
		GroupIDAiozDepin,
		GroupIDAiozAi,
		GroupIDAiozStream,
		GroupIDAiozStorage,
		GroupIDAiozPin,
		GroupIDAiozWallet,
		GroupIDAiozAds,
		GroupIDAiozBridge,
		GroupIDAiozDex,
		GroupIDAiozExplorer,
		GroupIDAiozTransfer,
		GroupIDAiozVault:
		return true
	}
	return false
}

const RoadmapContentID = "singleton"
