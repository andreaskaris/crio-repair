FROM quay.io/centos/centos:stream8
RUN dnf install -y 'dnf-command(config-manager)'
RUN dnf config-manager --set-enabled powertools
# dnf install epel-release epel-next-release
RUN yum -y install libseccomp-devel gpgme-devel device-mapper-devel libassuan-devel libassuan make golang

WORKDIR /crio-repair
COPY . .
RUN make build
