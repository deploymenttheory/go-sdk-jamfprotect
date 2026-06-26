package analytic

// GraphQL fragments and queries for Analytics

const analyticFields = `
fragment AnalyticFields on Analytic {
	uuid
	name
	label
	inputType
	filter
	description
	longDescription
	created
	updated
	actions
	analyticActions {
		name
		parameters
	}
	tenantActions {
		name
		parameters
	}
	tags
	level
	severity
	tenantSeverity
	snapshotFiles
	context {
		name
		type
		exprs
	}
	categories
	jamf
	remediation
}
`

const createAnalyticMutation = `
mutation createAnalytic(
	$name: String!,
	$inputType: String!,
	$description: String!,
	$actions: [String],
	$analyticActions: [AnalyticActionsInput]!,
	$tags: [String]!,
	$categories: [String]!,
	$filter: String!,
	$context: [AnalyticContextInput]!,
	$level: Int!,
	$severity: SEVERITY!,
	$snapshotFiles: [String]!
) {
	createAnalytic(input: {
		name: $name,
		inputType: $inputType,
		description: $description,
		actions: $actions,
		analyticActions: $analyticActions,
		tags: $tags,
		categories: $categories,
		filter: $filter,
		context: $context,
		level: $level,
		severity: $severity,
		snapshotFiles: $snapshotFiles
	}) {
		...AnalyticFields
	}
}
` + analyticFields

const getAnalyticQuery = `
query getAnalytic($uuid: ID!) {
	getAnalytic(uuid: $uuid) {
		...AnalyticFields
	}
}
` + analyticFields

const updateAnalyticMutation = `
mutation updateAnalytic(
	$uuid: ID!,
	$name: String!,
	$inputType: String!,
	$description: String!,
	$actions: [String],
	$analyticActions: [AnalyticActionsInput]!,
	$tags: [String]!,
	$categories: [String]!,
	$filter: String!,
	$context: [AnalyticContextInput]!,
	$level: Int!,
	$severity: SEVERITY,
	$snapshotFiles: [String]!
) {
	updateAnalytic(uuid: $uuid, input: {
		name: $name,
		inputType: $inputType,
		description: $description,
		actions: $actions,
		analyticActions: $analyticActions,
		categories: $categories,
		tags: $tags,
		filter: $filter,
		context: $context,
		level: $level,
		severity: $severity,
		snapshotFiles: $snapshotFiles
	}) {
		...AnalyticFields
	}
}
` + analyticFields

const updateInternalAnalyticMutation = `
mutation updateInternalAnalytic(
	$uuid: ID!,
	$tenantActions: [AnalyticActionsInput],
	$tenantSeverity: SEVERITY
) {
	updateInternalAnalytic(
		uuid: $uuid,
		input: {tenantActions: $tenantActions, tenantSeverity: $tenantSeverity}
	) {
		...AnalyticFields
	}
}
` + analyticFields

const deleteAnalyticMutation = `
mutation deleteAnalytic($uuid: ID!) {
	deleteAnalytic(uuid: $uuid) {
		uuid
	}
}
`

const listAnalyticsQuery = `
query listAnalytics {
	listAnalytics {
		items {
			...AnalyticFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + analyticFields

const listAnalyticsLiteQuery = `
query listAnalyticsLite {
	listAnalytics {
		items {
			name
			label
			uuid
			longDescription
			description
			tags
			inputType
			remediation
		}
		pageInfo {
			next
			total
		}
	}
}
`

const listAnalyticsNamesQuery = `
query listAnalyticsNames {
	listAnalyticsNames: listAnalytics {
		items {
			name
		}
	}
}
`

const listAnalyticsCategoriesQuery = `
query listAnalyticsCategories {
	listAnalyticsCategories {
		value
		count
	}
}
`

const listAnalyticsTagsQuery = `
query listAnalyticsTags {
	listAnalyticsTags {
		value
		count
	}
}
`

const listAnalyticsFilterOptionsQuery = `
query listAnalyticsFilterOptions {
	listAnalyticsTags {
		value
		count
	}
	listAnalyticsCategories {
		value
		count
	}
}
`
