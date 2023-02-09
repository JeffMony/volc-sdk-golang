// Code Generated by gadget/xsdk, DO NOT EDIT

package privatezone

type ListRegionsRequest struct {
	Tag *string `form:"Tag" json:"Tag,omitempty"`
}

type ListRegionsResponse struct {
	Regions []Region `form:"Regions" json:"Regions,omitempty"`
}

type BindVPCRequest struct {
	VPCs map[string][]string `form:"VPCs" json:"VPCs,omitempty"`
	ZID  *int64              `form:"ZID" json:"ZID,omitempty"`
}

type CreatePrivateZoneRequest struct {
	LineMode      *int64              `form:"LineMode" json:"LineMode,omitempty"`
	RecursionMode *bool               `form:"RecursionMode" json:"RecursionMode,omitempty"`
	Remark        *string             `form:"Remark" json:"Remark,omitempty"`
	VPCs          map[string][]string `form:"VPCs" json:"VPCs,omitempty"`
	ZoneName      *string             `form:"ZoneName" json:"ZoneName,omitempty"`
}

type CreatePrivateZoneResponse struct {
	ZID *int64 `form:"ZID" json:"ZID,omitempty"`
}

type DeletePrivateZoneRequest struct {
	DeleteWhenEmpty *bool  `form:"DeleteWhenEmpty" json:"DeleteWhenEmpty,omitempty"`
	ZID             *int64 `form:"ZID" json:"ZID,omitempty"`
}

type IncBindVPCRequest struct {
	Bind   map[string][]string `form:"Bind" json:"Bind,omitempty"`
	Unbind map[string][]string `form:"Unbind" json:"Unbind,omitempty"`
	ZID    *int64              `form:"ZID" json:"ZID,omitempty"`
}

type ListBindVPCRequest struct {
	Region *string `form:"-" json:"-"`
	ZID    *string `form:"-" json:"-"`
}

type ListBindVPCResponse struct {
	BindVPCs []VPC  `form:"BindVPCs" json:"BindVPCs,omitempty"`
	Total    *int64 `form:"Total" json:"Total,omitempty"`
}

type ListPrivateZonesRequest struct {
	KeyWord       *string `form:"-" json:"-"`
	LineMode      *string `form:"-" json:"-"`
	PageNumber    *string `form:"-" json:"-"`
	PageSize      *string `form:"-" json:"-"`
	RecursionMode *string `form:"-" json:"-"`
	Region        *string `form:"-" json:"-"`
	SearchMode    *string `form:"-" json:"-"`
	VpcID         *string `form:"-" json:"-"`
}

type ListPrivateZonesResponse struct {
	Total *int64                   `form:"Total" json:"Total,omitempty"`
	Zones []TopPrivateZoneResponse `form:"Zones" json:"Zones,omitempty"`
}

type QueryPrivateZoneRequest struct {
	ZID *string `form:"-" json:"-"`
}

type QueryPrivateZoneResponse struct {
	BindVPCs      []VPC    `form:"BindVPCs" json:"BindVPCs,omitempty"`
	CreatedAt     *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	LastOperator  *string  `form:"LastOperator" json:"LastOperator,omitempty"`
	LineMode      *int64   `form:"LineMode" json:"LineMode,omitempty"`
	RecordCount   *int64   `form:"RecordCount" json:"RecordCount,omitempty"`
	RecursionMode *bool    `form:"RecursionMode" json:"RecursionMode,omitempty"`
	Region        []string `form:"Region" json:"Region,omitempty"`
	Remark        *string  `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt     *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZoneName      *string  `form:"ZoneName" json:"ZoneName,omitempty"`
}

type UpdatePrivateZoneRequest struct {
	LineMode      *int64  `form:"LineMode" json:"LineMode,omitempty"`
	RecursionMode *bool   `form:"RecursionMode" json:"RecursionMode,omitempty"`
	Remark        *string `form:"Remark" json:"Remark,omitempty"`
	ZID           *int64  `form:"ZID" json:"ZID,omitempty"`
}

type UpdatePrivateZoneResponse struct {
	CreatedAt     *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	LastOperator  *string  `form:"LastOperator" json:"LastOperator,omitempty"`
	LineMode      *int64   `form:"LineMode" json:"LineMode,omitempty"`
	RecordCount   *int64   `form:"RecordCount" json:"RecordCount,omitempty"`
	RecursionMode *bool    `form:"RecursionMode" json:"RecursionMode,omitempty"`
	Region        []string `form:"Region" json:"Region,omitempty"`
	Remark        *string  `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt     *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZID           *int64   `form:"ZID" json:"ZID,omitempty"`
	ZoneName      *string  `form:"ZoneName" json:"ZoneName,omitempty"`
}

