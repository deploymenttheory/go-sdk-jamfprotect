package user

// GraphQL fragments, queries, and mutations for User

const userFields = `
fragment UserFields on User {
  id
  email
  sub @skip(if: $hasLimitedAppAccess)
  connection @include(if: $RBAC_Connection) { id name requireKnownUsers source }
  assignedRoles @skip(if: $hasLimitedAppAccess) @include(if: $RBAC_Role) { id name }
  assignedGroups @skip(if: $hasLimitedAppAccess) @include(if: $RBAC_Group) {
    id
    name
    assignedRoles @include(if: $RBAC_Role) { id name }
  }
  lastLogin
  source @skip(if: $hasLimitedAppAccess)
  receiveEmailAlert
  emailAlertMinSeverity
  extraAttrs @skip(if: $hasLimitedAppAccess)
  created
  updated
}
`

const listUsersQuery = `
query listUsers($pageSize: Int, $nextToken: String, $direction: OrderDirection!, $field: UserOrderField!, $hasLimitedAppAccess: Boolean!, $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!, $RBAC_Group: Boolean!) {
  listUsers(input: {next: $nextToken, pageSize: $pageSize, order: {direction: $direction, field: $field}}) {
    items { ...UserFields }
    pageInfo { next total }
  }
}
` + userFields

const getUserQuery = `
query getUser($id: ID!, $hasLimitedAppAccess: Boolean!, $RBAC_Connection: Boolean!, $RBAC_Role: Boolean!, $RBAC_Group: Boolean!) {
  getUser(id: $id) { ...UserFields }
}
` + userFields

const createUserMutation = `
mutation createUser($email: AWSEmail!, $roleIds: [ID], $groupIds: [ID], $connectionId: ID, $receiveEmailAlert: Boolean!, $emailAlertMinSeverity: SEVERITY!, $RBAC_Role: Boolean!, $RBAC_Group: Boolean!, $RBAC_Connection: Boolean!, $hasLimitedAppAccess: Boolean!) {
  createUser(input: {email: $email, roleIds: $roleIds, groupIds: $groupIds, connectionId: $connectionId, receiveEmailAlert: $receiveEmailAlert, emailAlertMinSeverity: $emailAlertMinSeverity}) {
    ...UserFields
  }
}
` + userFields

const updateUserMutation = `
mutation updateUser($id: ID!, $roleIds: [ID], $groupIds: [ID], $receiveEmailAlert: Boolean!, $emailAlertMinSeverity: SEVERITY!, $RBAC_Role: Boolean!, $RBAC_Group: Boolean!, $RBAC_Connection: Boolean!, $hasLimitedAppAccess: Boolean!) {
  updateUser(id: $id, input: {roleIds: $roleIds, groupIds: $groupIds, receiveEmailAlert: $receiveEmailAlert, emailAlertMinSeverity: $emailAlertMinSeverity}) {
    ...UserFields
  }
}
` + userFields

const deleteUserMutation = `
mutation deleteUser($id: ID!) {
  deleteUser(id: $id) { id }
}
`
