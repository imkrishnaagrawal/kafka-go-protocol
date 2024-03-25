package types

type ResponseErrorSet struct {
	Errors map[int]KafkaResponseError
}
type KafkaResponseError struct {
	Name        string
	Code        int
	Retirable   bool
	Description string
}

func NewResponseErrorSet() *ResponseErrorSet {
	responseError := ResponseErrorSet{}
	responseError.ConfigureErrorSet()
	return &responseError
}
func (responseError *ResponseErrorSet) Get(code int) KafkaResponseError {
	return responseError.Errors[code]
}

func (responseError *ResponseErrorSet) ConfigureErrorSet() {
	responseError.Errors = map[int]KafkaResponseError{
		-1: {
			Name:        "UNKNOWN_SERVER_ERROR",
			Code:        -1,
			Retirable:   false,
			Description: "The server experienced an unexpected error when processing the request.",
		},

		0: {
			Name:        "NONE",
			Code:        0,
			Retirable:   false,
			Description: "",
		},

		1: {
			Name:        "OFFSET_OUT_OF_RANGE",
			Code:        1,
			Retirable:   false,
			Description: "The requested offset is not within the range of offsets maintained by the server.",
		},

		2: {
			Name:        "CORRUPT_MESSAGE",
			Code:        2,
			Retirable:   true,
			Description: "This message has failed its CRC checksum, exceeds the valid size, has a null key for a compacted topic, or is otherwise corrupt.",
		},

		3: {
			Name:        "UNKNOWN_TOPIC_OR_PARTITION",
			Code:        3,
			Retirable:   true,
			Description: "This server does not host this topic-partition.",
		},

		4: {
			Name:        "INVALID_FETCH_SIZE",
			Code:        4,
			Retirable:   false,
			Description: "The requested fetch size is invalid.",
		},

		5: {
			Name:        "LEADER_NOT_AVAILABLE",
			Code:        5,
			Retirable:   true,
			Description: "There is no leader for this topic-partition as we are in the middle of a leadership election.",
		},

		6: {
			Name:        "NOT_LEADER_OR_FOLLOWER",
			Code:        6,
			Retirable:   true,
			Description: "For requests intended only for the leader, this error indicates that the broker is not the current leader. For requests intended for any replica, this error indicates that the broker is not a replica of the topic partition.",
		},

		7: {
			Name:        "REQUEST_TIMED_OUT",
			Code:        7,
			Retirable:   true,
			Description: "The request timed out.",
		},

		8: {
			Name:        "BROKER_NOT_AVAILABLE",
			Code:        8,
			Retirable:   false,
			Description: "The broker is not available.",
		},

		9: {
			Name:        "REPLICA_NOT_AVAILABLE",
			Code:        9,
			Retirable:   true,
			Description: "The replica is not available for the requested topic-partition. Produce/Fetch requests and other requests intended only for the leader or follower return NOT_LEADER_OR_FOLLOWER if the broker is not a replica of the topic-partition.",
		},

		10: {
			Name:        "MESSAGE_TOO_LARGE",
			Code:        10,
			Retirable:   false,
			Description: "The request included a message larger than the max message size the server will accept.",
		},

		11: {
			Name:        "STALE_CONTROLLER_EPOCH",
			Code:        11,
			Retirable:   false,
			Description: "The controller moved to another broker.",
		},

		12: {
			Name:        "OFFSET_METADATA_TOO_LARGE",
			Code:        12,
			Retirable:   false,
			Description: "The metadata field of the offset request was too large.",
		},

		13: {
			Name:        "NETWORK_EXCEPTION",
			Code:        13,
			Retirable:   true,
			Description: "The server disconnected before a response was received.",
		},

		14: {
			Name:        "COORDINATOR_LOAD_IN_PROGRESS",
			Code:        14,
			Retirable:   true,
			Description: "The coordinator is loading and hence can't process requests.",
		},

		15: {
			Name:        "COORDINATOR_NOT_AVAILABLE",
			Code:        15,
			Retirable:   true,
			Description: "The coordinator is not available.",
		},

		16: {
			Name:        "NOT_COORDINATOR",
			Code:        16,
			Retirable:   true,
			Description: "This is not the correct coordinator.",
		},

		17: {
			Name:        "INVALID_TOPIC_EXCEPTION",
			Code:        17,
			Retirable:   false,
			Description: "The request attempted to perform an operation on an invalid topic.",
		},

		18: {
			Name:        "RECORD_LIST_TOO_LARGE",
			Code:        18,
			Retirable:   false,
			Description: "The request included message batch larger than the configured segment size on the server.",
		},

		19: {
			Name:        "NOT_ENOUGH_REPLICAS",
			Code:        19,
			Retirable:   true,
			Description: "Messages are rejected since there are fewer in-sync replicas than required.",
		},

		20: {
			Name:        "NOT_ENOUGH_REPLICAS_AFTER_APPEND",
			Code:        20,
			Retirable:   true,
			Description: "Messages are written to the log, but to fewer in-sync replicas than required.",
		},

		21: {
			Name:        "INVALID_REQUIRED_ACKS",
			Code:        21,
			Retirable:   false,
			Description: "Produce request specified an invalid value for required acks.",
		},

		22: {
			Name:        "ILLEGAL_GENERATION",
			Code:        22,
			Retirable:   false,
			Description: "Specified group generation id is not valid.",
		},

		23: {
			Name:        "INCONSISTENT_GROUP_PROTOCOL",
			Code:        23,
			Retirable:   false,
			Description: "The group member's supported protocols are incompatible with those of existing members or first group member tried to join with empty protocol type or empty protocol list.",
		},

		24: {
			Name:        "INVALID_GROUP_ID",
			Code:        24,
			Retirable:   false,
			Description: "The configured groupId is invalid.",
		},

		25: {
			Name:        "UNKNOWN_MEMBER_ID",
			Code:        25,
			Retirable:   false,
			Description: "The coordinator is not aware of this member.",
		},

		26: {
			Name:        "INVALID_SESSION_TIMEOUT",
			Code:        26,
			Retirable:   false,
			Description: "The session timeout is not within the range allowed by the broker (as configured by group.min.session.timeout.ms and group.max.session.timeout.ms).",
		},

		27: {
			Name:        "REBALANCE_IN_PROGRESS",
			Code:        27,
			Retirable:   false,
			Description: "The group is rebalancing, so a rejoin is needed.",
		},

		28: {
			Name:        "INVALID_COMMIT_OFFSET_SIZE",
			Code:        28,
			Retirable:   false,
			Description: "The committing offset data size is not valid.",
		},

		29: {
			Name:        "TOPIC_AUTHORIZATION_FAILED",
			Code:        29,
			Retirable:   false,
			Description: "Topic authorization failed.",
		},

		30: {
			Name:        "GROUP_AUTHORIZATION_FAILED",
			Code:        30,
			Retirable:   false,
			Description: "Group authorization failed.",
		},

		31: {
			Name:        "CLUSTER_AUTHORIZATION_FAILED",
			Code:        31,
			Retirable:   false,
			Description: "Cluster authorization failed.",
		},

		32: {
			Name:        "INVALID_TIMESTAMP",
			Code:        32,
			Retirable:   false,
			Description: "The timestamp of the message is out of acceptable range.",
		},

		33: {
			Name:        "UNSUPPORTED_SASL_MECHANISM",
			Code:        33,
			Retirable:   false,
			Description: "The broker does not support the requested SASL mechanism.",
		},

		34: {
			Name:        "ILLEGAL_SASL_STATE",
			Code:        34,
			Retirable:   false,
			Description: "Request is not valid given the current SASL state.",
		},

		35: {
			Name:        "UNSUPPORTED_VERSION",
			Code:        35,
			Retirable:   false,
			Description: "The version of API is not supported.",
		},

		36: {
			Name:        "TOPIC_ALREADY_EXISTS",
			Code:        36,
			Retirable:   false,
			Description: "Topic with this name already exists.",
		},

		37: {
			Name:        "INVALID_PARTITIONS",
			Code:        37,
			Retirable:   false,
			Description: "Number of partitions is below 1.",
		},

		38: {
			Name:        "INVALID_REPLICATION_FACTOR",
			Code:        38,
			Retirable:   false,
			Description: "Replication factor is below 1 or larger than the number of available brokers.",
		},

		39: {
			Name:        "INVALID_REPLICA_ASSIGNMENT",
			Code:        39,
			Retirable:   false,
			Description: "Replica assignment is invalid.",
		},

		40: {
			Name:        "INVALID_CONFIG",
			Code:        40,
			Retirable:   false,
			Description: "Configuration is invalid.",
		},

		41: {
			Name:        "NOT_CONTROLLER",
			Code:        41,
			Retirable:   true,
			Description: "This is not the correct controller for this cluster.",
		},

		42: {
			Name:        "INVALID_REQUEST",
			Code:        42,
			Retirable:   false,
			Description: "This most likely occurs because of a request being malformed by the client library or the message was sent to an incompatible broker. See the broker logs for more details.",
		},

		43: {
			Name:        "UNSUPPORTED_FOR_MESSAGE_FORMAT",
			Code:        43,
			Retirable:   false,
			Description: "The message format version on the broker does not support the request.",
		},

		44: {
			Name:        "POLICY_VIOLATION",
			Code:        44,
			Retirable:   false,
			Description: "Request parameters do not satisfy the configured policy.",
		},

		45: {
			Name:        "OUT_OF_ORDER_SEQUENCE_NUMBER",
			Code:        45,
			Retirable:   false,
			Description: "The broker received an out of order sequence number.",
		},

		46: {
			Name:        "DUPLICATE_SEQUENCE_NUMBER",
			Code:        46,
			Retirable:   false,
			Description: "The broker received a duplicate sequence number.",
		},

		47: {
			Name:        "INVALID_PRODUCER_EPOCH",
			Code:        47,
			Retirable:   false,
			Description: "Producer attempted to produce with an old epoch.",
		},

		48: {
			Name:        "INVALID_TXN_STATE",
			Code:        48,
			Retirable:   false,
			Description: "The producer attempted a transactional operation in an invalid state.",
		},

		49: {
			Name:        "INVALID_PRODUCER_ID_MAPPING",
			Code:        49,
			Retirable:   false,
			Description: "The producer attempted to use a producer id which is not currently assigned to its transactional id.",
		},

		50: {
			Name:        "INVALID_TRANSACTION_TIMEOUT",
			Code:        50,
			Retirable:   false,
			Description: "The transaction timeout is larger than the maximum value allowed by the broker (as configured by transaction.max.timeout.ms).",
		},

		51: {
			Name:        "CONCURRENT_TRANSACTIONS",
			Code:        51,
			Retirable:   true,
			Description: "The producer attempted to update a transaction while another concurrent operation on the same transaction was ongoing.",
		},

		52: {
			Name:        "TRANSACTION_COORDINATOR_FENCED",
			Code:        52,
			Retirable:   false,
			Description: "Indicates that the transaction coordinator sending a WriteTxnMarker is no longer the current coordinator for a given producer.",
		},

		53: {
			Name:        "TRANSACTIONAL_ID_AUTHORIZATION_FAILED",
			Code:        53,
			Retirable:   false,
			Description: "Transactional Id authorization failed.",
		},

		54: {
			Name:        "SECURITY_DISABLED",
			Code:        54,
			Retirable:   false,
			Description: "Security features are disabled.",
		},

		55: {
			Name:        "OPERATION_NOT_ATTEMPTED",
			Code:        55,
			Retirable:   false,
			Description: "The broker did not attempt to execute this operation. This may happen for batched RPCs where some operations in the batch failed, causing the broker to respond without trying the rest.",
		},

		56: {
			Name:        "KAFKA_STORAGE_ERROR",
			Code:        56,
			Retirable:   true,
			Description: "Disk error when trying to access log file on the disk.",
		},

		57: {
			Name:        "LOG_DIR_NOT_FOUND",
			Code:        57,
			Retirable:   false,
			Description: "The user-specified log directory is not found in the broker config.",
		},

		58: {
			Name:        "SASL_AUTHENTICATION_FAILED",
			Code:        58,
			Retirable:   false,
			Description: "SASL Authentication failed.",
		},

		59: {
			Name:        "UNKNOWN_PRODUCER_ID",
			Code:        59,
			Retirable:   false,
			Description: "This exception is raised by the broker if it could not locate the producer metadata associated with the producerId in question. This could happen if, for instance, the producer's records were deleted because their retention time had elapsed. Once the last records of the producerId are removed, the producer's metadata is removed from the broker, and future appends by the producer will return this exception.",
		},

		60: {
			Name:        "REASSIGNMENT_IN_PROGRESS",
			Code:        60,
			Retirable:   false,
			Description: "A partition reassignment is in progress.",
		},

		61: {
			Name:        "DELEGATION_TOKEN_AUTH_DISABLED",
			Code:        61,
			Retirable:   false,
			Description: "Delegation Token feature is not enabled.",
		},

		62: {
			Name:        "DELEGATION_TOKEN_NOT_FOUND",
			Code:        62,
			Retirable:   false,
			Description: "Delegation Token is not found on server.",
		},

		63: {
			Name:        "DELEGATION_TOKEN_OWNER_MISMATCH",
			Code:        63,
			Retirable:   false,
			Description: "Specified Principal is not valid Owner/Renewer.",
		},

		64: {
			Name:        "DELEGATION_TOKEN_REQUEST_NOT_ALLOWED",
			Code:        64,
			Retirable:   false,
			Description: "Delegation Token requests are not allowed on PLAINTEXT/1-way SSL channels and on delegation token authenticated channels.",
		},

		65: {
			Name:        "DELEGATION_TOKEN_AUTHORIZATION_FAILED",
			Code:        65,
			Retirable:   false,
			Description: "Delegation Token authorization failed.",
		},

		66: {
			Name:        "DELEGATION_TOKEN_EXPIRED",
			Code:        66,
			Retirable:   false,
			Description: "Delegation Token is expired.",
		},

		67: {
			Name:        "INVALID_PRINCIPAL_TYPE",
			Code:        67,
			Retirable:   false,
			Description: "Supplied principalType is not supported.",
		},

		68: {
			Name:        "NON_EMPTY_GROUP",
			Code:        68,
			Retirable:   false,
			Description: "The group is not empty.",
		},

		69: {
			Name:        "GROUP_ID_NOT_FOUND",
			Code:        69,
			Retirable:   false,
			Description: "The group id does not exist.",
		},

		70: {
			Name:        "FETCH_SESSION_ID_NOT_FOUND",
			Code:        70,
			Retirable:   true,
			Description: "The fetch session ID was not found.",
		},

		71: {
			Name:        "INVALID_FETCH_SESSION_EPOCH",
			Code:        71,
			Retirable:   true,
			Description: "The fetch session epoch is invalid.",
		},

		72: {
			Name:        "LISTENER_NOT_FOUND",
			Code:        72,
			Retirable:   true,
			Description: "There is no listener on the leader broker that matches the listener on which metadata request was processed.",
		},

		73: {
			Name:        "TOPIC_DELETION_DISABLED",
			Code:        73,
			Retirable:   false,
			Description: "Topic deletion is disabled.",
		},

		74: {
			Name:        "FENCED_LEADER_EPOCH",
			Code:        74,
			Retirable:   true,
			Description: "The leader epoch in the request is older than the epoch on the broker.",
		},

		75: {
			Name:        "UNKNOWN_LEADER_EPOCH",
			Code:        75,
			Retirable:   true,
			Description: "The leader epoch in the request is newer than the epoch on the broker.",
		},

		76: {
			Name:        "UNSUPPORTED_COMPRESSION_TYPE",
			Code:        76,
			Retirable:   false,
			Description: "The requesting client does not support the compression type of given partition.",
		},

		77: {
			Name:        "STALE_BROKER_EPOCH",
			Code:        77,
			Retirable:   false,
			Description: "Broker epoch has changed.",
		},

		78: {
			Name:        "OFFSET_NOT_AVAILABLE",
			Code:        78,
			Retirable:   true,
			Description: "The leader high watermark has not caught up from a recent leader election so the offsets cannot be guaranteed to be monotonically increasing.",
		},

		79: {
			Name:        "MEMBER_ID_REQUIRED",
			Code:        79,
			Retirable:   false,
			Description: "The group member needs to have a valid member id before actually entering a consumer group.",
		},

		80: {
			Name:        "PREFERRED_LEADER_NOT_AVAILABLE",
			Code:        80,
			Retirable:   true,
			Description: "The preferred leader was not available.",
		},

		81: {
			Name:        "GROUP_MAX_SIZE_REACHED",
			Code:        81,
			Retirable:   false,
			Description: "The consumer group has reached its max size.",
		},

		82: {
			Name:        "FENCED_INSTANCE_ID",
			Code:        82,
			Retirable:   false,
			Description: "The broker rejected this static consumer since another consumer with the same group.instance.id has registered with a different member.id.",
		},

		83: {
			Name:        "ELIGIBLE_LEADERS_NOT_AVAILABLE",
			Code:        83,
			Retirable:   true,
			Description: "Eligible topic partition leaders are not available.",
		},

		84: {
			Name:        "ELECTION_NOT_NEEDED",
			Code:        84,
			Retirable:   true,
			Description: "Leader election not needed for topic partition.",
		},

		85: {
			Name:        "NO_REASSIGNMENT_IN_PROGRESS",
			Code:        85,
			Retirable:   false,
			Description: "No partition reassignment is in progress.",
		},

		86: {
			Name:        "GROUP_SUBSCRIBED_TO_TOPIC",
			Code:        86,
			Retirable:   false,
			Description: "Deleting offsets of a topic is forbidden while the consumer group is actively subscribed to it.",
		},

		87: {
			Name:        "INVALID_RECORD",
			Code:        87,
			Retirable:   false,
			Description: "This record has failed the validation on broker and hence will be rejected.",
		},

		88: {
			Name:        "UNSTABLE_OFFSET_COMMIT",
			Code:        88,
			Retirable:   true,
			Description: "There are unstable offsets that need to be cleared.",
		},

		89: {
			Name:        "THROTTLING_QUOTA_EXCEEDED",
			Code:        89,
			Retirable:   true,
			Description: "The throttling quota has been exceeded.",
		},

		90: {
			Name:        "PRODUCER_FENCED",
			Code:        90,
			Retirable:   false,
			Description: "There is a newer producer with the same transactionalId which fences the current one.",
		},

		91: {
			Name:        "RESOURCE_NOT_FOUND",
			Code:        91,
			Retirable:   false,
			Description: "A request illegally referred to a resource that does not exist.",
		},

		92: {
			Name:        "DUPLICATE_RESOURCE",
			Code:        92,
			Retirable:   false,
			Description: "A request illegally referred to the same resource twice.",
		},

		93: {
			Name:        "UNACCEPTABLE_CREDENTIAL",
			Code:        93,
			Retirable:   false,
			Description: "Requested credential would not meet criteria for acceptability.",
		},

		94: {
			Name:        "INCONSISTENT_VOTER_SET",
			Code:        94,
			Retirable:   false,
			Description: "Indicates that the either the sender or recipient of a voter-only request is not one of the expected voters",
		},

		95: {
			Name:        "INVALID_UPDATE_VERSION",
			Code:        95,
			Retirable:   false,
			Description: "The given update version was invalid.",
		},

		96: {
			Name:        "FEATURE_UPDATE_FAILED",
			Code:        96,
			Retirable:   false,
			Description: "Unable to update finalized features due to an unexpected server error.",
		},

		97: {
			Name:        "PRINCIPAL_DESERIALIZATION_FAILURE",
			Code:        97,
			Retirable:   false,
			Description: "Request principal deserialization failed during forwarding. This indicates an internal error on the broker cluster security setup.",
		},

		98: {
			Name:        "SNAPSHOT_NOT_FOUND",
			Code:        98,
			Retirable:   false,
			Description: "Requested snapshot was not found",
		},

		99: {
			Name:        "POSITION_OUT_OF_RANGE",
			Code:        99,
			Retirable:   false,
			Description: "Requested position is not greater than or equal to zero, and less than the size of the snapshot.",
		},

		100: {
			Name:        "UNKNOWN_TOPIC_ID",
			Code:        100,
			Retirable:   true,
			Description: "This server does not host this topic ID.",
		},

		101: {
			Name:        "DUPLICATE_BROKER_REGISTRATION",
			Code:        101,
			Retirable:   false,
			Description: "This broker ID is already in use.",
		},

		102: {
			Name:        "BROKER_ID_NOT_REGISTERED",
			Code:        102,
			Retirable:   false,
			Description: "The given broker ID was not registered.",
		},

		103: {
			Name:        "INCONSISTENT_TOPIC_ID",
			Code:        103,
			Retirable:   true,
			Description: "The log's topic ID did not match the topic ID in the request",
		},

		104: {
			Name:        "INCONSISTENT_CLUSTER_ID",
			Code:        104,
			Retirable:   false,
			Description: "The clusterId in the request does not match that found on the server",
		},

		105: {
			Name:        "TRANSACTIONAL_ID_NOT_FOUND",
			Code:        105,
			Retirable:   false,
			Description: "The transactionalId could not be found",
		},

		106: {
			Name:        "FETCH_SESSION_TOPIC_ID_ERROR",
			Code:        106,
			Retirable:   true,
			Description: "The fetch session encountered inconsistent topic ID usage",
		},

		107: {
			Name:        "INELIGIBLE_REPLICA",
			Code:        107,
			Retirable:   false,
			Description: "The new ISR contains at least one ineligible replica.",
		},

		108: {
			Name:        "NEW_LEADER_ELECTED",
			Code:        108,
			Retirable:   false,
			Description: "The AlterPartition request successfully updated the partition state but the leader has changed.",
		},

		109: {
			Name:        "OFFSET_MOVED_TO_TIERED_STORAGE",
			Code:        109,
			Retirable:   false,
			Description: "The requested offset is moved to tiered storage.",
		},

		110: {
			Name:        "FENCED_MEMBER_EPOCH",
			Code:        110,
			Retirable:   false,
			Description: "The member epoch is fenced by the group coordinator. The member must abandon all its partitions and rejoin.",
		},

		111: {
			Name:        "UNRELEASED_INSTANCE_ID",
			Code:        111,
			Retirable:   false,
			Description: "The instance ID is still used by another member in the consumer group. That member must leave first.",
		},

		112: {
			Name:        "UNSUPPORTED_ASSIGNOR",
			Code:        112,
			Retirable:   false,
			Description: "The assignor or its version range is not supported by the consumer group.",
		},

		113: {
			Name:        "STALE_MEMBER_EPOCH",
			Code:        113,
			Retirable:   false,
			Description: "The member epoch is stale. The member must retry after receiving its updated member epoch via the ConsumerGroupHeartbeat API.",
		},

		114: {
			Name:        "MISMATCHED_ENDPOINT_TYPE",
			Code:        114,
			Retirable:   false,
			Description: "The request was sent to an endpoint of the wrong type.",
		},

		115: {
			Name:        "UNSUPPORTED_ENDPOINT_TYPE",
			Code:        115,
			Retirable:   false,
			Description: "This endpoint type is not supported yet.",
		},

		116: {
			Name:        "UNKNOWN_CONTROLLER_ID",
			Code:        116,
			Retirable:   false,
			Description: "This controller ID is not known.",
		},

		117: {
			Name:        "UNKNOWN_SUBSCRIPTION_ID",
			Code:        117,
			Retirable:   false,
			Description: "Client sent a push telemetry request with an invalid or outdated subscription ID.",
		},

		118: {
			Name:        "TELEMETRY_TOO_LARGE",
			Code:        118,
			Retirable:   false,
			Description: "Client sent a push telemetry request larger than the maximum size the broker will accept.",
		},

		119: {
			Name:        "INVALID_REGISTRATION",
			Code:        119,
			Retirable:   false,
			Description: "The controller has considered the broker registration to be invalid.",
		},
	}
}
