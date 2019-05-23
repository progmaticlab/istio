#-----------------------------------------------------------------------------
# Target: test.conformance.*
#-----------------------------------------------------------------------------

# The following flags (in addition to ${V}) can be specified on the command-line, or the environment. This
# is primarily used by the CI systems.

# $(CI) specifies that the test is running in a CI system. This enables CI specific logging.
_CONFORMANCE_TEST_CIMODE_FLAG =
_CONFORMANCE_TEST_PULL_POLICY = Always
ifneq ($(CI),)
	_CONFORMANCE_TEST_CIMODE_FLAG = --istio.test.ci
	_CONFORMANCE_TEST_PULL_POLICY = IfNotPresent      # Using Always in CircleCI causes pull issues as images are local.
endif

# In Prow, ARTIFACTS_DIR points to the location where Prow captures the artifacts from the tests
CONFORMANCE_TEST_WORKDIR =
ifneq ($(ARTIFACTS_DIR),)
	CONFORMANCE_TEST_WORKDIR = ${ARTIFACTS_DIR}
endif

_CONFORMANCE_TEST_INGRESS_FLAG =
ifeq (${TEST_ENV},minikube)
    _CONFORMANCE_TEST_INGRESS_FLAG = --istio.test.kube.minikube
else ifeq (${TEST_ENV},minikube-none)
    _CONFORMANCE_TEST_INGRESS_FLAG = --istio.test.kube.minikube
endif


# $(CONFORMANCE_TEST_WORKDIR) specifies the working directory for the tests. If not specified, then a
# temporary folder is used.
_CONFORMANCE_TEST_WORKDIR_FLAG =
ifneq ($(CONFORMANCE_TEST_WORKDIR),)
    _CONFORMANCE_TEST_WORKDIR_FLAG = --istio.test.work_dir $(CONFORMANCE_TEST_WORKDIR)
endif

# $(CONFORMANCE_TEST_KUBECONFIG) specifies the kube config file to be used. If not specified, then
# ~/.kube/config is used.
CONFORMANCE_TEST_KUBECONFIG = ~/.kube/config
ifneq ($(KUBECONFIG),)
    CONFORMANCE_TEST_KUBECONFIG = $(KUBECONFIG)
endif

JUNIT_UNIT_TEST_XML ?= $(ISTIO_OUT)/junit_unit-tests.xml
JUNIT_REPORT = $(shell which go-junit-report 2> /dev/null || echo "${ISTIO_BIN}/go-junit-report")

test.conformance.kube: | $(JUNIT_REPORT)
	mkdir -p $(dir $(JUNIT_UNIT_TEST_XML))
	set -o pipefail; \
	$(GO) test -p 1 ${T} ./tests/conformance/... ${_CONFORMANCE_TEST_WORKDIR_FLAG} ${_CONFORMANCE_TEST_CIMODE_FLAG} -timeout 30m \
	--istio.test.env kube \
	--istio.test.kube.config ${CONFORMANCE_TEST_KUBECONFIG} \
	--istio.test.hub=${HUB} \
	--istio.test.tag=${TAG} \
	--istio.test.pullpolicy=${_CONFORMANCE_TEST_PULL_POLICY} \
	${_CONFORMANCE_TEST_INGRESS_FLAG} \
	2>&1 | tee >($(JUNIT_REPORT) > $(JUNIT_UNIT_TEST_XML))

test.conformance.local: | $(JUNIT_REPORT)
	mkdir -p $(dir $(JUNIT_UNIT_TEST_XML))
	set -o pipefail; \
	$(GO) test -p 1 ${T} ./tests/conformance/... \
	--istio.test.env native \
	2>&1 | tee >($(JUNIT_REPORT) > $(JUNIT_UNIT_TEST_XML))
