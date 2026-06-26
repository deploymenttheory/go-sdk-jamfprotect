package plan

// GraphQL fragments and queries for Plans

const planFields = `
fragment PlanFields on Plan {
	id
	hash
	name
	description
	created
	updated
	logLevel
	autoUpdate
	commsConfig {
		fqdn
		protocol
	}
	infoSync {
		attrs
		insightsSyncInterval
	}
	signaturesFeedConfig {
		mode
	}
	actionConfigs {
		id
		name
	}
	exceptionSets {
		uuid
		name
		managed
	}
	usbControlSet {
		id
		name
	}
	telemetry {
		id
		name
	}
	telemetryV2 {
		id
		name
	}
	analyticSets {
		type
		analyticSet {
			uuid
			name
			managed
			analytics {
				uuid
				categories
			}
		}
	}
}
`

const createPlanMutation = `
mutation createPlan(
	$name: String!,
	$description: String!,
	$logLevel: LOG_LEVEL_ENUM,
	$actionConfigs: ID!,
	$exceptionSets: [ID!],
	$telemetry: ID,
	$telemetryV2: ID,
	$analyticSets: [PlanAnalyticSetInput!],
	$usbControlSet: ID,
	$commsConfig: CommsConfigInput!,
	$infoSync: InfoSyncInput!,
	$autoUpdate: Boolean!,
	$signaturesFeedConfig: SignaturesFeedConfigInput!
) {
	createPlan(input: {
		name: $name,
		description: $description,
		logLevel: $logLevel,
		actionConfigs: $actionConfigs,
		exceptionSets: $exceptionSets,
		telemetry: $telemetry,
		telemetryV2: $telemetryV2,
		analyticSets: $analyticSets,
		usbControlSet: $usbControlSet,
		commsConfig: $commsConfig,
		infoSync: $infoSync,
		autoUpdate: $autoUpdate,
		signaturesFeedConfig: $signaturesFeedConfig
	}) {
		...PlanFields
	}
}
` + planFields

const getPlanQuery = `
query getPlan($id: ID!) {
	getPlan(id: $id) {
		...PlanFields
	}
}
` + planFields

const updatePlanMutation = `
mutation updatePlan(
	$id: ID!,
	$name: String!,
	$description: String!,
	$logLevel: LOG_LEVEL_ENUM,
	$actionConfigs: ID!,
	$exceptionSets: [ID!],
	$telemetry: ID,
	$telemetryV2: ID,
	$analyticSets: [PlanAnalyticSetInput!],
	$usbControlSet: ID,
	$commsConfig: CommsConfigInput!,
	$infoSync: InfoSyncInput!,
	$autoUpdate: Boolean!,
	$signaturesFeedConfig: SignaturesFeedConfigInput!
) {
	updatePlan(id: $id, input: {
		name: $name,
		description: $description,
		logLevel: $logLevel,
		actionConfigs: $actionConfigs,
		exceptionSets: $exceptionSets,
		telemetry: $telemetry,
		telemetryV2: $telemetryV2,
		analyticSets: $analyticSets,
		usbControlSet: $usbControlSet,
		commsConfig: $commsConfig,
		infoSync: $infoSync,
		autoUpdate: $autoUpdate,
		signaturesFeedConfig: $signaturesFeedConfig
	}) {
		...PlanFields
	}
}
` + planFields

const deletePlanMutation = `
mutation deletePlan($id: ID!) {
	deletePlan(id: $id) {
		id
	}
}
`

const listPlansQuery = `
query listPlans($nextToken: String, $direction: OrderDirection!, $field: PlanOrderField!) {
	listPlans(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...PlanFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + planFields

const listPlanNamesQuery = `
query listPlanNames($nextToken: String) {
	listPlanNames: listPlans(input: { next: $nextToken }) {
		items {
			name
		}
		pageInfo {
			next
			total
		}
	}
}
`

const getPlanConfigurationAndSetOptionsQuery = `
query getPlanConfigurationAndSetOptions(
	$RBAC_ActionConfigs: Boolean!
	$RBAC_Telemetry: Boolean!
	$RBAC_USBControlSet: Boolean!
	$RBAC_ExceptionSet: Boolean!
	$RBAC_AnalyticSet: Boolean!
) {
	actionConfigs: listActionConfigs(
		input: { order: { direction: DESC, field: created } }
	) @include(if: $RBAC_ActionConfigs) {
		items {
			name
			id
		}
	}
	telemetries: listTelemetries(
		input: { order: { direction: DESC, field: created } }
	) @include(if: $RBAC_Telemetry) {
		items {
			name
			id
		}
	}
	telemetriesV2: listTelemetriesV2(
		input: { order: { direction: DESC, field: created } }
	) @include(if: $RBAC_Telemetry) {
		items {
			name
			id
		}
	}
	usbControlSets: listUSBControlSets(
		input: { order: { direction: DESC, field: created } }
	) @include(if: $RBAC_USBControlSet) {
		items {
			name
			id
		}
	}
	exceptionSets: listExceptionSets(
		input: { order: { direction: DESC, field: created } }
	) @include(if: $RBAC_ExceptionSet) {
		items {
			name
			uuid
			managed
		}
	}
	analyticSets: listAnalyticSets(
		input: {
			order: { direction: DESC, field: created }
			filter: { managed: { equals: false } }
		}
	) @include(if: $RBAC_AnalyticSet) {
		items {
			name
			uuid
			managed
			types
		}
	}
		managedAnalyticSets: listAnalyticSets(
		input: {
			order: { direction: DESC, field: created }
			filter: { managed: { equals: true } }
		}
	) @include(if: $RBAC_AnalyticSet) {
		items {
			name
			description
			uuid
			managed
			types
		}
	}
}
`

const getPlansConfigProfileQuery = `
query getPlansConfigProfile($id: ID!, $input: ProfileOptionsInput) {
	getPlansConfigProfile(id: $id, input: $input)
}
`
