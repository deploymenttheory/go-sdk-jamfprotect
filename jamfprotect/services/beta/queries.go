package beta

const getBetaAcceptanceStatusQuery = `
query getBetaAcceptanceStatus {
	getAppInitializationData {
		betaAcceptanceStatus {
			betaName
			acceptedTimestamp
			acceptedUser
		}
	}
}
`

const updateBetaAcceptanceStatusMutation = `
mutation updateBetaAcceptanceStatus($betaName: BetaName!) {
	updateBetaAcceptanceStatus(input: {betaName: $betaName}) {
		betaAcceptanceStatus {
			betaName
			acceptedTimestamp
			acceptedUser
		}
	}
}
`