type BatchCreateRecordRequest struct {
	Records []CRecord `form:"Records" json:"Records,omitempty"`
	ZID     *int64    `form:"ZID" json:"ZID,omitempty"`
}

type BatchCreateRecordResponse struct {
	RecordIDs []string `form:"RecordIDs" json:"RecordIDs,omitempty"`
}

type BatchDeleteRecordRequest struct {
	RecordIDs []string `form:"RecordIDs" json:"RecordIDs,omitempty"`
	ZID       *int64   `form:"ZID" json:"ZID,omitempty"`
}

type BatchUpdateRecordRequest struct {
	Records []URecord `form:"Records" json:"Records,omitempty"`
	ZID     *int64    `form:"ZID" json:"ZID,omitempty"`
}

type CreateRecordRequest struct {
	Host   *string `form:"Host" json:"Host,omitempty"`
	Line   *string `form:"Line" json:"Line,omitempty"`
	Remark *string `form:"Remark" json:"Remark,omitempty"`
	TTL    *int64  `form:"TTL" json:"TTL,omitempty"`
	Type   *string `form:"Type" json:"Type,omitempty"`
	Value  *string `form:"Value" json:"Value,omitempty"`
	Weight *int64  `form:"Weight" json:"Weight,omitempty"`
	ZID    *int64  `form:"ZID" json:"ZID,omitempty"`
}

type CreateRecordResponse struct {
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
}

type DeleteRecordRequest struct {
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
}

type ListRecordAttributesRequest struct {
	ZID *string `form:"-" json:"-"`
}

type ListRecordAttributesResponse struct {
	TTLs        []int64  `form:"TTLs" json:"TTLs,omitempty"`
	Types       []string `form:"Types" json:"Types,omitempty"`
	WeightLimit *int64   `form:"WeightLimit" json:"WeightLimit,omitempty"`
}

type ListRecordSetsRequest struct {
	Host        *string `form:"-" json:"-"`
	PageNumber  *string `form:"-" json:"-"`
	PageSize    *string `form:"-" json:"-"`
	RecordSetID *string `form:"-" json:"-"`
	SearchMode  *string `form:"-" json:"-"`
	ZID         *string `form:"-" json:"-"`
}

type ListRecordSetsResponse struct {
	PageNumber *int64             `form:"PageNumber" json:"PageNumber,omitempty"`
	PageSize   *int64             `form:"PageSize" json:"PageSize,omitempty"`
	RecordSets []TopRecordSetResp `form:"RecordSets" json:"RecordSets,omitempty"`
	TotalCount *int64             `form:"TotalCount" json:"TotalCount,omitempty"`
}

type ListRecordsRequest struct {
	Host         *string `form:"-" json:"-"`
	LastOperator *string `form:"-" json:"-"`
	Line         *string `form:"-" json:"-"`
	Name         *string `form:"-" json:"-"`
	PageNumber   *string `form:"-" json:"-"`
	PageSize     *string `form:"-" json:"-"`
	SearchMode   *string `form:"-" json:"-"`
	SearchOrder  *string `form:"-" json:"-"`
	Type         *string `form:"-" json:"-"`
	Value        *string `form:"-" json:"-"`
	ZID          *string `form:"-" json:"-"`
}

type ListRecordsResponse struct {
	Records    []Record `form:"Records" json:"Records,omitempty"`
	PageNumber *int64   `form:"page_number" json:"page_number,omitempty"`
	PageSize   *int64   `form:"page_size" json:"page_size,omitempty"`
	Total      *int64   `form:"total" json:"total,omitempty"`
}

type UpdateRecordRequest struct {
	Enable   *bool   `form:"Enable" json:"Enable,omitempty"`
	Host     *string `form:"Host" json:"Host,omitempty"`
	Line     *string `form:"Line" json:"Line,omitempty"`
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
	Remark   *string `form:"Remark" json:"Remark,omitempty"`
	TTL      *int64  `form:"TTL" json:"TTL,omitempty"`
	Type     *string `form:"Type" json:"Type,omitempty"`
	Value    *string `form:"Value" json:"Value,omitempty"`
	Weight   *int64  `form:"Weight" json:"Weight,omitempty"`
}

type UpdateRecordSetRequest struct {
	RecordSetID   *string `form:"RecordSetID" json:"RecordSetID,omitempty"`
	WeightEnabled *bool   `form:"WeightEnabled" json:"WeightEnabled,omitempty"`
}

