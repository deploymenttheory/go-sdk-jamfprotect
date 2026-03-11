package role

// GraphQL fragments, queries, and mutations for Role

const roleFields = `
fragment RoleFields on Role {
  id
  name
  permissions {
    R
    W
  }
  created
  updated
}
`

const listRolesQuery = `
query listRoles($pageSize: Int, $nextToken: String, $direction: OrderDirection!, $field: RoleOrderField!) {
  listRoles(
    input: {next: $nextToken, pageSize: $pageSize, order: {direction: $direction, field: $field}}
  ) {
    items { ...RoleFields }
    pageInfo { next total }
  }
}
` + roleFields

const getRoleQuery = `
query getRole($id: ID!) {
  getRole(id: $id) { ...RoleFields }
}
` + roleFields

const createRoleMutation = `
mutation createRole($name: String!, $readResources: [RBAC_RESOURCE!]!, $writeResources: [RBAC_RESOURCE!]!) {
  createRole(input: {name: $name, readResources: $readResources, writeResources: $writeResources}) {
    ...RoleFields
  }
}
` + roleFields

const updateRoleMutation = `
mutation updateRole($id: ID!, $name: String!, $readResources: [RBAC_RESOURCE!]!, $writeResources: [RBAC_RESOURCE!]!) {
  updateRole(id: $id, input: {name: $name, readResources: $readResources, writeResources: $writeResources}) {
    ...RoleFields
  }
}
` + roleFields

const deleteRoleMutation = `
mutation deleteRole($id: ID!) {
  deleteRole(id: $id) { ...RoleFields }
}
` + roleFields
