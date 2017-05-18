package model

type ResponseStatus struct {
	Health struct {
		Health struct {
			HealthServices []struct {
				Mons []struct {
					Name         string `json:"name"`
					KbTotal      int64  `json:"kb_total"`
					KbUsed       int64  `json:"kb_used"`
					KbAvail      int64  `json:"kb_avail"`
					AvailPercent int8   `json:"avail_percent"`
					LastUpdated  string `json:"last_updated"`
					Health       string `json:"health"`
					StoreStats   struct {
						BytesTotal  int64  `json:"bytes_total"`
						BytesSst    int64  `json:"bytes_sst"`
						BytesLog    int64  `json:"bytes_log"`
						BytesMisc   int64  `json:"bytes_misc"`
						LastUpdated string `json:"last_updated"`
					} `json:"store_stats"`
				} `json:"mons"`
			} `json:"health_services"`
		} `json:"health"`
		Timeckecks struct {
			Epoch       int    `json:"epoch"`
			Round       int    `json:"round"`
			RoundStatus string `json:"round_status"`
			Mons        []struct {
				Name     string  `json:"name"`
				Skew     float32 `json:"skew"`
				Lantency float32 `json:"lantency"`
				Health   string  `json:"health"`
			} `json:"mons"`
		} `json:"timechecks"`
		Summary []struct {
			Severity string `json:"severity"`
			Summary  string `json:"summary"`
		} `json:"struct"`
		OverallStatus string `json:"overall_status"`
		Detail        []struct {
		} `json:"detail"`
	} `json:"health"`
	Fsid          string   `json:"fsid"`
	ElectionEpoch int      `json:"election_epoch"`
	Quorum        []int    `json:"quorum"`
	QuorumNames   []string `json:"quorum_names"`
	Monmap        struct {
		Epoch    int    `json:"epoch"`
		Fsid     string `json:"fsid"`
		Modified string `json:"modified"`
		Created  string `json:"created"`
		Mons     []struct {
			Rank int    `json:"rank"`
			Name string `json:"name"`
			Addr string `json:"addr"`
		} `json:"mons"`
	} `json:"monmap"`
	Osdmap struct {
		Osdmap struct {
			Epoch          int  `json:"epoch"`
			NumOsds        int  `json:"num_osds"`
			NumUpOsds      int  `json:"num_up_osds"`
			NumInOsds      int  `json:"num_in_osds"`
			Full           bool `json:"full"`
			Nearfull       bool `json:"nearfull"`
			NumRemappedPgs int  `json:"num_remapped_pgs"`
		} `json:"osdmap"`
	} `json:"osdmap"`
	Pgmap struct {
		PgsByState []struct {
			StateName string `json:"state_name"`
			Count     int    `json:"count"`
		} `json:"pgs_by_state"`
		Version         int     `json:"version"`
		NumPgs          int     `json:"num_pgs"`
		DataBytes       int64   `json:"data_bytes"`
		BytesUsed       int64   `json:"bytes_used"`
		BytesAvail      int64   `json:"bytes_avail"`
		BytesTotal      int64   `json:"bytes_total"`
		DegradedObjects int     `json:"degraded_objects"`
		DegradedTotal   int     `json:"degraded_total"`
		DegradedRatio   float32 `json:"degraded_ratio"`
	} `json:"pgmap"`
	Fsmap struct {
		Epoch  int        `json:"epoch"`
		ByRank []struct{} `json:"by_rank"`
	} `json:"fsmap"`
}

type ResponseVersion struct {
	Version string `json:"version"`
}

type ResponseOsdPoolLs []struct {
	PoolName            string `json:"pool_name"`
	Flags               int    `json:"flags"`
	FlagsNames          string `json:"flags_names"`
	Type                int    `json:"type"`
	Size                int    `json:"size"`
	MinSize             int    `json:"min_size"`
	CrushRuleset        int    `json:"cresh_ruleset"`
	ObjectHash          int    `json:"object_hash"`
	PgNum               int    `json:"pg_num"`
	PgPlacementNum      int    `json:"pg_placement_num"`
	CrashReplayInterval int    `json:"crash_replay_interval"`
	LastChange          string `json:"last_change"`
	LastForceOpResend   string `json:"last_force_op_resend"`
	Auid                uint64 `json:"auid"`
	SnapMode            string `json:"snap_mode"`
	SnapSeq             int    `json:"snap_seq"`
	SnapEpoch           int    `json:"snap_epoch"`
	PoolSnaps           []struct {
		Snapid int    `json:"snapid"`
		Stamp  string `json:"stamp"`
		Name   string `json:"name"`
	} `json:"pool_snaps"`
	RemovedSnaps    string `json:"removed_snaps"`
	QuotaMaxBytes   int64  `json:"quota_max_bytes"`
	QuotaMaxObjects int64  `json:"quota_max_objects"`
	Tiers           []struct {
	} `json:"tiers"`
	TiersOf                        int    `json:"tier_of"`
	ReadTier                       int    `json:"read_tier"`
	WriteTier                      int    `json:"write_tier"`
	CacheMode                      string `json:"cache_mode"`
	TargetMaxBytes                 int64  `json:"target_max_bytes"`
	TargetMaxObjects               int64  `json:"target_max_objects"`
	CacheTargetDirtyRatioMicro     int    `json:"cache_target_dirty_ratio_micro"`
	CacheTargetDirtyHighRatioMicro int    `json:"cache_target_dirty_high_ratio_micro"`
	CacheTargetFullRatioMicro      int    `json:"cache_target_full_ratio_micro"`
	CacheMinFlushAge               int    `json:"cache_min_flush_age"`
	CacheMinEvictAge               int    `json:"cache_min_evict_age"`
	ErasureCodeProfile             string `json:"erasure_code_profile"`
	HitSetParams                   struct {
		Type string `json:"type"`
	} `json:"hit_set_params"`
	HitSetPeriod              int        `json:"hit_set_period"`
	HitSetCount               int        `json:"hit_set_count"`
	UseGmtHitset              bool       `json:"use_gmt_hitset"`
	MinReadRecencyForPromote  int        `json:"min_read_recency_for_promote"`
	MinWriteRecencyForPromote int        `json:"min_write_recency_for_promote"`
	HitSetGradeDecayRate      int        `json:"hit_set_grade_decay_rate"`
	HitSetSearchLastN         int        `json:"hit_set_search_last_n"`
	GradeTable                []struct{} `json:"grade_table"`
	StripeWidth               int        `json:"stripe_width"`
	ExpectedNumObjects        int        `json:"expected_num_objects"`
	FastRead                  bool       `json:"fast_read"`
	Options                   struct{}   `json:"options"`
}

type ResponsePoolGet struct {
	Pool    string `json:"pool"`
	PoolId  int    `json:"pool_id"`
	Size    int    `json:"size"`
	MinSize int    `json:"min_size"`
}
