package common

type ResAuth struct {
	Caps   ResAuthCaps `json:"caps" required:"true" description:"The caps of the pool." valid:"Required"`
	Entity string      `json:"entity" required:"true" description:"The name of auth." valid:"Required"`
	Key    string      `json:"key" description:"The key of auth."`
}

type ResAuthCaps struct {
	Mon string `json:"mon" description:"The access permission of monitor."`
	Osd string `json:"osd" description:"The access permisstion of osd.'`
}

type ResAuthList struct {
	AuthDump []ResAuth `json:"auth_dump"`
}

type ResDf struct {
	Stats struct {
		TotalBytes      int64 `json:"total_bytes"`
		TotalUsedBytes  int64 `json:"total_used_bytes"`
		TotalAvailBytes int64 `json:"total_avail_bytes"`
	} `json:"stats"`
	Pools []struct {
		Name  string `json:"name"`
		Id    uint64 `json:"id"`
		Stats struct {
			KbUsed    uint64 `json:"kb_used"`
			BytesUsed uint64 `json:"bytes_used"`
			MaxAvail  uint64 `json:"max_avail"`
			Objects   uint64 `json:"objects"`
		} `json:"stats"`
	} `json:"pools"`
}

type ResFsmap struct {
	Epoch   uint64         `json:"epoch"`
	Id      uint64         `json:"id"`
	Up      uint64         `json:"up"`
	In      uint64         `json:"in"`
	Max     uint64         `json:"max"`
	Failed  uint64         `json:"failed"`
	Damaged uint64         `json:"damaged"`
	ByRank  []ResFsmapRank `json:"by_rank"`
}

type ResFsmapRank struct {
	FilesystemId uint64 `json:"filesystem_id"`
	Rank         uint64 `json:"rank"`
	Name         string `json:"name"`
	Status       string `json:"status"`
}

type ResMonStatus struct {
	Name            string                       `json:"name"`
	Rank            int64                        `json:"rank"`
	State           string                       `json:"state"`
	ElectionEpoch   int64                        `json:"election_epoch"`
	Quorum          []int64                      `json:"quorum"`
	OutsideQuorum   []string                     `json:"outside_quorum"`
	ExtraProbePeers []ResMonStatusExtraProbePeer `json:"extra_probe_peers"`
	SyncProvider    []ResMonStatusSyncProvider   `json:"sync_provider"`
	Sync            ResMonStatusSync             `json:"sync"`
	Monmap          ResMonmap                    `json:"monmap"`
	ProviderKillAt  int64                        `json:"provider_kill_at"`
	RequesterKillAt int64                        `json:"requester_kill_at"`
}

type ResMonStatusExtraProbePeer struct {
	Peer string `json:"peer"`
}

type ResMonStatusSyncProvider struct {
	Cookie        uint64 `json:“cookie”`
	Entity        string `json:"entity"`
	Timeout       string `json:"timeout"`
	LastCommitted uint64 `json:"last_committed"`
	LastKey       string `json:"last_key"`
}

type ResMonStatusSync struct {
	SyncProvider     string `json:"sync_provider"`
	SyncCookie       uint64 `json:"sync_cookie"`
	SyncStartVersion uint64 `json:"sync_start_version"`
}

type ResMonmap struct {
	Epoch    int            `json:"epoch"`
	Fsid     string         `json:"fsid"`
	Modified string         `json:"modified"`
	Created  string         `json:"created"`
	Mons     []ResMonmapMon `json:"mons"`
}

