package searching

import "time"

//SnapshotModel is a GO representation of the snapshot object
type SnapshotModel struct {
	SerialNumberInserv string `json:"serialNumberInserv"`
	System             struct {
		CompanyName string   `json:"companyName"`
		Model       string   `json:"model"`
		FullModel   string   `json:"fullModel"`
		OsVersion   string   `json:"osVersion"`
		Patches     []string `json:"patches"`
		Sp          struct {
			SpID      string `json:"spId"`
			SpModel   string `json:"spModel"`
			SpVersion string `json:"spVersion"`
		} `json:"sp"`
		ProductSKU    string `json:"productSKU"`
		ProductFamily string `json:"productFamily"`
		Recommended   struct {
			OsVersion string `json:"osVersion"`
		} `json:"recommended"`
		PortsHWInfo []struct {
			Nsp      string `json:"nsp"`
			Brand    string `json:"brand"`
			Model    string `json:"model"`
			Rev      string `json:"rev,omitempty"`
			Firmware string `json:"firmware"`
			Serial   string `json:"serial"`
		} `json:"portsHWInfo"`
	} `json:"system"`
	Capacity struct {
		Total struct {
			VirtualSizeTiB        float64 `json:"virtualSizeTiB"`
			SizeTiB               float64 `json:"sizeTiB"`
			FreeTiB               float64 `json:"freeTiB"`
			FreePct               float64 `json:"freePct"`
			AllocatedCapacityTiB  float64 `json:"allocatedCapacityTiB"`
			FailedCapacityTiB     int     `json:"failedCapacityTiB"`
			SystemAllocatedTiB    float64 `json:"systemAllocatedTiB"`
			VolumesAllocatedTiB   float64 `json:"volumesAllocatedTiB"`
			OtherAllocatedTiB     int     `json:"otherAllocatedTiB"`
			VolumesVirtualSizeTiB float64 `json:"volumesVirtualSizeTiB"`
			UsedSpaceTiB          float64 `json:"usedSpaceTiB"`
			CompactionRatio       float64 `json:"compactionRatio"`
			DedupeRatio           float64 `json:"dedupeRatio"`
			OverprovisioningRatio float64 `json:"overprovisioningRatio"`
		} `json:"total"`
		ByType struct {
			Ssd struct {
				SizeTiB         float64 `json:"sizeTiB"`
				FreeTiB         float64 `json:"freeTiB"`
				FreePct         float64 `json:"freePct"`
				UsedBalancedPct float64 `json:"usedBalancedPct"`
				SizeBalancedPct int     `json:"sizeBalancedPct"`
				LifeLeftPctMin  int     `json:"lifeLeftPctMin"`
			} `json:"ssd"`
		} `json:"byType"`
		ArrayType string `json:"arrayType"`
	} `json:"capacity"`
	Performance struct {
		PortBandwidthData struct {
			Total struct {
				DataRateKBPSAvg  int     `json:"dataRateKBPSAvg"`
				DataRateKBPSMax  int     `json:"dataRateKBPSMax"`
				DataRateKBPSMin  int     `json:"dataRateKBPSMin"`
				IopsAvg          int     `json:"iopsAvg"`
				IopsMax          int     `json:"iopsMax"`
				IopsMin          int     `json:"iopsMin"`
				IoSizeAvg        float64 `json:"ioSizeAvg"`
				IoSizeMax        int     `json:"ioSizeMax"`
				IoSizeMin        int     `json:"ioSizeMin"`
				QueueLengthAvg   int     `json:"queueLengthAvg"`
				QueueLengthMax   int     `json:"queueLengthMax"`
				QueueLengthMin   int     `json:"queueLengthMin"`
				ServiceTimeMSAvg int     `json:"serviceTimeMSAvg"`
				ServiceTimeMSMax int     `json:"serviceTimeMSMax"`
				ServiceTimeMSMin int     `json:"serviceTimeMSMin"`
			} `json:"total"`
			Read struct {
				DataRateKBPSAvg  int     `json:"dataRateKBPSAvg"`
				DataRateKBPSMax  int     `json:"dataRateKBPSMax"`
				DataRateKBPSMin  int     `json:"dataRateKBPSMin"`
				IopsAvg          int     `json:"iopsAvg"`
				IopsMax          int     `json:"iopsMax"`
				IopsMin          int     `json:"iopsMin"`
				IoSizeAvg        float64 `json:"ioSizeAvg"`
				IoSizeMax        int     `json:"ioSizeMax"`
				IoSizeMin        int     `json:"ioSizeMin"`
				QueueLengthAvg   int     `json:"queueLengthAvg"`
				QueueLengthMax   int     `json:"queueLengthMax"`
				QueueLengthMin   int     `json:"queueLengthMin"`
				ServiceTimeMSAvg int     `json:"serviceTimeMSAvg"`
				ServiceTimeMSMax int     `json:"serviceTimeMSMax"`
				ServiceTimeMSMin int     `json:"serviceTimeMSMin"`
			} `json:"read"`
			Write struct {
				DataRateKBPSAvg  int     `json:"dataRateKBPSAvg"`
				DataRateKBPSMax  int     `json:"dataRateKBPSMax"`
				DataRateKBPSMin  int     `json:"dataRateKBPSMin"`
				IopsAvg          int     `json:"iopsAvg"`
				IopsMax          int     `json:"iopsMax"`
				IopsMin          int     `json:"iopsMin"`
				IoSizeAvg        float64 `json:"ioSizeAvg"`
				IoSizeMax        int     `json:"ioSizeMax"`
				IoSizeMin        int     `json:"ioSizeMin"`
				QueueLengthAvg   int     `json:"queueLengthAvg"`
				QueueLengthMax   int     `json:"queueLengthMax"`
				QueueLengthMin   int     `json:"queueLengthMin"`
				ServiceTimeMSAvg int     `json:"serviceTimeMSAvg"`
				ServiceTimeMSMax int     `json:"serviceTimeMSMax"`
				ServiceTimeMSMin int     `json:"serviceTimeMSMin"`
			} `json:"write"`
		} `json:"portBandwidthData"`
		Summary struct {
			PortInfo struct {
				ReadServiceTimeColMillis  int `json:"readServiceTimeColMillis"`
				WriteServiceTimeColMillis int `json:"writeServiceTimeColMillis"`
				TotalServiceTimeColMillis int `json:"totalServiceTimeColMillis"`
			} `json:"portInfo"`
		} `json:"summary"`
	} `json:"performance"`
	Disks struct {
		Total struct {
			DiskCount       int `json:"diskCount"`
			DiskCountNormal int `json:"diskCountNormal"`
		} `json:"total"`
		ByType struct {
			Ssd struct {
				DiskCount            int    `json:"diskCount"`
				DiskCountBalancedPct int    `json:"diskCountBalancedPct"`
				DiskCountNormal      int    `json:"diskCountNormal"`
				State                string `json:"state"`
			} `json:"ssd"`
		} `json:"byType"`
		State string `json:"state"`
	} `json:"disks"`
	Nodes struct {
		NodeCount        int       `json:"nodeCount"`
		NodeCountOffline int       `json:"nodeCountOffline"`
		NodeCountMissing int       `json:"nodeCountMissing"`
		NodeTimeSkewSecs int       `json:"nodeTimeSkewSecs"`
		CageCount        int       `json:"cageCount"`
		CPUAvgMax        int       `json:"cpuAvgMax"`
		CPUMedianMax     int       `json:"cpuMedianMax"`
		BatteryExpiry    time.Time `json:"batteryExpiry"`
	} `json:"nodes"`
	Updated    time.Time `json:"updated"`
	Authorized struct {
		Tenants []string `json:"tenants"`
	} `json:"authorized"`
	Date time.Time `json:"date"`
}
