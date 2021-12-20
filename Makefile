naruto:
	@./scripts/naruto.sh

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

doctor: naruto checkGO checkGIT checkPreCommit checkSwagger checkGolangCILint checkGSED

init: naruto
	@./scripts/init.sh

install: naruto
	@./scripts/install.sh

swagger: naruto
	@./scripts/swagger.sh

verify: naruto
	@./scripts/verify.sh

test: naruto
	@./scripts/test.sh

build: naruto
	@./scripts/build.sh
