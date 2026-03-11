package changemanagement

// GraphQL queries and mutations for Change Management

const getConfigFreezeQuery = `
query getConfigFreeze {
	getAppInitializationData {
		configFreeze
	}
}
`

const updateConfigFreezeMutation = `
mutation updateOrganizationConfigFreeze($configFreeze: Boolean!) {
	updateOrganizationConfigFreeze(input: {configFreeze: $configFreeze}) {
		configFreeze
	}
}
`