type ResMonmapMon struct {
	Addr string `json:"addr"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type ResOsdmap struct {
	Epoch          int64  `json:"epoch"`
	NumOsds        int64  `json:"num_osds"`
	NumUpOsds      int64  `json:"num_up_osds"`
	NumInOsds      int64  `json:"num_in_osds"`
	Full           bool   `json:"full"`
	Nearfull       bool   `json:"nearfull"`
	NumRemappedPgs uint64 `json:"num_remapped_pgs"`
}

type ResOsdTree struct {
	Nodes []struct {
		Id              int64   `json:"id"`
		Name            string  `json:"name"`
		Type            string  `json:"type"`
		TypeId          int64   `json:"type_id"`
		CrushWeight     float64 `json:"crush_weight"`
		Depth           uint64  `json:"depth"`
		Exists          uint64  `json:"exists"`
		Status          string  `json:"status"`
		Reweight        float64 `json:"reweight"`
		PrimaryAffinity float64 `json:"primary_affinity"`
		Children        []int64 `json:"children"`
	} `json:"nodes"`
}

type ResPgStat struct {
	NumBytes      uint64            `json:"num_bytes"`
	NumPgByState  []ResNumPgByState `json:"num_pg_by_state"`
	NumPgs        uint64            `json:"num_pgs"`
	RawBytes      uint64            `json:"raw_bytes"`
	RawBytesAvail uint64            `json:"raw_bytes_avail"`
	RawBytesUsed  uint64            `json:"raw_bytes_used"`
	Version       uint64            `json:"version"`
}

type ResPgmap struct {
	PgsByState []ResPgByState `json:"pgs_by_state"`
	Version    uint64         `json:"version"`
	NumPgs     uint64         `json:"num_pgs"`
	NumPools   uint64         `json:"num_pools"`
	NumObjects uint64         `json:"num_objects"`
	DataBytes  uint64         `json:"data_bytes"`
	BytesUsed  uint64         `json:"bytes_used"`
	BytesAvail uint64         `json:"bytes_avail"`
	BytesTotal uint64         `json:"bytes_total"`
}

type ResPgByState struct {
	StateName string `json:"state_name"`
	Count     uint64 `json:"count"`
}

type ResPoolVar struct {
	Pool                       string  `json:"pool"`
	PoolId                     int64   `json:"pool_id"`
	PgNum                      int64   `json:"pg_num"`
	PgpNum                     int64   `json:"pgp_num"`
	Auid                       int64   `json:"auid"`
	Size                       int64   `json:"size"`
	MinSize                    int64   `json:"min_size"`
	CrashReplayInterval        int64   `json:"crash_replay_interval"`
	CrushRule                  string  `json:"crush_rule"`
	CrushRuleset               int64   `json:"crush_ruleset"`
	Hashpspool                 string  `json:"hashpspool"`
	Nodelete                   string  `json:"nodelete"`
	Nopgchange                 string  `json:"nopgchange"`
	Nosizechange               string  `json:"nosizechange"`
	WriteFadviseDontneed       string  `json:"write_fadvise_dontneed"`
	Noscrub                    string  `json:"noscrub"`
	NodeepScrub                string  `json:"nodeep_scrub"`
	HitSetPeriod               int64   `json:"hit_set_period"`
	HitSetCount                int64   `json:"hit_set_count"`
	HitSetType                 string  `json:"hit_set_type"`
	HitSetFpp                  float64 `json:"hit_set_fpp"`
	UseGmtHitset               bool    `json:"use_gmt_hitset"`
	TargetMaxObjects           uint64  `json:"target_max_objects"`
	TargetMaxBytes             uint64  `json:"target_max_bytes"`
	CacheTargetDirtyRatioMicro uint64  `json:"cache_target_dirty_ratio_micro"`
	CacheTargetDirtyRatio      float64 `json:"cache_target_dirty_ratio"`
	CacheTargetFullRatioMicro  uint64  `json:"cache_target_full_ratio_micro"`
	CacheTargetFullRatio       float64 `json:"cache_target_full_ratio"`
	CacheMinFlushAge           uint64  `json:"cache_min_flush_age"`
	CacheMinEvictAge           uint64  `json:"cache_min_evict_age"`
	ErasureCodeProfile         string  `json:"erasure_code_profile"`
	MinReadRecencyForPromote   int64   `json:"min_read_recency_for_promote"`
	MinWriteRecencyForPromote  int64   `json:"min_write_recency_for_promote"`
	FastRead                   int64   `json:"fast_read"`
	HitSetGradeDecayRate       int64   `json:"hit_set_grade_decay_rate"`
	HitSetSearchLastN          int64   `json:"hit_set_search_last_n"`
	ScrubMinInterval           string  `json:"scrub_min_interval"`
	ScrubMaxInterval           string  `json:"scrub_max_interval"`
	DeepScrubInterval          string  `json:"deep_scrub_interval"`
	RecoveryPriority           string  `json:"recovery_priority"`
	RecoveryOpPriority         string  `json:"recovery_op_priority"`
	ScrubPriority              string  `json:"scrub_priority"`
}

type ResPoolStat struct {
	PoolName     string                  `json:"pool_name"`
	PoolId       uint64                  `json:"pool_id"`
	Recovery     ResPoolStatRecovery     `json:"recovery"`
	RecoveryRate ResPoolStatRecoveryRate `json:"recovery_rate"`
	ClientIoRate ResPoolStatClientIoRate `json:"client_io_rate"`
	CacheIoRate  ResPoolStatCacheIoRate  `json:"cache_io_rate"`
}

type ResPoolStats []ResPoolStat

type ResPoolStatCacheIoRate struct {
	FlushBytesSec    int64 `json:"flush_bytes_sec"`
	EvictBytesSec    int64 `json:"evict_bytes_sec"`
	PromoteOpPerSec  int64 `json:"promote_op_per_sec"`
	NumFlushModeLow  int64 `json:"num_flush_mode_low"`
	NumFlushModeHigh int64 `json:"num_flush_mode_high"`
	NumEvictModeSome int64 `json:"num_evict_mode_some"`
	NumEvictModeFull int64 `json:"num_evict_mode_full"`
}

type ResPoolStatRecovery struct {
	DegradedObjects  uint64  `json:"degraded_objects"`
	DegradedTotal    uint64  `json:"degraded_total"`
	DegradedRatio    float32 `json:"degraded_ratio"`
	MisplacedObjects uint64  `json:"misplaced_objects"`
	MisplacedTotal   uint64  `json:"misplaced_total"`
	MisplacedRatio   float32 `json:"misplaced_ratio"`
	UnfoundObjects   uint64  `json:"unfound_objects"`
	UnfoundTotal     uint64  `json:"unfound_total"`
	UnfoundRatio     float32 `json:"unfound_ratio"`
}

type ResPoolStatRecoveryRate struct {
	RecoveringObjectsPerSec int64 `json:"recovering_objects_per_sec"`
	RecoveringBytesPerSec   int64 `json:"recovering_bytes_per_sec"`
	RecoveringKeysPerSec    int64 `json:"recovering_keys_per_sec"`
	NumObjectsRecovered     int64 `json:"num_objects_recovered"`
	NumBytesRecovered       int64 `json:"num_bytes_recovered"`
	NumKeysRecovered        int64 `json:"num_keys_recovered"`
}

type ResPoolStatClientIoRate struct {
	ReadBytesSec  int64 `json:"read_bytes_sec"`
	WriteBytesSec int64 `json:"write_bytes_sec"`
	ReadOpPerSec  int64 `json:"read_op_per_sec"`
	WriteOpPerSec int64 `json:"write_op_per_sec"`
}

type ResPoolLs []string

type ResPoolLsDetail []ResPoolDetail

type ResPoolDetail struct {
	PoolName                       string          `json:"pool_name"`
	Flags                          uint64          `json:"flags"`
	FlagsNames                     string          `json:"flags_names"`
	Type                           int64           `json:"type"`
	Size                           int64           `json:"size"`
	MinSize                        int64           `json:"min_size"`
	CrushRuleset                   int64           `json:"crush_ruleset"`
	ObjectHash                     int64           `json:"object_hash"`
	PgNum                          uint64          `json:"pg_num"`
	PgPlacementNum                 uint64          `json:"pg_placement_num"`
	CrashReplayInterval            uint64          `json:"crash_reply_interval"`
	LastChange                     string          `json:"last_change"`
	LastForceOpResend              string          `json:"last_force_op_resend"`
	LastForceOpResendPreluminous   string          `json:"last_force_op_resend_preluminous"`
	Auid                           uint64          `json:"auid"`
	SnapMode                       string          `json:"snap_mode"`
	SnapSeq                        uint64          `json:"snap_seq"`
	SnapEpoch                      uint64          `json:"snap_epoch"`
	PoolSnaps                      []ResPoolSnap   `json:"pool_snaps"`
	RemovedSnaps                   string          `json:"removed_snaps"`
	QuotaMaxBytes                  uint64          `json:"quota_max_bytes"`
	QuotaMaxObjects                uint64          `json:"quota_max_objects"`
	Tiers                          []uint64        `json:"tiers"`
	TierOf                         int64           `json:"tier_of"`
	ReadTier                       int64           `json:"read_tier"`
	WriteTier                      int64           `json:"write_tier"`
	CacheMode                      string          `json:"cache_mode"`
	TargetMaxBytes                 uint64          `json:"target_max_bytes"`
	TargetMaxObjects               uint64          `json:"target_max_objects"`
	CacheTargetDirtyRatioMicro     uint64          `json:"cache_target_dirty_ratio_micro"`
	CacheTargetDirtyHighRatioMicro uint64          `json:"cache_target_dirty_high_ratio_micro"`
	CacheTargetFullRatioMicro      uint64          `json:"cache_target_full_ratio_micro"`
	CacheMinFlushAge               uint64          `json:"cache_min_flush_age"`
	CacheMinEvictAge               uint64          `json:"cache_min_evict_age"`
	ErasureCodeProfile             string          `json:"erasure_code_profile"`
	HitSetParams                   ResHitSetParams `json:"hit_set_params"`
	HitSetPeriod                   uint64          `json:"hit_set_period"`
	HitSetCount                    uint64          `json:"hit_set_count"`
	UseGmtHitset                   bool            `json:"use_gmt_hitset"`
	MinReadRecencyForPromote       uint64          `json:"min_read_recency_for_promote"`
	MinWriteRecencyForPromote      uint64          `json:"min_write_recency_for_promote"`
	HitSetGradeDecayRate           uint64          `json:"hit_set_grade_decay_rate"`
	HitSetSearchLastN              uint64          `json:"hit_set_search_last_n"`
	GradeTable                     []uint64        `json:"grade_table"`
	StripeWidth                    uint64          `json:"stripe_width"`
	ExpectedNumObjects             uint64          `json:"expected_num_objects"`
	FastRead                       bool            `json:"fast_read"`
	Options                        ResOptions      `json:"options"`
}

type ResOptions struct {
	ScrubMinInterval         float64 `json:"scrub_min_interval"`
	ScrubMaxInterval         float64 `json:"scrub_max_interval"`
	DeepScrubInterval        float64 `json:"deep_scrub_interval"`
	RecoveryPriority         int64   `json:"recovery_priority"`
	RecoveryOpPriority       int64   `json:"recovery_op_priority"`
	ScrubPriority            int64   `json:"scrub_priority"`
	CompressionMode          string  `json:"compression_mode"`
	CompressionAlgorithm     string  `json:"compression_algorithm"`
	CompressionRequiredRatio float64 `json:"compression_required_ratio"`
	CompressionMaxBlobSize   int64   `json:"compression_max_blob_size"`
	CompressionMinBlobSize   int64   `json:"compression_min_blob_size"`
	CsumType                 int64   `json:"csum_type"`
	CsumMaxBlock             int64   `json:"csum_max_block"`
	CsumMinBlock             int64   `json:"csum_min_block"`
}

type ResHitSetParams struct {
	Type string `json:"type"`
}

type ResPoolSnap struct {
	Name   string `json:"name"`
	Snapid uint64 `json:"snapid"`
	Stamp  string `json:"stamp"`
}

type ResNumPgByState struct {
	Name string `json:"name"`
	Num  uint64 `json:"num"`
}

type ResVersion struct {
	Version string `json:"version"`
}

type ResHealth struct {
	Health        ResHealthHealth `json:"health"`
	Timechecks    ResTimechecks   `json:"timechecks"`
	Summary       []ResSummary    `json:"summary"`
	OverallStatus string          `json:"overall_status"`
	Detail        []string        `json:"detail"`
}

type ResHealthHealth struct {
	HealthServices []ResHealthService `json:"health_services"`
}

type ResHealthService struct {
	Mons []ResHealthServiceMon `json:"mons"`
}

type ResHealthServiceMon struct {
	Name         string        `json:"name"`
	KbTotal      int64         `json:"kb_total"`
	KbUsed       int64         `json:"kb_used"`
	KbAvail      int64         `json:"kb_avail"`
	AvailPercent int           `json:"avail_percent"`
	LastUpdated  string        `json:"last_updated"`
	Health       string        `json:"health"`
	StoreStats   ResStoreStats `json:"store_stats"`
}

type ResStoreStats struct {
	BytesTotal  int64  `json:"bytes_total"`
	BytesSst    int64  `json:"bytes_sst"`
	BytesLog    int64  `json:"bytes_log"`
	BytesMisc   int64  `json:"bytes_misc"`
	LastUpdated string `json:"last_updated"`
}

type ResSummary struct {
	Severity string `json:"severity"`
	Summary  string `json:"summary"`
}

type ResTimechecks struct {
	Epoch       uint64             `json:"epoch"`
	Round       int64              `json:"round"`
	RoundStatus string             `json:"round_status"`
	Mons        []ResTimechecksMon `json:"mons"`
}

type ResTimechecksMon struct {
	Name     string  `json:"name"`
	Skew     float32 `json:"skew"`
	Lantency float32 `json:"lantency"`
	Health   string  `json:"health"`
}

type ResStatus struct {
	Health        ResHealth       `json:"health"`
	Fsid          string          `json:"fsid"`
	ElectionEpoch uint64          `json:"election_epoch"`
	Quorum        []int           `json:"quorum"`
	QuorumNames   []string        `json:"quorum_names"`
	Monmap        ResMonmap       `json:"monmap"`
	Osdmap        ResStatusOsdmap `json:"osdmap"`
	Pgmap         ResPgmap        `json:"pgmap"`
	Fsmap         ResFsmap        `json:"fsmap"`
}

type ResStatusOsdmap struct {
	Osdmap ResOsdmap `json:"osdmap"`
}

type ResQuorumStatus struct {
	ElectionEpoch    int64     `json:"election_epoch"`
	Quorum           []uint64  `json:"quorum"`
	QuorumNames      []string  `json:"quorum_names"`
	QuorumLeaderName string    `json:"quorum_leader_name"`
	Monmap           ResMonmap `json:"monmap"`
}
