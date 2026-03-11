package group

// GraphQL fragments, queries, and mutations for Group

const groupFields = `
fragment GroupFields on Group {
  id
  name
  connection @include(if: $RBAC_Connection) {
    id
    name
  }
  assignedRoles @include(if: $RBAC_Role) {
    id
    name
  }
  accessGroup
  created
  updated
}
`

const listGroupsQuery = `
query listGroups($pageSize: Int, $nextToken: String, $direction: OrderDirection!, $field: GroupOrderField!, $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!) {
  listGroups(
    input: {next: $nextToken, pageSize: $pageSize, order: {direction: $direction, field: $field}}
  ) {
    items { ...GroupFields }
    pageInfo { next total }
  }
}
` + groupFields

const getGroupQuery = `
query getGroup($id: ID!, $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!) {
  getGroup(id: $id) { ...GroupFields }
}
` + groupFields

const createGroupMutation = `
mutation createGroup($name: String!, $connectionId: ID, $accessGroup: Boolean, $roleIds: [ID], $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!) {
  createGroup(input: {name: $name, connectionId: $connectionId, accessGroup: $accessGroup, roleIds: $roleIds}) {
    ...GroupFields
  }
}
` + groupFields

const updateGroupMutation = `
mutation updateGroup($id: ID!, $name: String!, $accessGroup: Boolean, $roleIds: [ID], $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!) {
  updateGroup(id: $id, input: {name: $name, accessGroup: $accessGroup, roleIds: $roleIds}) {
    ...GroupFields
  }
}
` + groupFields

const deleteGroupMutation = `
mutation deleteGroup($id: ID!) {
  deleteGroup(id: $id) { id }
}
`
