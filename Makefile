UNBOUND = unbound-buddy

BUILD_DIR = ./build
PREFIX = /usr/local

all: ${BUILD_DIR} ${BUILD_DIR}/${UNBOUND}

${BUILD_DIR}:
	mkdir -p "${BUILD_DIR}"

${BUILD_DIR}/${UNBOUND}:
	go build -v -o "${BUILD_DIR}/${UNBOUND}" cmd/${UNBOUND}/${UNBOUND}.go

strip:
	strip -v "${BUILD_DIR}/${UNBOUND}"

clean:
	[ -d "${BUILD_DIR}" ] && rm -rf "${BUILD_DIR}" || true