package api

import (
	"time"
)

// Snapshot is the generated struct that we can use for unmarshalling strings from the database.
type Snapshot struct {
	SerialNumberInserv string `json:"serialNumberInserv"`
	System             struct {
		CompanyName string   `json:"companyName"`
		Model       string   `json:"model"`
		FullModel   string   `json:"fullModel"`
		OsVersion   string   `json:"osVersion"`
		Patches     []string `json:"patches"`
		Sp          struct {
			SpID      string   `json:"spId"`
			SpModel   string   `json:"spModel"`
			SpVersion string   `json:"spVersion"`
			SpPatches []string `json:"spPatches"`
		} `json:"sp"`
		ProductSKU    string `json:"productSKU"`
		ProductFamily string `json:"productFamily"`
		Recommended   struct {
			CriticalPatches []string `json:"criticalPatches"`
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
			FailedCapacityTiB     float64 `json:"failedCapacityTiB"`
			SystemAllocatedTiB    float64 `json:"systemAllocatedTiB"`
			VolumesAllocatedTiB   float64 `json:"volumesAllocatedTiB"`
			OtherAllocatedTiB     float64 `json:"otherAllocatedTiB"`
			VolumesVirtualSizeTiB float64 `json:"volumesVirtualSizeTiB"`
			UsedSpaceTiB          float64 `json:"usedSpaceTiB"`
			CompactionRatio       float64 `json:"compactionRatio"`
			OverprovisioningRatio float64 `json:"overprovisioningRatio"`
		} `json:"total"`
		ByType struct {
			Ssd struct {
				SizeTiB         float64 `json:"sizeTiB"`
				FreeTiB         float64 `json:"freeTiB"`
				FreePct         float64 `json:"freePct"`
				UsedBalancedPct float64 `json:"usedBalancedPct"`
				SizeBalancedPct float64 `json:"sizeBalancedPct"`
				LifeLeftPctMin  float64 `json:"lifeLeftPctMin"`
			} `json:"ssd"`
		} `json:"byType"`
		ArrayType string `json:"arrayType"`
	} `json:"capacity"`
	Performance struct {
		PortBandwidthData struct {
			Total BandwidthData `json:"total"`
			Read BandwidthData `json:"read"`
			Write BandwidthData `json:"write"`
		} `json:"portBandwidthData"`
		Summary struct {
			PortInfo struct {
				ReadServiceTimeColMillis  float64 `json:"readServiceTimeColMillis"`
				WriteServiceTimeColMillis float64 `json:"writeServiceTimeColMillis"`
				TotalServiceTimeColMillis float64 `json:"totalServiceTimeColMillis"`
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
				DiskCount            float64 `json:"diskCount"`
				DiskCountBalancedPct float64 `json:"diskCountBalancedPct"`
				DiskCountNormal      float64 `json:"diskCountNormal"`
				State                string  `json:"state"`
			} `json:"ssd"`
		} `json:"byType"`
		State string `json:"state"`
	} `json:"disks"`
	Nodes struct {
		NodeCount        int `json:"nodeCount"`
		NodeCountOffline int `json:"nodeCountOffline"`
		NodeCountMissing int `json:"nodeCountMissing"`
		NodeTimeSkewSecs int `json:"nodeTimeSkewSecs"`
		CageCount        int `json:"cageCount"`
		CPUAvgMax        int `json:"cpuAvgMax"`
		CPUMedianMax     int `json:"cpuMedianMax"`
	} `json:"nodes"`
	Updated    time.Time `json:"updated"`
	Authorized struct {
		Tenants []string `json:"tenants"`
	} `json:"authorized"`
	Date time.Time `json:"date"`
}

// BandwidthData is what is in the Port Bandwidth Data part of the struct
type BandwidthData struct {
	DataRateKBPSAvg  float64 `json:"dataRateKBPSAvg"`
	DataRateKBPSMax  float64 `json:"dataRateKBPSMax"`
	DataRateKBPSMin  float64 `json:"dataRateKBPSMin"`
	IopsAvg          float64 `json:"iopsAvg"`
	IopsMax          float64 `json:"iopsMax"`
	IopsMin          float64 `json:"iopsMin"`
	IoSizeAvg        float64 `json:"ioSizeAvg"`
	IoSizeMax        float64 `json:"ioSizeMax"`
	IoSizeMin        float64 `json:"ioSizeMin"`
	QueueLengthAvg   float64 `json:"queueLengthAvg"`
	QueueLengthMax   float64 `json:"queueLengthMax"`
	QueueLengthMin   float64 `json:"queueLengthMin"`
	ServiceTimeMSAvg float64 `json:"serviceTimeMSAvg"`
	ServiceTimeMSMax float64 `json:"serviceTimeMSMax"`
	ServiceTimeMSMin float64 `json:"serviceTimeMSMin"`
}
