//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-che-mt/$CHE_VERSION/fabric8-tenant-che-mt-$CHE_VERSION-openshift.yml > fabric8-tenant-che-mt-openshift.yml"

//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-che/$CHE_VERSION/fabric8-tenant-che-$CHE_VERSION-openshift.yml > fabric8-tenant-che-openshift.yml"
//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-che-quotas-oso/$CHE_VERSION/fabric8-tenant-che-quotas-oso-$CHE_VERSION-openshift.yml > fabric8-tenant-che-quotas-oso-openshift.yml"

//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-jenkins/$JENKINS_VERSION/fabric8-tenant-jenkins-$JENKINS_VERSION-openshift.yml > fabric8-tenant-jenkins-openshift.yml"
//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-jenkins-quotas-oso/$JENKINS_VERSION/fabric8-tenant-jenkins-quotas-oso-$JENKINS_VERSION-openshift.yml > fabric8-tenant-jenkins-quotas-oso-openshift.yml"

//go:generate sh -c "curl http://central.maven.org/maven2/io/fabric8/tenant/packages/fabric8-tenant-team/$TEAM_VERSION/fabric8-tenant-team-$TEAM_VERSION-openshift.yml > fabric8-tenant-team-openshift.yml"

package template
