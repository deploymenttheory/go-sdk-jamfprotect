package computer

// GraphQL fragments and queries for Computer

const computerFields = `
fragment ComputerFields on Computer {
	uuid
	serial
	hostName
	modelName
	osMajor
	osMinor
	osPatch
	arch @skip(if: $isList)
	certid @skip(if: $isList)
	memorySize @skip(if: $isList)
	osString
	kernelVersion @skip(if: $isList)
	installType @skip(if: $isList)
	label
	created
	updated
	version
	checkin
	configHash
	tags
	signaturesVersion @include(if: $RBAC_ThreatPreventionVersion)
	plan @include(if: $RBAC_Plan) {
		id
		name
		hash
	}
	insightsStatsFail @include(if: $RBAC_Insight)
	insightsUpdated @include(if: $RBAC_Insight)
	connectionStatus
	lastConnection
	lastConnectionIp
	lastDisconnection
	lastDisconnectionReason
	webProtectionActive
	fullDiskAccess
	pendingPlan
}
`

const getComputerQuery = `
query getComputer(
	$uuid: ID!,
	$isList: Boolean!,
	$RBAC_ThreatPreventionVersion: Boolean!,
	$RBAC_Plan: Boolean!,
	$RBAC_Insight: Boolean!
) {
	getComputer(uuid: $uuid) {
		...ComputerFields
	}
}
` + computerFields

const listComputersQuery = `
query listComputers(
	$pageSize: Int,
	$direction: OrderDirection!,
	$field: [ComputerOrderField!],
	$isList: Boolean!,
	$RBAC_ThreatPreventionVersion: Boolean!,
	$RBAC_Plan: Boolean!,
	$RBAC_Insight: Boolean!
) {
	listComputers(
		input: {
			order: {direction: $direction, field: $field},
			pageSize: $pageSize
		}
	) {
		items {
			...ComputerFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + computerFields
