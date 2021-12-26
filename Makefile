wsLoadTest:
	@./scripts/amx.sh

checkGO:
	@./scripts/checkGO.sh

checkGIT:
	@./scripts/checkGIT.sh

checkPreCommit:
	@./scripts/checkPreCommit.sh

checkSwagger:
	@./scripts/checkSwagger.sh

checkGSED:
	@./scripts/checkGSED.sh

checkGolangCILint:
	@./scripts/checkGolangCILint.sh

doctor: wsLoadTest checkGO checkGIT checkPreCommit checkSwagger checkGolangCILint checkGSED

init: wsLoadTest
	@./scripts/init.sh

install: wsLoadTest
	@./scripts/install.sh

swagger: wsLoadTest
	@./scripts/swagger.sh

verify: wsLoadTest
	@./scripts/verify.sh

test: wsLoadTest
	@./scripts/test.sh

build: wsLoadTest
	@./scripts/build.sh