type UpdateRecordSetResponse struct {
	FQDN          *string `form:"FQDN" json:"FQDN,omitempty"`
	Host          *string `form:"Host" json:"Host,omitempty"`
	ID            *string `form:"ID" json:"ID,omitempty"`
	Line          *string `form:"Line" json:"Line,omitempty"`
	Type          *string `form:"Type" json:"Type,omitempty"`
	WeightEnabled *bool   `form:"WeightEnabled" json:"WeightEnabled,omitempty"`
}

type Region struct {
	Code *string `form:"Code" json:"Code,omitempty"`
	Name *string `form:"Name" json:"Name,omitempty"`
}

type VPC struct {
	AccountID  *string `form:"AccountID" json:"AccountID,omitempty"`
	ID         *string `form:"ID" json:"ID,omitempty"`
	Region     *string `form:"Region" json:"Region,omitempty"`
	RegionName *string `form:"RegionName" json:"RegionName,omitempty"`
}

type TopPrivateZoneResponse struct {
	CreatedAt     *string  `form:"CreatedAt" json:"CreatedAt,omitempty"`
	LastOperator  *string  `form:"LastOperator" json:"LastOperator,omitempty"`
	LineMode      *int64   `form:"LineMode" json:"LineMode,omitempty"`
	RecordCount   *int64   `form:"RecordCount" json:"RecordCount,omitempty"`
	RecursionMode *bool    `form:"RecursionMode" json:"RecursionMode,omitempty"`
	Region        []string `form:"Region" json:"Region,omitempty"`
	Remark        *string  `form:"Remark" json:"Remark,omitempty"`
	UpdatedAt     *string  `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	ZID           *int64   `form:"ZID" json:"ZID,omitempty"`
	ZoneName      *string  `form:"ZoneName" json:"ZoneName,omitempty"`
}

type CRecord struct {
	Host   *string `form:"Host" json:"Host,omitempty"`
	Line   *string `form:"Line" json:"Line,omitempty"`
	Remark *string `form:"Remark" json:"Remark,omitempty"`
	TTL    *int64  `form:"TTL" json:"TTL,omitempty"`
	Type   *string `form:"Type" json:"Type,omitempty"`
	Value  *string `form:"Value" json:"Value,omitempty"`
	Weight *int64  `form:"Weight" json:"Weight,omitempty"`
}

type URecord struct {
	Enable   *bool   `form:"Enable" json:"Enable,omitempty"`
	Host     *string `form:"Host" json:"Host,omitempty"`
	Line     *string `form:"Line" json:"Line,omitempty"`
	RecordID *string `form:"RecordID" json:"RecordID,omitempty"`
	Remark   *string `form:"Remark" json:"Remark,omitempty"`
	TTL      *int64  `form:"TTL" json:"TTL,omitempty"`
	Type     *string `form:"Type" json:"Type,omitempty"`
	Value    *string `form:"Value" json:"Value,omitempty"`
	Weight   *int64  `form:"Weight" json:"Weight,omitempty"`
}

type TopRecordSetResp struct {
	FQDN          *string `form:"FQDN" json:"FQDN,omitempty"`
	Host          *string `form:"Host" json:"Host,omitempty"`
	ID            *string `form:"ID" json:"ID,omitempty"`
	Line          *string `form:"Line" json:"Line,omitempty"`
	Type          *string `form:"Type" json:"Type,omitempty"`
	WeightEnabled *bool   `form:"WeightEnabled" json:"WeightEnabled,omitempty"`
}

type Record struct {
	CreatedAt    *string `form:"CreatedAt" json:"CreatedAt,omitempty"`
	Host         *string `form:"Host" json:"Host,omitempty"`
	LastOperator *string `form:"LastOperator" json:"LastOperator,omitempty"`
	Line         *string `form:"Line" json:"Line,omitempty"`
	RecordID     *string `form:"RecordID" json:"RecordID,omitempty"`
	Remark       *string `form:"Remark" json:"Remark,omitempty"`
	TTL          *int64  `form:"TTL" json:"TTL,omitempty"`
	Type         *string `form:"Type" json:"Type,omitempty"`
	UpdatedAt    *string `form:"UpdatedAt" json:"UpdatedAt,omitempty"`
	Value        *string `form:"Value" json:"Value,omitempty"`
	Weight       *int64  `form:"Weight" json:"Weight,omitempty"`
	ZID          *int64  `form:"ZID" json:"ZID,omitempty"`
}