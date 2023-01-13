# You can pick any Debian/Ubuntu-based image. 😊
FROM mcr.microsoft.com/devcontainers/base:bullseye

# Options for setup script
ARG INSTALL_ZSH="true"
ARG UPGRADE_PACKAGES="false"
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID
# Install needed packages and setup non-root user. Use a separate RUN statement to add your own dependencies.
COPY library-scripts/*.sh /tmp/library-scripts/
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && /bin/bash /tmp/library-scripts/common-debian.sh "${INSTALL_ZSH}" "${USERNAME}" "${USER_UID}" "${USER_GID}" "${UPGRADE_PACKAGES}" "true" "true" \
    && apt-get install -y pkg-config zip g++ zlib1g-dev unzip python \
    && apt-get autoremove -y && apt-get clean -y && rm -rf /var/lib/apt/lists/* /tmp/library-scripts

# Install Bazel
ARG BAZELISK_VERSION=v1.10.1
ARG BAZELISK_DOWNLOAD_SHA=dev-mode
RUN curl -fSsL -o /usr/local/bin/bazelisk https://github.com/bazelbuild/bazelisk/releases/download/${BAZELISK_VERSION}/bazelisk-linux-amd64 \
    && ([ "${BAZELISK_DOWNLOAD_SHA}" = "dev-mode" ] || echo "${BAZELISK_DOWNLOAD_SHA} */usr/local/bin/bazelisk" | sha256sum --check - ) \
    && chmod 0755 /usr/local/bin/bazelisk \
    && ln -s  /usr/local/bin/bazelisk  /usr/local/bin/bazel

RUN curl -fSsl -o /tmp/go1.19.5.linux-amd64.tar.gz https://dl.google.com/go/go1.19.5.linux-amd64.tar.gz \
    && rm -rf /usr/local/go \
    && tar -C /usr/local -xzf /tmp/go1.19.5.linux-amd64.tar.gz

RUN echo 'export PATH=$PATH:/usr/local/go/bin' >> "/home/${USERNAME}/.profile" \
    && echo 'export PATH=$PATH:$(go env GOPATH)' >> "/home/${USERNAME}/.profile" \
    && sudo -u "${USERNAME}" /usr/local/go/bin/go install -v golang.org/x/tools/gopls@latest \
    && sudo -u "${USERNAME}" /usr/local/go/bin/go install -v github.com/go-delve/delve/cmd/dlv@latest \
    && sudo -u "${USERNAME}" /usr/local/go/bin/go install -v honnef.co/go/tools/cmd/staticcheck@latest