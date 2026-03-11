package dataretention

// GraphQL queries and mutations for Data Retention

const getDataRetentionQuery = `
query getDataRetention {
	getOrganization {
		retention {
			database {
				log {
					numberOfDays
				}
				alert {
					numberOfDays
				}
			}
			cold {
				alert {
					numberOfDays
				}
			}
			updated
		}
	}
}
`

const updateDataRetentionMutation = `
mutation updateOrganizationRetention($databaseLogDays: Int!, $databaseAlertDays: Int!, $coldAlertDays: Int!) {
	updateOrganizationRetention(
		input: {
			retention: {
				database: {
					log: {numberOfDays: $databaseLogDays}
					alert: {numberOfDays: $databaseAlertDays}
				}
				cold: {alert: {numberOfDays: $coldAlertDays}}
			}
		}
	) {
		retention {
			database {
				log { numberOfDays }
				alert { numberOfDays }
			}
			cold {
				alert { numberOfDays }
			}
			updated
		}
	}
}
`
