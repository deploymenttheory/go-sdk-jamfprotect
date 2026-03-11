package dataforwarding

// GraphQL queries and mutations for Data Forwarding

const dataForwardingFields = `
fragment DataForwardFields on Organization {
	uuid
	forward {
		s3 {
			bucket
			enabled
			encrypted
			prefix
			role
			cloudformation
		}
		sentinel {
			enabled
			customerId
			sharedKey
			logType
			domain
		}
		sentinelV2 {
			enabled
			secretExists
			azureTenantId
			azureClientId
			endpoint
			alerts {
				enabled
				dcrImmutableId
				streamName
			}
			ulogs {
				enabled
				dcrImmutableId
				streamName
			}
			telemetries {
				enabled
				dcrImmutableId
				streamName
			}
			telemetriesV2 {
				enabled
				dcrImmutableId
				streamName
			}
		}
	}
}
`

const getDataForwardingQuery = `
query getDataForwarding {
	getOrganization {
		...DataForwardFields
	}
}
` + dataForwardingFields

const updateDataForwardingMutation = `
mutation updateOrganizationForward($s3: OrganizationS3ForwardInput!, $sentinel: OrganizationSentinelForwardInput!, $sentinelV2: OrganizationSentinelV2ForwardInput!) {
	updateOrganizationForward(
		input: {s3: $s3, sentinel: $sentinel, sentinelV2: $sentinelV2}
	) {
		...DataForwardFields
	}
}
` + dataForwardingFields
