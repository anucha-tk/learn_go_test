test:
	go test .

test-v:
	go test -v .

test-cover:
	go test -cover .

test-profile:
	go test -coverprofile=coverage.out

test-cover-html:
	go tool cover -html=coverage.out -o coverage.html

open-cover:
	explorer.exe coverage.html

open-cover-w:
	$(MAKE) test-profile;
	$(MAKE) test-cover-html;
	$(MAKE) open-cover;
