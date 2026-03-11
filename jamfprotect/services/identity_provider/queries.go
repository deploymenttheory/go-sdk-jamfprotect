package identityprovider

// GraphQL fragments and queries for IdentityProvider (Connections)

const connectionFields = `
fragment ConnectionFields on Connection {
  id
  name
  requireKnownUsers
  button
  created
  updated
  strategy
  groupsSupport
  source
}
`

const listConnectionsQuery = `
query listConnections($pageSize: Int, $nextToken: String, $direction: OrderDirection!, $field: ConnectionOrderField!) {
  listConnections(
    input: {next: $nextToken, pageSize: $pageSize, order: {direction: $direction, field: $field}}
  ) {
    items { ...ConnectionFields }
    pageInfo { next total }
  }
}
` + connectionFields
