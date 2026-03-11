package downloads

// GraphQL queries for Downloads

const getOrganizationDownloadsQuery = `
query getOrganizationDownloads {
  downloads: getOrganizationDownloads {
    pppc
    rootCA
    csr
    installerUuid
    vanillaPackage {
      version
    }
    websocket_auth
    tamperPreventionProfile
  }
}
`
