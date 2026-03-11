package apiclient

// GraphQL fragments and queries for ApiClient

const apiClientFields = `
fragment ApiClientFields on ApiClient {
	clientId
	created
	name
	assignedRoles {
		id
		name
	}
	password
}
`

const createApiClientMutation = `
mutation createApiClient($name: String!, $roleIds: [ID]) {
	createApiClient(input: {name: $name, roleIds: $roleIds}) {
		...ApiClientFields
	}
}
` + apiClientFields

const getApiClientQuery = `
query getApiClient($clientId: ID!) {
	getApiClient(clientId: $clientId) {
		...ApiClientFields
	}
}
` + apiClientFields

const updateApiClientMutation = `
mutation updateApiClient($clientId: ID!, $name: String!, $roleIds: [ID]) {
	updateApiClient(clientId: $clientId, input: {name: $name, roleIds: $roleIds}) {
		...ApiClientFields
	}
}
` + apiClientFields

const deleteApiClientMutation = `
mutation deleteApiClient($clientId: ID!) {
	deleteApiClient(clientId: $clientId) {
		clientId
	}
}
`

const listApiClientsQuery = `
query listApiClients($direction: OrderDirection!, $field: ApiClientOrderField!) {
	listApiClients(
		input: {order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...ApiClientFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + apiClientFields
